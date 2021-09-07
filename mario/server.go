package mario

import (
	"github.com/bpcoder16/Mario/config"
	"go.uber.org/zap"
	"io"
)

var (
	Server *config.Server

	PanicIOWriter io.Writer

	ZapLogger        *zap.Logger
	ZapSugaredLogger *zap.SugaredLogger
)
