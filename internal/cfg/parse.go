package cfg

import "github.com/spf13/viper"

func ParseConfig() error {
	viper.SetConfigFile(".cfg/local_conf.yaml")
	err := viper.ReadInConfig()
	return err
}
