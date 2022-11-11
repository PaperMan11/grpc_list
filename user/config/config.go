package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ServerConf Server `mapstructure:"server"`
	MysqlConf  Mysql  `mapstructure:"mysql"`
	EtcdConf   Etcd   `mapstructure:"etcd"`
}

type Server struct {
	GrpcAddress string `mapstructure:"grpcAddress"`
}

type Mysql struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DataBase string `mapstructure:"database"`
	UserName string `mapstructure:"username"`
	PassWord string `mapstructure:"password"`
}

type Etcd struct {
	Address string `mapstructure:"address"`
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
