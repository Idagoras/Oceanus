package config

import (
	"github.com/spf13/viper"
	"time"
)

const (
	PORT           = ":8080"
	APP_NAME       = "bluebell"
	APP_SECRET     = "6YJSuc50uJ18zj45"
	API_EXPIRY     = "120"
	LOG_FILE_PATH  = "/var/log/bluebell"
	LOG_FILE_NAME  = "app.log"
	DB_DRIVER      = "postgres"
	DBSOURCE       = "postgresql://idagoras:314159@localhost:3306/adventure?sslmode=disable"
	SERVER_ADDRESS = "127.0.0.1"
)

type Config struct {
	Port                string        `mapstructure:"PORT"`
	AppName             string        `mapstructure:"APP_NAME"`
	AppSecret           string        `mapstructure:"APP_SECRET"`
	ApiExpiry           string        `mapstructure:"API_EXPIRY"`
	LogFilePath         string        `mapstructure:"LOG_FILE_PATH"`
	LogFileName         string        `mapstructure:"LOG_FILE_NAME"`
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DBSOURCE"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuartion time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	HttpServerAddress   string        `mapstructure:"HTTP_SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
