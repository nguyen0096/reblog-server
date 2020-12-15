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

// NewConfig returns struct of application configurations
func NewConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	AppConfig := &Config{}
	AppConfig.Database.Host = viper.GetString("rb_db_host")
	AppConfig.Database.Port = viper.GetInt("rb_db_port")
	AppConfig.Database.User = viper.GetString("rb_db_user")
	AppConfig.Database.Password = viper.GetString("rb_db_password")
	AppConfig.Database.Name = viper.GetString("rb_db_name")
	AppConfig.Database.MaxConnection = viper.GetInt("RB_DB_MAXCONN")
	AppConfig.Database.MinConnection = viper.GetInt("rb_DB_MINCONN")
	AppConfig.Database.MaxLifeTimeConnection = viper.GetFloat64("rb_DB_CONN_LIFETIME")
	AppConfig.Database.MaxIdleTimeConnection = viper.GetFloat64("rb_DB_CONN_IDLETIME")

	return AppConfig
}

func (c *Config) GetHostname() string {
	return c.Database.Host
}

func (c *Config) GetPort() int {
	return c.Database.Port
}

func (c *Config) GetUser() string {
	return c.Database.User
}

func (c *Config) GetPassword() string {
	return c.Database.Password
}

func (c *Config) GetDatabase() string {
	return c.Database.Name
}
