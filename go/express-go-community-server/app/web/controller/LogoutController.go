package controller

import (
	"context"
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter-restful/api/vo"
	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

type LogoutController struct {
	markup.Component `class:"rest-controller"`

	Subjects keeper.SubjectManager `inject:"#keeper-subject-manager"`
}

func (inst *LogoutController) _Impl() glass.Controller {
	return inst
}

func (inst *LogoutController) Init(ec glass.EngineConnection) error {
	ec = ec.RequestMapping("logout")
	ec.Handle(http.MethodPost, "", inst.handlePost)
	return nil
}

func (inst *LogoutController) handlePost(c *gin.Context) {

	subject, err := inst.Subjects.GetSubject(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = inst.doLogout(c, subject)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	status := http.StatusOK
	o := &vo.BaseVO{}
	o.Status = status
	c.JSON(status, o)
}

func (inst *LogoutController) doLogout(c context.Context, s keeper.Subject) error {
	session, err := s.GetSession(true)
	if err != nil {
		return err
	}
	tr := session.BeginTransaction()
	defer tr.Close()
	setter := session.Properties().Setter()
	setter.SetBool(keeper.SessionFieldAuthenticated, false)
	return tr.Commit()
}
