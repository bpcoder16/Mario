package core

import (
	"github.com/bpcoder16/Mario/mario"
	"github.com/bpcoder16/Mario/middleware"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/sync/errgroup"
)

type RouterManager func(r *gin.Engine)

func runGin(manager RouterManager) (r *gin.Engine) {
	binding.Validator = &MultiLangValidator{
		Locale:  "zh",
		TagName: "binding",
	}
	// Creates a router without any middleware by default
	r = gin.New()

	r.Use(middleware.DebugLog())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.RecoveryWithWriter(mario.PanicIOWriter))

	manager(r)

	return
}

func RunHttpServer(manager RouterManager) {
	r := runGin(manager)
	mario.ZapSugaredLogger.Fatal(r.Run(":" + mario.Server.System.Port))
}

type HttpServerConfig struct {
	Port    string
	Manager RouterManager
}

func RunMultiHttpServer(configList []HttpServerConfig) {
	var g errgroup.Group
	for _, config := range configList {
		r := runGin(config.Manager)
		addr := ":" + config.Port
		g.Go(func() (err error) {
			err = r.Run(addr)
			return
		})
	}

	// Wait for all HTTP fetches to complete.
	if err := g.Wait(); err != nil {
		mario.ZapSugaredLogger.Fatal(err)
	}
}
