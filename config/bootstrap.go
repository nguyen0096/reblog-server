package config

import (
	"fmt"

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
}

var (
	AppConfig = &Config{}
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	AppConfig.Database.Host = viper.GetString("rb_db_host")
	AppConfig.Database.Port = viper.GetInt("rb_db_port")
	AppConfig.Database.User = viper.GetString("rb_db_user")
	AppConfig.Database.Password = viper.GetString("rb_db_password")
	AppConfig.Database.Name = viper.GetString("rb_db_name")
	AppConfig.Database.MaxConnection = viper.GetInt("RB_DB_MAXCONN")
	AppConfig.Database.MinConnection = viper.GetInt("rb_DB_MINCONN")
	AppConfig.Database.MaxLifeTimeConnection = viper.GetFloat64("rb_DB_CONN_LIFETIME")
	AppConfig.Database.MaxIdleTimeConnection = viper.GetFloat64("rb_DB_CONN_IDLETIME")
}
