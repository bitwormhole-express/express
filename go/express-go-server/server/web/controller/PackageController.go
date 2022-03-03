package controller

import (
	"net/http"

	"bitwomrhole.com/djaf/express-go-server/server/service"
	"bitwomrhole.com/djaf/express-go-server/server/web/dto"
	"bitwomrhole.com/djaf/express-go-server/server/web/vo"
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

// PackageController 属于是包裹对象控制器
type PackageController struct {
	markup.Component `class:"rest-controller"`

	Responder      glass.MainResponder    `inject:"#glass-main-responder"`
	PackageService service.PackageService `inject:"#express-PackageService"`
}

func (inst *PackageController) _Impl() glass.Controller {
	return inst
}

func (inst *PackageController) Init(ec glass.EngineConnection) error {
	ec = ec.RequestMapping("packages")
	ec.Handle(http.MethodDelete, ":id", inst.doDelete)
	ec.Handle(http.MethodGet, "", inst.doGetList)
	ec.Handle(http.MethodGet, ":id", inst.doGetOne)
	ec.Handle(http.MethodPost, "", inst.doPost)
	ec.Handle(http.MethodPut, ":id", inst.doPut)
	return nil
}

func (inst *PackageController) doDelete(c *gin.Context) {
	req := innerPackageRequest{
		controller: inst,
		gc:         c,
	}
	req.wantID = true
	req.wantParams = false
	req.wantRequestBody = false
	err := req.open()
	if err == nil {
		err = req.doRemove()
	}
	req.send(err)
}

func (inst *PackageController) doGetList(c *gin.Context) {
	req := innerPackageRequest{
		controller: inst,
		gc:         c,
	}
	req.wantID = false
	req.wantParams = true
	req.wantRequestBody = false
	err := req.open()
	if err == nil {
		err = req.doFindList()
	}
	req.send(err)
}

func (inst *PackageController) doGetOne(c *gin.Context) {
	req := innerPackageRequest{
		controller: inst,
		gc:         c,
	}
	req.wantID = true
	req.wantParams = false
	req.wantRequestBody = false
	err := req.open()
	if err == nil {
		err = req.doFindOne()
	}
	req.send(err)
}

func (inst *PackageController) doPost(c *gin.Context) {
	req := innerPackageRequest{
		controller: inst,
		gc:         c,
	}
	req.wantID = false
	req.wantParams = false
	req.wantRequestBody = true
	err := req.open()
	if err == nil {
		err = req.doInsert()
	}
	req.send(err)
}

func (inst *PackageController) doPut(c *gin.Context) {
	req := innerPackageRequest{
		controller: inst,
		gc:         c,
	}
	req.wantID = true
	req.wantParams = false
	req.wantRequestBody = true
	err := req.open()
	if err == nil {
		err = req.doUpdate()
	}
	req.send(err)
}

///////////////////////////////////////////////////////////////////////////////

type innerPackageRequest struct {
	controller *PackageController
	gc         *gin.Context

	wantID          bool
	wantParams      bool
	wantRequestBody bool

	id string
	rx vo.Package
	tx vo.Package
}

func (inst *innerPackageRequest) doFindList() error {
	rxlist := inst.rx.Packages
	o1 := rxlist[0]
	o2list, err := inst.controller.PackageService.Find(o1)
	if err != nil {
		return err
	}
	inst.tx.Packages = o2list
	return nil
}

func (inst *innerPackageRequest) doFindOne() error {
	id := dto.PackageID(inst.id)
	o2, err := inst.controller.PackageService.FindOne(id)
	if err != nil {
		return err
	}
	inst.tx.Packages = []*dto.Package{o2}
	return nil
}

func (inst *innerPackageRequest) doInsert() error {
	rxlist := inst.rx.Packages
	o1 := rxlist[0]
	_, err := inst.controller.PackageService.Insert(o1)
	return err
}

func (inst *innerPackageRequest) doRemove() error {
	id := dto.PackageID(inst.id)
	return inst.controller.PackageService.Remove(id)
}

func (inst *innerPackageRequest) doUpdate() error {
	id := dto.PackageID(inst.id)
	rxlist := inst.rx.Packages
	o1 := rxlist[0]
	_, err := inst.controller.PackageService.Update(id, o1)
	return err
}

func (inst *innerPackageRequest) open() error {

	gc := inst.gc

	if inst.wantID {
		inst.id = gc.Param("id")
	}

	if inst.wantParams {
		pack := &dto.Package{}
		pack.FromAddress = gc.Param("from")
		pack.ToAddress = gc.Param("to")
		inst.rx.Packages = []*dto.Package{pack}
	}

	if inst.wantRequestBody {
		gc.BindJSON(&inst.rx)
	}

	return nil
}

func (inst *innerPackageRequest) send(err error) {
	data := &inst.tx
	status := inst.tx.Status
	resp := &glass.Response{
		Context: inst.gc,
		Data:    data,
		Error:   err,
		Status:  status,
	}
	inst.controller.Responder.Send(resp)
}
