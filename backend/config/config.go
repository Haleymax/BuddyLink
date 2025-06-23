package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Server ServerConfig `mapstructure:"server"`
	MySQL  MySQLConfig  `mapstructure:"mysql"`
	Minio  MinioConfig  `mapstructure:"minio"`
	SMTP   SMTPConfig   `mapstructure:"smtp"`
}

func LoadConfig() Config {
	workDir, _ := os.Getwd()
	configFile := filepath.Join(workDir, "config", "config.yaml")
	log.Println("Loading config from ", configFile)

	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	log.Println("Using config file:", viper.AllSettings())

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshalling config, %s", err)
	}
	return config
}
