package util

import "github.com/spf13/viper"

type Config struct {
	DataFile string `mapstructure:"DATA_FILE"`
	DataPath string `mapstructure:"DATA_PATH"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return

}
