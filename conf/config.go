package conf

import "github.com/spf13/viper"

type Config struct {
	NDServerConfig struct {
		URL      string `mapstructure:"url"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	} `mapstructure:"nd_server"`

	Parameters struct {
		AddTime        float64 `mapstructure:"add_time"`
		LastPlayTime   float64 `mapstructure:"last_play_time"`
		PlayBetweenDay int64   `mapstructure:"play_between_day"`
	} `mapstructure:"parameters"`
}

var Conf *Config

func LoadConfig() (*Config, error) {
	if Conf != nil {
		return Conf, nil
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	Conf = &Config{}
	err = viper.Unmarshal(Conf)
	if err != nil {
		return nil, err
	}
	return Conf, nil
}
