package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Configurations exported
type Configurations struct {
	Server      ServerConfig
	MailJet     MailJetConfig
	Me          MeConfig
	HegdeFlutes HegdeFlutesConfig
}

//MeConfig Exported
type MeConfig struct {
	Email string
}

//HegdeFlutes Exported
type HegdeFlutesConfig struct {
	Email string
}

// ServerConfig exported
type ServerConfig struct {
	Port int
}

// MailJetConfig exported
type MailJetConfig struct {
	Templates struct {
		ContactMe            int
		HegdeFlutesContactUs int
	}
	ApiKey struct {
		Public  string
		Private string
	}
}

var (
	bindMap = map[string]string{
		"server.port": "PORT",
	}
)

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Configurations, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	//bind env to struct
	for k, v := range bindMap {
		viper.BindEnv(k, v)
	}

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
	fmt.Println(config)
	fmt.Println(os.Getenv("MAILJET_APIKEY_PUBLIC"))
	return
}
