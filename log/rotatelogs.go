package log

import (
	"github.com/bpcoder16/Mario/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"time"
)

func GetRotateLogWriter(logConfig config.LogItem) io.Writer {
	rotateLogP, err := rotatelogs.New(
		logConfig.Config.Filename+".%Y%m%d%H",
		rotatelogs.WithLinkName(logConfig.Config.Filename),                                     // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Duration(86400*logConfig.Config.MaxSaveDayNum)*time.Second), // 文件最大保存份数
		rotatelogs.WithRotationTime(time.Hour),                                                 // 日志切割时间间隔
	)
	if err != nil {
		panic(err)
	}
	return rotateLogP
}
