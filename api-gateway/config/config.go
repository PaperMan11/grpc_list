package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type EtcdConf struct {
	Address string `mapstructure:"address"`
}

type ServerConf struct {
	Port int `mapstructure:"port"`
}

type Config struct {
	EtcdConf   EtcdConf   `mapstructure:"etcd"`
	ServerConf ServerConf `mapstructure:"server"`
}

var Conf = new(Config)

func Init() error {
	viper.SetConfigFile("../config/config.yml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Read config failed")
		return err
	}
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Println("viper.Unmarshal failed")
		return err
	}
	fmt.Println("Conf:", *Conf)
	return nil
}
