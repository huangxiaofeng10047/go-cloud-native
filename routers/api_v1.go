package routers

import (
	"github.com/CodeLine-95/go-cloud-native/middleware/xlog"
	"github.com/CodeLine-95/go-cloud-native/routers/docker"
	"github.com/gin-gonic/gin"
)

// ApiV1 group: v1
func ApiV1(r *gin.Engine) *gin.Engine {
	r.Use(xlog.Logger())
	v1 := r.Group("/v1")
	docker.RouterDocker(v1)
	return r
}