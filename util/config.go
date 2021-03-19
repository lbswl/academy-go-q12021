package util

import "github.com/spf13/viper"

type Config struct {
	DataPath               string `mapstructure:"DATA_PATH"`
	DataFile               string `mapstructure:"DATA_FILE"`
	NumberCallsExternalApi int    `mapstructure:"NUMBER_CALLS_EXTERNAL_API"`
	UrlExternalApi         string `mapstructure:"URL_EXTERNAL_API"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("server")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return

}
