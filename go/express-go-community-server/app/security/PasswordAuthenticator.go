package security

import (
	"context"
	"errors"
	"strings"

	"github.com/bitwomrhole-express/express/community-server/app/data/dao"
	"github.com/bitwomrhole-express/express/community-server/app/service"
	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter-security/keeper/users"
	"github.com/bitwormhole/starter/markup"
)

type PasswordAuthenticator struct {
	markup.Component `class:"keeper-authenticator-registry"`

	AccountDAO      dao.Account             `inject:"#express-data-account-dao"`
	PasswordService service.PasswordService `inject:"#express-PasswordService"`
}

func (inst *PasswordAuthenticator) _Impl() (keeper.AuthenticatorRegistry, keeper.Authenticator) {
	return inst, inst
}

func (inst *PasswordAuthenticator) GetRegistrationList() []*keeper.AuthenticatorRegistration {
	ar := &keeper.AuthenticatorRegistration{
		Name:          "password",
		Mechanism:     "password",
		Authenticator: inst,
	}
	return []*keeper.AuthenticatorRegistration{ar}
}

func (inst *PasswordAuthenticator) Supports(ctx context.Context, a keeper.Authentication) bool {
	mech := a.Mechanism()
	return strings.ToLower(mech) == "password"
}

// Verify 验证密码是否正确
func (inst *PasswordAuthenticator) Verify(ctx context.Context, a keeper.Authentication) (keeper.Identity, error) {

	// find account
	word := a.User()
	account, err := inst.AccountDAO.Find(word)
	if err != nil {
		return nil, err
	}

	// verify
	err = inst.PasswordService.Verify(account.ID.String(), a.Secret())
	if err != nil {
		return nil, errors.New("bad username or password")
	}

	// make identity
	ib := keeper.IdentityBuilder{}
	ib.Avatar = account.Avatar
	ib.Email = account.Email
	ib.Nickname = account.Nickname
	ib.Roles = users.Roles(account.Roles)
	ib.UserID = users.UserID(account.ID)
	ib.UserName = users.UserName(account.Username)
	ib.UserUUID = users.UserUUID(account.UUID)
	identity := ib.Identity()
	return identity, nil
}
