package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const ConfPath = "/etc/confd/demo.json"

type Config struct {
	DatabaseUrl  string `json:"database_url" mapstructure:"database_url"`
	DatabaseUser string `json:"database_user" mapstructure:"database_user"`
}

func main() {
	//viperReadConfigFile()
	viperToWatchConfigFile()
	<-make(chan struct{})
}

func viperToWatchConfigFile() {
	viper.SetConfigFile(ConfPath)
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		var config Config
		if err := viper.Unmarshal(&config); err != nil {
			fmt.Println("Failed to unmarshal config:", err)
		} else {
			fmt.Println("Updated config:", config)
		}
	})

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Failed to read config:", err)
		return
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("Failed to unmarshal config:", err)
		return
	}

	fmt.Println("Initial config:", config)
}
