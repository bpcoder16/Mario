package middleware

import (
	"github.com/bpcoder16/Mario/core"
	"github.com/bpcoder16/Mario/mario"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
)

func DebugLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method == "POST" && mario.Server.System.RunMode == "dev" {
			body, err := core.GetBodyClone(ctx)
			if err == nil {
				rawData, err := ioutil.ReadAll(body)
				if err == nil {
					mario.ZapLogger.Debug("requestBody", zap.ByteString("requestBody", rawData))
				}
			}
		}

		// Process request
		ctx.Next()
	}
}
