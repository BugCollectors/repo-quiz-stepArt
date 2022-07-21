package cfg

import "github.com/spf13/viper"

type Config struct {
	PGUsername string `mapstructure:"POSTGRES_USERNAME"`
	PGPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PGHost     string `mapstructure:"POSTGRES_HOST"`
	PGPort     string `mapstructure:"POSTGRES_PORT"`
	PGDBName   string `mapstructure:"POSTGRES_DBNAME"`
	PGSSLMode  string `mapstructure:"POSTGRES_SSLMODE"`
}

func LoadConfig() (*Config, error) {
	config := &Config{}

	vi := viper.New()
	vi.SetConfigFile(".cfg/local_conf.yaml")
	err := vi.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = vi.Unmarshal(&config)

	return config, err
}
