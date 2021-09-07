package log

import (
	"github.com/bpcoder16/Mario/config"
	"github.com/bpcoder16/Mario/mario"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func getRotateLogSyncWriter(logConfig config.LogItem) zapcore.WriteSyncer {
	return zapcore.AddSync(GetRotateLogWriter(logConfig))
}

func getSyncLogWriter(logConfig config.LogItem) zapcore.WriteSyncer {
	switch logConfig.LogType {
	case "rotatelogs":
		return getRotateLogSyncWriter(logConfig)
	}
	return getRotateLogSyncWriter(logConfig)
}

func getEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
}

func getEncoder(encoderType string) zapcore.Encoder {
	switch encoderType {
	case "JSONEncoder":
		return zapcore.NewJSONEncoder(getEncoderConfig())
	case "ConsoleEncoder":
		return zapcore.NewConsoleEncoder(getEncoderConfig())
	}
	return zapcore.NewJSONEncoder(getEncoderConfig())
}

func SetZapDefaultLogger() {
	var cores []zapcore.Core

	infoLogWriter := getSyncLogWriter(mario.Server.Log.InfoLog)
	infoEncoder := getEncoder(mario.Server.Log.InfoLog.Config.EncoderType)
	errorLogWriter := getSyncLogWriter(mario.Server.Log.ErrorLog)
	errorEncoder := getEncoder(mario.Server.Log.ErrorLog.Config.EncoderType)

	var highPriority = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	var lowPriority = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		if mario.Server.System.RunMode == "release" {
			return lvl == zapcore.InfoLevel
		}
		return lvl < zapcore.WarnLevel
	})

	if mario.Server.System.RunMode == "dev" {
		cores = append(cores, zapcore.NewCore(infoEncoder, zapcore.NewMultiWriteSyncer(
			infoLogWriter,
			zapcore.AddSync(os.Stdout),
		), lowPriority))
		cores = append(cores, zapcore.NewCore(errorEncoder, zapcore.NewMultiWriteSyncer(
			errorLogWriter,
			zapcore.AddSync(os.Stdout),
		), highPriority))
	} else {
		cores = append(cores, zapcore.NewCore(infoEncoder, infoLogWriter, lowPriority))
		cores = append(cores, zapcore.NewCore(errorEncoder, errorLogWriter, highPriority))
	}

	core := zapcore.NewTee(cores...)

	mario.ZapLogger = zap.New(core, zap.AddCaller())
	mario.ZapSugaredLogger = mario.ZapLogger.Sugar()
}

func DeferZapSync() {
	_ = mario.ZapSugaredLogger.Sync()
	_ = mario.ZapLogger.Sync()
}
