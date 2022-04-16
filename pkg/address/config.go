package address

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

//Available config variables
type Config struct {
	DbHost     string
	DbPort     string
	DbUserName string
	DbPassword string
	DbDatabase string
	DbEngine   string
	ConfigType string
	ConfigFile string
}

//Initialize config variables
func InitConfig(c *Config) {
	dir, _ := os.Getwd()
	viper.AddConfigPath(dir)
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.ReadInConfig()
	if c.ConfigType != "" && c.ConfigFile != "" {
		viper.SetConfigType(c.ConfigType)
		viper.SetConfigName(c.ConfigFile)
		viper.MergeInConfig()
	}
	analyzeConfig()
}

//Get config by key
func GetConfig(key string) string {
	return viper.GetString(key)
}

//analyze config variables and merged it
func analyzeConfig() {
	configs := viper.AllKeys()
	for _, config := range configs {
		if strings.Contains(config, "db") {
			switch {
			case strings.Contains(config, "host"):
				viper.SetDefault("DB_HOST", viper.GetString(config))
			case strings.Contains(config, "port"):
				viper.SetDefault("DB_PORT", viper.GetString(config))
			case strings.Contains(config, "user"):
				viper.SetDefault("DB_USERNAME", viper.GetString(config))
			case strings.Contains(config, "password"):
				viper.SetDefault("DB_PASSWORD", viper.GetString(config))
			case strings.Contains(config, "database") || strings.Contains(config, "name") || strings.Contains(config, "db"):
				viper.SetDefault("DB_DATABASE", viper.GetString(config))
			}
		}
	}
}
