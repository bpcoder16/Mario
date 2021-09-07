package middleware

import (
	"github.com/gin-gonic/gin"
)

func Log() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//if ctx.Request.Method == "POST" && mario.Server.System.RunMode == "dev" {
		//	body, err := utils.GetBodyClone(ctx)
		//	if err == nil {
		//		rawData, err := ioutil.ReadAll(body)
		//		if err == nil {
		//			mario.ZapSugaredLogger.With(string(rawData))
		//		}
		//	}
		//}

		// Process request
		ctx.Next()
	}
}
