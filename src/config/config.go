package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

// Configurations exported
type Configurations struct {
	Server  ServerConfig
	MailJet MailJetConfig
	Me      MeConfig
}

//MeConfig Exported
type MeConfig struct {
	Email string
}

// ServerConfig exported
type ServerConfig struct {
	Port int
}

// MailJetConfig exported
type MailJetConfig struct {
	Templates struct {
		ContactMe int
	}
	ApiKey struct {
		Public  string
		Private string
	}
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Configurations, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// for k, v := range defaults {
	// 	viper.SetDefault(k, v)
	// }

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("Config file not found %v", err)
		} else {
			return
		}
	}

	err = viper.Unmarshal(&config)
	return
}
