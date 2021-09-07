package core

import (
	"github.com/bpcoder16/Mario/log"
	"github.com/bpcoder16/Mario/mario"
)

func SetPanicLogger() {
	switch mario.Server.Log.PanicLog.LogType {
	case "rotatelogs":
		mario.PanicIOWriter = log.GetRotateLogWriter(mario.Server.Log.PanicLog)
	default:
		mario.PanicIOWriter = log.GetRotateLogWriter(mario.Server.Log.PanicLog)
	}
}

func SetZapLogger() {
	log.SetZapDefaultLogger()
	defer log.DeferZapSync()
}
