package utils

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func SetConfigWithFile(configPath string, configPtr interface{}) {
	var err error
	v := viper.New()
	v.SetConfigFile(configPath)
	err = v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		if err = v.Unmarshal(configPtr); err != nil {
			panic(err)
		}
	})
	if err = v.Unmarshal(configPtr); err != nil {
		panic(err)
	}
	return
}
