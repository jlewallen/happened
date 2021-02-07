package main

import (
	_ "fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

func startup() error {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigName("happened.json")
	viper.AddConfigPath("$HOME/")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("HPN")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	if err := viper.WriteConfigAs("happened-template.json"); err != nil {
		return err
	}

	return nil
}

func main() {
	log.Printf("starting")

	if err := startup(); err != nil {
		panic(err)
	}

	sm := NewStreamManager()

	defer sm.Close()

	go listen(sm)

	web(sm)
}
