package main

import (
	"fmt"
	"os"

	"github.com/deuxksy/jasmine/configuration"
	"github.com/fsnotify/fsnotify"
	"github.com/gocolly/colly/v2"
	"github.com/spf13/viper"
)

// main 패키지에 init() 메서드를 만들어놓으면
// main문이 실행되기전에 먼저 실행됩니다.
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
		fmt.Println("Config file changed:", e.Name)
		var err error
		err = viper.ReadInConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
		err = viper.Unmarshal(&configuration.RuntimeConf)
		if err != nil {
			fmt.Println(err)
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
	fmt.Println("GO_PROFILE: " + profile)
	return profile
}

func collyTest() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("http://go-colly.org/")
}

func main() {
	// 어디서든 가져다 쓸 수 있습니다.
	fmt.Println("db type: ", configuration.RuntimeConf.Datasource.DbType)
}