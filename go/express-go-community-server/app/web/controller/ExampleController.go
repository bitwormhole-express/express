package controller

import (
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter-restful/api/vo"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

type ExampleController struct {
	markup.Component `class:"rest-controller"`
}

func (inst *ExampleController) _Impl() glass.Controller {
	return inst
}

func (inst *ExampleController) Init(ec glass.EngineConnection) error {
	ec = ec.RequestMapping("example")
	ec.Handle(http.MethodGet, "", inst.handle)
	ec.Handle(http.MethodPost, "", inst.handle)
	return nil
}

func (inst *ExampleController) handle(c *gin.Context) {
	o := &vo.Auth{}
	c.BindJSON(o)
	c.JSON(200, o)
}
