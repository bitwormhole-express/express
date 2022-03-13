package controller

import (
	"net/http"

	"github.com/bitwomrhole-express/express/community-server/app/service"
	"github.com/bitwomrhole-express/express/community-server/app/web/vo"
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

// PasswordController 处理跟密码相关的功能
type PasswordController struct {
	markup.Component `class:"rest-controller"`

	Responder       glass.MainResponder     `inject:"#glass-main-responder"`
	PasswordService service.PasswordService `inject:"#express-PasswordService"`
}

func (inst *PasswordController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *PasswordController) Init(ec glass.EngineConnection) error {
	ec = ec.RequestMapping("password")
	ec.Handle(http.MethodPost, "", inst.handlePost)
	ec.Handle(http.MethodPut, "", inst.handlePut)
	return nil
}

// 注册，并创建密码
func (inst *PasswordController) handlePost(c *gin.Context) {
	req := innerPasswordRequest{
		controller: inst,
		gc:         c,
	}
	req.wantRequestBody = true
	err := req.open()
	if err == nil {
		err = req.doCreateAccount()
	}
	req.send(err)
}

// 重设密码
func (inst *PasswordController) handlePut(c *gin.Context) {
	req := innerPasswordRequest{
		controller: inst,
		gc:         c,
	}
	req.wantRequestBody = true
	err := req.open()
	if err == nil {
		err = req.doResetPassword()
	}
	req.send(err)
}

////////////////////////////////////////////////////////////////////////////////

type innerPasswordRequest struct {
	controller      *PasswordController
	gc              *gin.Context
	wantRequestBody bool
	rx              vo.Password
	tx              vo.Password
}

func (inst *innerPasswordRequest) open() error {
	if inst.wantRequestBody {
		inst.gc.BindJSON(&inst.rx)
	}
	return nil
}

func (inst *innerPasswordRequest) send(err error) {
	resp := &glass.Response{
		Context: inst.gc,
		Data:    &inst.tx,
		Error:   err,
		Status:  0,
	}
	inst.controller.Responder.Send(resp)
}

func (inst *innerPasswordRequest) doCreateAccount() error {
	o := &inst.rx.Password
	return inst.controller.PasswordService.CreateAccount(o)
}

func (inst *innerPasswordRequest) doResetPassword() error {
	o := &inst.rx.Password
	return inst.controller.PasswordService.ResetPassword(o)
}

////////////////////////////////////////////////////////////////////////////////
