package impl

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"sort"
	"strconv"
	"strings"

	"github.com/bitwomrhole-express/express/community-server/app/service"
	"github.com/bitwomrhole-express/express/community-server/app/web/dto"
	"github.com/bitwormhole/starter-mail/mail"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
)

// EmailVerificationServiceImpl 实现邮件验证服务
type EmailVerificationServiceImpl struct {
	markup.Component `id:"express-EmailVerificationService" initMethod:"Init"`

	Context          application.Context `inject:"context"`
	MailTemplateName string              `inject:"${email.verification.template}"` // 验证邮件标题
	MailTitle        string              `inject:"${email.verification.title}"`    // 验证邮件标题
	Sender           mail.Sender         `inject:"#mail.Sender"`
	SenderAddr       string              `inject:"${mail.sender.address}"`

	mailContentType     string
	mailContentTemplate string // 验证邮件模板
	secret              string
}

func (inst *EmailVerificationServiceImpl) _Impl() service.EmailVerificationService {
	return inst
}

// Init 初始化
func (inst *EmailVerificationServiceImpl) Init() error {
	inst.secret = inst.makeRandomWithSize(256)
	return inst.initMailTemplate()
}

// 加载邮件模板
func (inst *EmailVerificationServiceImpl) initMailTemplate() error {
	name := inst.MailTemplateName
	text, err := inst.Context.GetResources().GetText(name)
	if err != nil {
		return err
	}
	ctype := "text/plain"
	if strings.HasSuffix(name, ".html") {
		ctype = "text/html"
	}
	inst.mailContentType = ctype
	inst.mailContentTemplate = text
	return nil
}

// StartVerification 开始验证：生成验证码，并发送邮件；
func (inst *EmailVerificationServiceImpl) StartVerification(email string) (*dto.EmailVerification, error) {

	const ttl = (1000 * 300) // 验证码保质期（单位：毫秒）
	vcode := inst.generateVeriCode()
	now := util.Now()

	verification := &dto.EmailVerification{
		Code:      vcode,
		Email:     email,
		TimeBegin: now,
		TimeEnd:   now + ttl,
		// Nonce: nonce,
		// Secret:    secret,
	}

	err := inst.sendMail(verification)
	if err != nil {
		return nil, err
	}

	nonce := inst.generateNonce()
	secret := inst.getServiceSecret()

	verification.Secret = secret // 设置密钥
	verification.Nonce = nonce
	sum := inst.computeSum(verification)

	verification.Secret = "" // 清除密钥
	verification.Code = ""   // 清除验证码
	verification.Sha256sum = sum
	return verification, nil
}

// Verify 执行验证：检查验证码是否正确
func (inst *EmailVerificationServiceImpl) Verify(v *dto.EmailVerification) error {

	secret := inst.getServiceSecret()
	sum1 := v.Sha256sum
	now := util.Now()

	if now < v.TimeBegin || v.TimeEnd < now {
		return errors.New("bad verification")
	}

	v2 := dto.EmailVerification{}
	v2 = *v
	v2.Secret = secret
	sum2 := inst.computeSum(&v2)

	if sum1 != sum2 {
		return errors.New("bad verification")
	}

	return nil
}

func (inst *EmailVerificationServiceImpl) computeSum(v *dto.EmailVerification) string {

	t1 := strconv.FormatInt(v.TimeBegin.Int64(), 10)
	t2 := strconv.FormatInt(v.TimeEnd.Int64(), 10)
	list := make([]string, 0)

	list = append(list, v.Code)
	list = append(list, v.Email)
	list = append(list, v.Nonce)
	list = append(list, v.Secret)
	list = append(list, t1)
	list = append(list, t2)

	sort.Strings(list)
	builder := strings.Builder{}

	for _, item := range list {
		builder.WriteString(item)
		builder.WriteRune('\n')
	}

	text := builder.String()
	sum := sha256.Sum256([]byte(text))
	return util.StringifyBytes(sum[:])
}

func (inst *EmailVerificationServiceImpl) generateVeriCode() string {
	const wantLen = 8
	nonce := inst.makeRandomWithSize(32)
	b64 := util.Base64FromString(nonce)
	hex := util.HexFromBytes(b64.Bytes())
	text := hex.String()
	chs := []rune(text)
	builder := strings.Builder{}
	for _, c := range chs {
		if '0' <= c && c <= '9' {
			builder.WriteRune(c)
		}
		if builder.Len() >= wantLen {
			break
		}
	}
	return builder.String()
}

func (inst *EmailVerificationServiceImpl) generateNonce() string {
	return inst.makeRandomWithSize(48)
}

func (inst *EmailVerificationServiceImpl) makeRandomWithSize(size int) string {

	// step 1
	now := util.Now()
	bin1 := make([]byte, size)
	rand.Reader.Read(bin1)

	builder := bytes.Buffer{}
	builder.WriteString(now.String())
	builder.Write(bin1)
	sum := sha256.Sum256(builder.Bytes())

	// step 2
	bin2 := make([]byte, size)
	rand.Reader.Read(bin2)

	builder.Reset()
	builder.Write(sum[:])
	builder.Write(bin2)

	// final
	bin3 := builder.Bytes()
	b64 := util.Base64FromBytes(bin3)
	return b64.String()
}

func (inst *EmailVerificationServiceImpl) getServiceSecret() string {
	return inst.secret
}

func (inst *EmailVerificationServiceImpl) sendMail(v *dto.EmailVerification) error {
	from := inst.SenderAddr
	to := []string{v.Email}
	content := inst.makeMailContent(v)
	m := &mail.Mail{
		FromAddr:    from,
		ToAddr:      to,
		Title:       inst.MailTitle,
		ContentType: inst.mailContentType,
		Content:     content,
	}
	return inst.Sender.Send(m)
}

func (inst *EmailVerificationServiceImpl) makeMailContent(v *dto.EmailVerification) string {
	ttl := (v.TimeEnd - v.TimeBegin) / (1000 * 60)
	ttlstr := strconv.FormatInt(ttl.Int64(), 10)
	params := map[string]string{
		"verification.code":    v.Code,
		"verification.address": v.Email,
		"verification.ttl.min": ttlstr,
	}
	text := inst.mailContentTemplate
	for k, v := range params {
		text = strings.ReplaceAll(text, "{{"+k+"}}", v)
	}
	return text
}
