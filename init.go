package Mario

import (
	"github.com/bpcoder16/Mario/config"
	"github.com/bpcoder16/Mario/core"
	"github.com/bpcoder16/Mario/mario"
	"github.com/bpcoder16/Mario/utils"
)

func init() {
	initServerConfig()

	initLogger()
}

func initServerConfig() {
	var server config.Server
	utils.SetConfigWithFile("./config/server.toml", &server)
	mario.Server = &server
}

func initLogger() {
	core.SetPanicLogger()
	core.SetZapLogger()
}
