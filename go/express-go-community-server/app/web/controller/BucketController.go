package controller

import (
	"errors"
	"net/http"

	"github.com/bitwomrhole-express/express/community-server/app/web/vo"
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

type BucketController struct {
	markup.Component `class:"rest-controller"`

	Responder glass.MainResponder `inject:"#glass-main-responder"`
}

func (inst *BucketController) _Impl() glass.Controller {
	return inst
}

func (inst *BucketController) Init(ec glass.EngineConnection) error {
	ec = ec.RequestMapping("buckets")

	ec.Handle(http.MethodGet, "", inst.handleGetList)
	ec.Handle(http.MethodPost, "", inst.handlePost)

	ec.Handle(http.MethodDelete, ":id", inst.handleDelete)
	ec.Handle(http.MethodGet, ":id", inst.handleGetOne)
	ec.Handle(http.MethodPut, ":id", inst.handlePut)
	return nil
}

func (inst *BucketController) handleGetOne(c *gin.Context) {
	req := innerBucketsRequest{
		controller: inst,
		gc:         c,
	}
	err := req.open()
	if err == nil {
		err = req.doGetOne()
	}
	req.send(err)
}

func (inst *BucketController) handleGetList(c *gin.Context) {
	req := innerBucketsRequest{
		controller: inst,
		gc:         c,
	}
	err := req.open()
	if err == nil {
		err = req.doGetList()
	}
	req.send(err)
}

func (inst *BucketController) handlePost(c *gin.Context) {
	req := innerBucketsRequest{
		controller: inst,
		gc:         c,
	}
	err := req.open()
	if err == nil {
		err = req.doPost()
	}
	req.send(err)
}

func (inst *BucketController) handlePut(c *gin.Context) {
	req := innerBucketsRequest{
		controller: inst,
		gc:         c,
	}
	err := req.open()
	if err == nil {
		err = req.doPut()
	}
	req.send(err)
}

func (inst *BucketController) handleDelete(c *gin.Context) {
	req := innerBucketsRequest{
		controller: inst,
		gc:         c,
	}
	err := req.open()
	if err == nil {
		err = req.doDelete()
	}
	req.send(err)
}

///////////////////////////////////////////////////////////////////////////////

type innerBucketsRequest struct {
	controller *BucketController
	gc         *gin.Context

	wantID          bool
	wantRequestBody bool

	id    string
	body1 vo.Bucket
	body2 vo.Bucket
}

func (inst *innerBucketsRequest) open() error {

	if inst.wantID {
		inst.id = inst.gc.Param("id")
	}

	if inst.wantRequestBody {
		inst.gc.BindJSON(&inst.body1)
	}

	return nil
}

func (inst *innerBucketsRequest) send(err error) {
	resp := &glass.Response{
		Context: inst.gc,
		Data:    &inst.body2,
		Error:   err,
		Status:  inst.body2.Status,
	}
	inst.controller.Responder.Send(resp)
}

func (inst *innerBucketsRequest) doPost() error {
	return errors.New("no impl")
}

func (inst *innerBucketsRequest) doPut() error {
	return errors.New("no impl")
}

func (inst *innerBucketsRequest) doDelete() error {
	return errors.New("no impl")
}

func (inst *innerBucketsRequest) doGetOne() error {
	return errors.New("no impl")
}

func (inst *innerBucketsRequest) doGetList() error {
	return errors.New("no impl")
}
