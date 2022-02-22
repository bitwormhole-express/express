package security

import (
	"context"
	"strings"

	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter-security/keeper/users"
	"github.com/bitwormhole/starter/markup"
)

type LoginWithPasswordHandler struct {
	markup.Component `id:"keeper-authenticator-registry-2233" class:"keeper-authenticator-registry"`
}

func (inst *LoginWithPasswordHandler) _Impl() (keeper.AuthenticatorRegistry, keeper.Authenticator) {
	return inst, inst
}

func (inst *LoginWithPasswordHandler) GetRegistrationList() []*keeper.AuthenticatorRegistration {
	ar := &keeper.AuthenticatorRegistration{
		Name:          "password",
		Mechanism:     "password",
		Authenticator: inst,
	}
	return []*keeper.AuthenticatorRegistration{ar}
}

func (inst *LoginWithPasswordHandler) Supports(ctx context.Context, a keeper.Authentication) bool {
	mech := a.Mechanism()
	return strings.ToLower(mech) == "password"
}

func (inst *LoginWithPasswordHandler) Verify(ctx context.Context, a keeper.Authentication) (keeper.Identity, error) {

	user := a.User()
	ib := keeper.IdentityBuilder{}

	ib.Avatar = "https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png"
	ib.Nickname = user
	ib.Roles = "guest,user"
	ib.UserID = 0
	ib.UserName = users.UserName(user)
	ib.UserUUID = ""

	id := ib.Identity()
	// return nil, errors.New("todo: no impl")
	return id, nil
}
