package security

import (
	"context"
	"encoding/json"
	"strings"

	"bitwomrhole.com/djaf/express-go-server/server/data/dao"
	"bitwomrhole.com/djaf/express-go-server/server/service"
	"bitwomrhole.com/djaf/express-go-server/server/web/dto"
	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter-security/keeper/users"
	"github.com/bitwormhole/starter/markup"
)

type EmailAuthenticator struct {
	markup.Component `class:"keeper-authenticator-registry"`

	AccountDAO       dao.Account                      `inject:"#express-data-account-dao"`
	EmailVeriService service.EmailVerificationService `inject:"#express-EmailVerificationService"`
}

func (inst *EmailAuthenticator) _Impl() (keeper.AuthenticatorRegistry, keeper.Authenticator) {
	return inst, inst
}

func (inst *EmailAuthenticator) GetRegistrationList() []*keeper.AuthenticatorRegistration {
	ar := &keeper.AuthenticatorRegistration{
		Name:          "email",
		Mechanism:     "email",
		Authenticator: inst,
	}
	return []*keeper.AuthenticatorRegistration{ar}
}

func (inst *EmailAuthenticator) Supports(ctx context.Context, a keeper.Authentication) bool {
	mech := a.Mechanism()
	return strings.ToLower(mech) == "email"
}

func (inst *EmailAuthenticator) Verify(ctx context.Context, a keeper.Authentication) (keeper.Identity, error) {

	user := a.User()
	secret := a.Secret()

	// 验证邮件
	veri := &dto.EmailVerification{}
	err := json.Unmarshal(secret, veri)
	if err != nil {
		return nil, err
	}
	err = inst.EmailVeriService.Verify(veri)
	if err != nil {
		return nil, err
	}

	// 查找账号
	account, err := inst.AccountDAO.FindByEmail(user)
	if err != nil {
		return nil, err
	}

	// 创建结果
	ib := keeper.IdentityBuilder{}
	ib.Avatar = account.Avatar
	ib.Nickname = account.Nickname
	ib.Roles = users.Roles(account.Roles)
	ib.UserID = users.UserID(account.ID)
	ib.UserName = users.UserName(account.Username)
	ib.UserUUID = users.UserUUID(account.UUID)

	id := ib.Identity()
	return id, nil
}
