package Mario

import (
	"github.com/bpcoder16/Mario/core"
)

func RunHttpServer(manager core.RouterManager) {
	core.RunHttpServer(manager)
}

func RunMultiHttpServer(configList []core.HttpServerConfig) {
	core.RunMultiHttpServer(configList)
}
