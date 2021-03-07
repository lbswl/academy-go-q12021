package config

import (
	"github.com/tkanos/gonfig"
)

// Configuration struct contains variables
type Configuration struct {
	CsvPath string
}

// GetConfig returns variables for configuration
func GetConfig() Configuration {
	configuration := Configuration{}

	fileName := "config.json"

	gonfig.GetConf(fileName, &configuration)
	return configuration

}
