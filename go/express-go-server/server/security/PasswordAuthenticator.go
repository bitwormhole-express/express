package security

import (
	"context"
	"errors"
	"strings"

	"bitwomrhole.com/djaf/express-go-server/server/data/dao"
	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter-security/keeper/users"
	"github.com/bitwormhole/starter/markup"
)

type PasswordAuthenticator struct {
	markup.Component `class:"keeper-authenticator-registry"`

	AccountDAO dao.Account `inject:"#express-data-account-dao"`
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

func (inst *PasswordAuthenticator) Verify(ctx context.Context, a keeper.Authentication) (keeper.Identity, error) {

	word := a.User()
	account, err := inst.AccountDAO.Find(word)
	if err != nil {
		return nil, err
	}

	pass1 := string(a.Secret())
	pass2 := account.Password
	if pass1 != pass2 {
		return nil, errors.New("bad username or password")
	}

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
