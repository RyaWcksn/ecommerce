package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		ENV      string `mapstructure:"env"`
		APPNAME  string `mapstructure:"app_name"`
		LOGLEVEL string `mapstructure:"log_level"`
	} `mapstructure:"app"`
	Server struct {
		HTTPAddress string `mapstructure:"httpaddress"`
	} `mapstructure:"server"`
	Database struct {
		Host        string `mapstructure:"host"`
		Username    string `mapstructure:"username"`
		Password    string `mapstructure:"password"`
		Database    string `mapstructure:"database"`
		MaxIdleConn int    `mapstructure:"maxiddle"`
		MaxOpenConn int    `mapstructure:"maxopen"`
	} `mapstructure:"database"`
}

var Cfg *Config

func init() {
	Cfg = LoadConfig()
}

func LoadConfig() *Config {
	// Initialize viper
	cfg := &Config{}

	viper.SetConfigName("config")     // name of the config file (without extension)
	viper.AddConfigPath(".")          // search the root directory for the config file
	viper.AddConfigPath("./configs/") // search the root directory for the config file
	viper.SetConfigType("yaml")       // use YAML config format

	// Load the config file
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config file: %s", err))
	}

	// Unmarshal the config into a struct
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(fmt.Errorf("failed to unmarshal config: %s", err))
	}

	// Use the config values
	return cfg
}
