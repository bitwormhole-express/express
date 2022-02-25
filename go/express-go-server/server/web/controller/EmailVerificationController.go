package controller

import (
	"net/http"

	"bitwomrhole.com/djaf/express-go-server/server/service"
	"bitwomrhole.com/djaf/express-go-server/server/web/vo"
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

// EmailVerificationController 邮件验证控制器
type EmailVerificationController struct {
	markup.Component `class:"rest-controller"`

	EmailVeriService service.EmailVerificationService `inject:"#express-EmailVerificationService"`
	Responder        glass.MainResponder              `inject:"#glass-main-responder"`
}

func (inst *EmailVerificationController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *EmailVerificationController) Init(ec glass.EngineConnection) error {
	ec = ec.RequestMapping("email-verification")
	ec.Handle(http.MethodPost, "", inst.handlePost)
	return nil
}

func (inst *EmailVerificationController) handlePost(c *gin.Context) {
	req := innerEmailVerificationRequest{
		controller: inst,
		gc:         c,
	}
	req.wantRequestBody = true
	err := req.open()
	if err == nil {
		err = req.doPost()
	}
	req.send(err)
}

////////////////////////////////////////////////////////////////////////////////

type innerEmailVerificationRequest struct {
	controller *EmailVerificationController
	gc         *gin.Context

	wantRequestBody bool

	rx vo.EmailVerification
	tx vo.EmailVerification
}

func (inst *innerEmailVerificationRequest) open() error {
	if inst.wantRequestBody {
		inst.gc.BindJSON(&inst.rx)
	}
	return nil
}

func (inst *innerEmailVerificationRequest) send(err error) {
	resp := &glass.Response{}
	resp.Context = inst.gc
	resp.Data = &inst.tx
	resp.Error = err
	resp.Status = 0
	inst.controller.Responder.Send(resp)
}

func (inst *innerEmailVerificationRequest) doPost() error {
	email := inst.rx.Verification.Email
	o, err := inst.controller.EmailVeriService.StartVerification(email)
	if err != nil {
		return err
	}
	inst.tx.Verification = *o
	return nil
}
