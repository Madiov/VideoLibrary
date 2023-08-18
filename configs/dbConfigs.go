package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Configs struct {
	Driver      string `mapstructure:"DB_DRIVER"`
	User        string `mapstructure:"DB_USER"`
	Password    string `mapstructure:"DB_PASSWORD"`
	IP          string `mapstructure:"DB_IP"`
	Port        uint   `mapstructure:"DB_PORT"`
	Name        string `mapstructure:"DB_NAME"`
	HostAddress string `mapstructure:"HOST_ADDRESS"`
	HostPort    string `mapstructure:"HOST_PORT"`
}

func LoadDbConfigs(path string) (Configs, error) {
	var conf Configs
	viper.AddConfigPath(path)
	viper.SetConfigName("conf")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	err = viper.Unmarshal(&conf)
	return conf, err
}
