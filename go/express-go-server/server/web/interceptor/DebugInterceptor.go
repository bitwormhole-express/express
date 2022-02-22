package interceptor

import (
	"strings"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
	"github.com/gin-gonic/gin"
)

type DebugInterceptor struct {
	markup.Component `class:"rest-interceptor-registry"`
}

func (inst *DebugInterceptor) _Impl() (glass.InterceptorRegistry, glass.Interceptor) {
	return inst, inst
}

func (inst *DebugInterceptor) GetRegistrationList() []*glass.InterceptorRegistration {
	ir := &glass.InterceptorRegistration{}
	ir.Name = "debug-interceptor"
	ir.Order = 99
	ir.Interceptor = inst
	return []*glass.InterceptorRegistration{ir}
}

func (inst *DebugInterceptor) Intercept(h gin.HandlerFunc) gin.HandlerFunc {
	h2 := func(c *gin.Context) {
		inst.check(c)
		h(c)
	}
	return h2
}

func (inst *DebugInterceptor) check(c *gin.Context) {

	p1 := c.Request.URL.Path
	p2 := c.Request.URL.RawPath

	if !strings.HasPrefix(p1, "/api/") {
		return
	}

	vlog.Debug(p1, p2)
}
