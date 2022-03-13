package controller

import (
	"net/http"

	"github.com/bitwomrhole-express/express/community-server/app/service"
	"github.com/bitwomrhole-express/express/community-server/app/web/dto"
	"github.com/bitwomrhole-express/express/community-server/app/web/vo"
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

// UploadController 处理跟密码相关的功能
type UploadController struct {
	markup.Component `class:"rest-controller"`

	Responder     glass.MainResponder   `inject:"#glass-main-responder"`
	UploadService service.UploadService `inject:"#express-UploadService"`
}

func (inst *UploadController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *UploadController) Init(ec glass.EngineConnection) error {
	ec = ec.RequestMapping("upload")
	ec.Handle(http.MethodPost, "", inst.handlePost)
	ec.Handle(http.MethodPut, "", inst.handlePut)
	ec.Handle(http.MethodPut, ":id", inst.handlePut)
	return nil
}

func (inst *UploadController) handlePost(c *gin.Context) {
	req := innerUploadRequest{
		controller: inst,
		gc:         c,
	}
	req.wantID = false
	req.wantRequestBody = true
	err := req.open()
	if err == nil {
		err = req.doPost()
	}
	req.send(err)
}

func (inst *UploadController) handlePut(c *gin.Context) {
	req := innerUploadRequest{
		controller: inst,
		gc:         c,
	}
	req.wantID = true
	req.wantRequestBody = true
	err := req.open()
	if err == nil {
		err = req.doPut()
	}
	req.send(err)
}

///////////////////////////////////////////////////////////////////////////////

type innerUploadRequest struct {
	controller *UploadController
	gc         *gin.Context

	wantID          bool
	wantRequestBody bool

	id    string
	body1 vo.Upload
	body2 vo.Upload
}

func (inst *innerUploadRequest) open() error {
	c := inst.gc
	if inst.wantID {
		inst.id = c.Param("id")
	}
	if inst.wantRequestBody {
		c.BindJSON(&inst.body1)
	}
	return nil
}

func (inst *innerUploadRequest) send(err error) {
	status := inst.body2.Status
	resp := &glass.Response{
		Context: inst.gc,
		Data:    &inst.body2,
		Error:   err,
		Status:  status,
	}
	inst.controller.Responder.Send(resp)
}

func (inst *innerUploadRequest) doPost() error {
	ctx := inst.gc
	ups := inst.controller.UploadService
	items1 := inst.body1.Items
	items2 := make([]*dto.Upload, 0)
	for _, item1 := range items1 {
		item2, err := ups.BeginUpload(ctx, item1)
		if err != nil {
			return err
		}
		items2 = append(items2, item2)
	}
	inst.body2.Items = items2
	return nil
}

func (inst *innerUploadRequest) doPut() error {
	ctx := inst.gc
	ups := inst.controller.UploadService
	items1 := inst.body1.Items
	items2 := make([]*dto.Upload, 0)
	for _, item1 := range items1 {
		item2, err := ups.EndUpload(ctx, item1.ID, item1)
		if err != nil {
			return err
		}
		items2 = append(items2, item2)
	}
	inst.body2.Items = items2
	return nil
}
