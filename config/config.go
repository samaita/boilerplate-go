package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Port string `mapstructure:"PORT"`
	} `mapstructure:"APP"`
	Datastore struct {
		Database struct {
			Postgres struct {
				DBName            string        `mapstructure:"DBNAME"`
				Host              string        `mapstructure:"HOST"`
				MaxConnection     int           `mapstructure:"MAX_CONNECTION"`
				MaxIdleConnection int           `mapstructure:"MAX_IDLE_CONNECTION"`
				Password          string        `mapstructure:"PASSWORD"`
				Port              string        `mapstructure:"PORT"`
				SSLMode           string        `mapstructure:"SSLMODE"`
				User              string        `mapstructure:"USER"`
				Timeout           time.Duration `mapstructure:"TIMEOUT"`
			} `mapstructure:"POSTGRES"`
		} `mapstructure:"DATABASE"`
	} `mapstructure:"DATASTORE"`
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
