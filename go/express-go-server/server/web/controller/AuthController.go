package controller

import (
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter-restful/api/vo"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	markup.Component `class:"rest-controller"`
}

func (inst *AuthController) _Impl() glass.Controller {
	return inst
}

func (inst *AuthController) Init(ec glass.EngineConnection) error {
	ec = ec.RequestMapping("auth")
	ec.Handle(http.MethodGet, "", inst.handleGet)
	return nil
}

func (inst *AuthController) handleGet(c *gin.Context) {
	o := &vo.Auth{}
	c.JSON(http.StatusOK, o)
}
