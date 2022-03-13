package impl

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"strings"

	"github.com/bitwomrhole-express/express/community-server/app/data/dao"
	"github.com/bitwomrhole-express/express/community-server/app/data/entity"
	"github.com/bitwomrhole-express/express/community-server/app/service"
	"github.com/bitwomrhole-express/express/community-server/app/web/dto"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
)

// PasswordServiceImpl 实现密码服务
type PasswordServiceImpl struct {
	markup.Component `id:"express-PasswordService"`

	AccountDAO       dao.Account                      `inject:"#express-data-account-dao"`
	EmailVeriService service.EmailVerificationService `inject:"#express-EmailVerificationService"`
}

func (inst *PasswordServiceImpl) _Impl() service.PasswordService {
	return inst
}

func (inst *PasswordServiceImpl) verifyEmail(o *dto.Password) error {
	return inst.EmailVeriService.Verify(&o.Verification)
}

func (inst *PasswordServiceImpl) verifyPassword(username string, password []byte) error {

	// find account
	account, err := inst.AccountDAO.Find(username)
	if err != nil {
		return err
	}

	// check password
	salt := account.Salt
	want := account.Password
	pc := passwordComputer{}
	pc.purePassword = password
	pc.setSalt(salt)
	return pc.verify(want, 10)
}

// Verify 验证密码和用户名
func (inst *PasswordServiceImpl) Verify(username string, password []byte) error {
	return inst.verifyPassword(username, password)
}

func (inst *PasswordServiceImpl) makeSalt(o *dto.Password) util.Base64 {

	now := util.Now()
	nonce := make([]byte, 64)
	rand.Reader.Read(nonce)
	hex := util.HexFromBytes(nonce)

	builder := strings.Builder{}
	builder.WriteString(now.String())
	builder.WriteString(hex.String())
	builder.WriteString(o.Verification.Email)

	text := builder.String()
	sum := md5.Sum([]byte(text))
	return util.Base64FromBytes(sum[:])
}

// CreateAccount 创建账号
func (inst *PasswordServiceImpl) CreateAccount(o *dto.Password) error {

	// verify
	err := inst.verifyEmail(o)
	if err != nil {
		return err
	}

	// make password
	email := o.Verification.Email
	salt := inst.makeSalt(o)

	pc := passwordComputer{}
	pc.setSalt(salt)
	pc.purePassword = o.NewPassword.Bytes()
	saltyPassword := pc.makeSaltyPassword(5)

	// insert new account
	account := &entity.Account{
		Email:    email,
		Salt:     salt,
		Password: saltyPassword,
		Username: "todo",
		Nickname: "todo",
	}
	_, err = inst.AccountDAO.Insert(account)
	return err
}

// ResetPassword 重设密码、修改密码
func (inst *PasswordServiceImpl) ResetPassword(o *dto.Password) error {

	// verify
	err := inst.verifyEmail(o)
	if err != nil {
		return err
	}

	// find account
	email := o.Verification.Email
	account, err := inst.AccountDAO.FindByEmail(email)
	if err != nil {
		return err
	}

	// make new password
	const repeat = 5
	salt := inst.makeSalt(o)
	pc := passwordComputer{}
	pc.setSalt(salt)
	pc.purePassword = o.NewPassword.Bytes()
	account.Salt = salt
	account.Password = pc.makeSaltyPassword(repeat)

	// save account
	_, err = inst.AccountDAO.Update(account.ID, account)
	return err
}

////////////////////////////////////////////////////////////////////////////////

type passwordComputer struct {
	purePassword  []byte
	saltyPassword []byte
	salt          []byte
	generation    int
}

func (inst *passwordComputer) sha256sum(data []byte) []byte {
	sum := sha256.Sum256(data)
	return sum[:]
}

func (inst *passwordComputer) next() {

	pp1 := inst.purePassword
	pp2 := inst.sha256sum(pp1)

	mixer := bytes.Buffer{}
	mixer.Write(pp2)
	mixer.Write(inst.salt)

	sp := inst.sha256sum(mixer.Bytes())
	sp = inst.sha256sum(sp)
	sp = inst.sha256sum(sp)
	sp = inst.sha256sum(sp)

	inst.saltyPassword = sp
	inst.purePassword = pp2
	inst.generation++
}

func (inst *passwordComputer) setSalt(salt util.Base64) {
	inst.salt = salt.Bytes()
}

func (inst *passwordComputer) verify(want string, ttl int) error {
	pc2 := passwordComputer{}
	pc2 = *inst
	pc2.next()
	pc2.next()
	for ; ttl > 0; ttl-- {
		hex := util.HexFromBytes(pc2.saltyPassword)
		have := hex.String()
		if want == have {
			return nil
		}
		pc2.next()
	}
	return errors.New("bad auth")
}

func (inst *passwordComputer) makeSaltyPassword(repeat int) string {
	pc2 := passwordComputer{}
	pc2 = *inst
	for ; repeat > 0; repeat-- {
		pc2.next()
	}
	hex := util.HexFromBytes(pc2.saltyPassword)
	return hex.String()
}
