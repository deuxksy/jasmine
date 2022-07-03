package main

import (
	"os"

	"github.com/deuxksy/template-go-application/internal/configuration"
	"github.com/deuxksy/template-go-application/internal/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/samber/lo"
	"github.com/spf13/viper"
)


func init() {
	profile := initProfile()
	setRuntimeConfig(profile)
}

func setRuntimeConfig(profile string) {
	viper.AddConfigPath("configs")
	viper.SetConfigName(profile)
	viper.SetConfigType("yaml")
	viper.Set("Verbose", true)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&configuration.RuntimeConf)
	if err != nil {
		panic(err)
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		logger.Warn("Config file changed: %s", e.Name)
		var err error
		err = viper.ReadInConfig()
		if err != nil {
			logger.Error(err.Error())
			return
		}
		err = viper.Unmarshal(&configuration.RuntimeConf)
		if err != nil {
			logger.Error(err.Error())
			return
		}
	})
	viper.WatchConfig()
}

func initProfile() string {
	var profile string
	profile = os.Getenv("GO_PROFILE")
	if len(profile) <= 0 {
		profile = "local"
	}
	logger.Debug(profile)
	return profile
}

func main () {
	names := lo.Uniq([]string{"Samuel", "Marc", "Samuel"})
	logger.Info("%s, %d", names, len(names))
	logger.Info("%d", configuration.RuntimeConf.Server.Port)
	logger.Info(configuration.RuntimeConf.Datasource.Url)
	logger.Error(configuration.RuntimeConf.Datasource.DbType)
}