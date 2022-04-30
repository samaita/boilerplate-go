package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Port string `mapstructure:"PORT"`
	} `mapstructure:"APP"`
}

// GetConfig initialize the config by load a config from a path
func GetConfig() (conf Config) {
	conf = loadConfig(".", ".env")
	return
}

func loadConfig(path, name string) (config Config) {
	var err error

	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		log.Fatalln("Err on read config:", err)
	}

	if err = viper.Unmarshal(&config); err != nil {
		log.Fatalln("Err on unmarshal config:", err)
	}
	return
}
