package config

import (
	"os"
	"reblog-server/utils"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host                  string
		Port                  int
		User                  string
		Password              string
		Name                  string
		MaxConnection         int
		MinConnection         int
		MaxLifeTimeConnection float64
		MaxIdleTimeConnection float64
	}

	API struct {
		Port int
	}

	Controller struct {
		HashCost int
	}

	Auth struct {
		JWTSecret string
	}
}

var (
	App *Config
)

func InitConfig() {
	var err error

	currentWD, err := os.Getwd()
	utils.Info("currentWD: %v", currentWD)
	if err != nil {
		utils.Error("[config.InitConfig] cannot get current working dir. err: %v", err)
		currentWD = "."
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(currentWD)
	err = viper.ReadInConfig()
	if err != nil {
		utils.Panic("[config.InitConfig] cannot get config file. err: %v", err)
	}

	App = &Config{}
	App.Database.Host = viper.GetString("rb_db_host")
	App.Database.Port = viper.GetInt("rb_db_port")
	App.Database.User = viper.GetString("rb_db_user")
	App.Database.Password = viper.GetString("rb_db_password")
	App.Database.Name = viper.GetString("rb_db_name")
	App.Database.MaxConnection = viper.GetInt("rb_DB_MAXCONN")
	App.Database.MinConnection = viper.GetInt("rb_DB_MINCONN")
	App.Database.MaxLifeTimeConnection = viper.GetFloat64("rb_DB_CONN_LIFETIME")
	App.Database.MaxIdleTimeConnection = viper.GetFloat64("rb_DB_CONN_IDLETIME")

	App.API.Port = viper.GetInt("rb_API_PORT")

	App.Controller.HashCost = viper.GetInt("rb_APP_HASHING_COST")

	App.Auth.JWTSecret = viper.GetString("rb_APP_JWT_SECRET")
}
