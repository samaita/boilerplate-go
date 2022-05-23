package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
		Cache struct {
			Redis struct {
				DB       int           `mapstructure:"DB"`
				Host     string        `mapstructure:"HOST"`
				Password string        `mapstructure:"PASSWORD"`
				Port     string        `mapstructure:"PORT"`
				Timeout  time.Duration `mapstructure:"TIMEOUT"`
			} `mapstructure:"REDIS"`
		} `mapstructure:"CACHE"`
	} `mapstructure:"DATASTORE"`
}

// GetConfig initialize the config by load config from ENV and/or .env file
func GetConfig() (conf Config) {
	loadConfigFromEnv()
	conf = loadConfigFromFile(".", ".env")
	return
}

func loadConfigFromFile(path, name string) (config Config) {
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

func loadConfigFromEnv() {
	envKeys := []string{
		"APP.PORT",
		"DATASTORE.DATABASE.POSTGRES.DBNAME",
		"DATASTORE.DATABASE.POSTGRES.HOST",
		"DATASTORE.DATABASE.POSTGRES.MAX_CONNECTION",
		"DATASTORE.DATABASE.POSTGRES.MAX_IDLE_CONNECTION",
		"DATASTORE.DATABASE.POSTGRES.PASSWORD",
		"DATASTORE.DATABASE.POSTGRES.PORT",
		"DATASTORE.DATABASE.POSTGRES.SSLMODE",
		"DATASTORE.DATABASE.POSTGRES.USER",
		"DATASTORE.DATABASE.POSTGRES.TIMEOUT",
		"DATASTORE.CACHE.REDIS.DB",
		"DATASTORE.CACHE.REDIS.HOST",
		"DATASTORE.CACHE.REDIS.PASSWORD",
		"DATASTORE.CACHE.REDIS.TIMEOUT",
	}

	for i, key := range envKeys {
		envKeys[i] = fmt.Sprintf(`%s="%s"`, envKeys[i], os.Getenv(key))
	}
	createConfigFile(envKeys)
}

func createConfigFile(lines []string) {
	f, err := os.OpenFile(".env", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}

	fileContent := ""
	for _, line := range lines {
		fileContent += line
		fileContent += "\n"
	}

	if err = ioutil.WriteFile(".env", []byte(fileContent), 0644); err != nil {
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
