package config

type Log struct {
	PanicLog LogItem
	InfoLog  LogItem
	ErrorLog LogItem
}

type LogItem struct {
	LogType string
	Config  struct {
		Filename      string
		MaxSaveDayNum int
		EncoderType   string
	}
}
