package middleware

import (
	"github.com/CodeLine-95/go-cloud-native/pkg/logz"
	"github.com/CodeLine-95/go-cloud-native/pkg/utils/traceId"
	"github.com/CodeLine-95/go-cloud-native/pkg/xlog"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				trace := debug.Stack()
				xlog.Error(traceId.GetLogContext(c, "panic", logz.F("err", err), logz.F("stack", string(trace))))
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
