package config

import (
	"os"
	"strconv"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	PRODUCTION string = "prod"
	STAGING    string = "staging"
)

type MongoDbConfig struct {
	Driver   string `mapstructure:"Driver"`
	Name     string `mapstructure:"Name"`
	Host     string `mapstructure:"Host"`
	Port     int    `mapstructure:"Port"`
	Username string `mapstructure:"Username"`
	Password string `mapstructure:"Password"`
}

type AppConfig struct {
	Name      string `mapstructure:"Name"`
	Port      int    `mapstructure:"Port"`
	Secret    string `mapstructure:"Secret"`
	SystemEnv string `mapstructure:"SystemEnv"`
	//database
	MasterMongoDb MongoDbConfig `mapstructure:"MasterMongoDb"`
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfigs() *AppConfig {
	port := os.Getenv("PORT")
	if appConfig != nil {
		if port != "" {
			appConfig.Port, _ = strconv.Atoi(port)
		}
		return appConfig
	}

	lock.Lock()
	defer lock.Unlock()

	if appConfig != nil {
		if port != "" {
			appConfig.Port, _ = strconv.Atoi(port)
		}
		return appConfig
	}

	appConfig = initConfig()
	if port != "" {
		appConfig.Port, _ = strconv.Atoi(port)
	}
	return appConfig
}

func initConfig() *AppConfig {
	var finalConfig AppConfig
	viper.BindEnv("Name", "APP_NAME")
	viper.BindEnv("Port", "APP_PORT")
	viper.BindEnv("Secret", "APP_SECRET")
	viper.BindEnv("SystemEnv", "SYSTEM_ENV")
	viper.BindEnv("JwtExpired", "JWT_EXPIRED")
	//database
	viper.BindEnv("MasterMongoDb.Driver", "MASTER_MONGO_DB_DRIVER")
	viper.BindEnv("MasterMongoDb.Name", "MASTER_MONGO_DB_NAME")
	viper.BindEnv("MasterMongoDb.Host", "MASTER_MONGO_DB_HOST")
	viper.BindEnv("MasterMongoDb.Port", "MASTER_MONGO_DB_PORT")
	viper.BindEnv("MasterMongoDb.Username", "MASTER_MONGO_DB_USERNAME")
	viper.BindEnv("MasterMongoDb.Password", "MASTER_MONGO_DB_PASSWORD")
	err := viper.Unmarshal(&finalConfig)
	if err != nil {
		log.Info("failed to extract config")
	}
	return &finalConfig
}
