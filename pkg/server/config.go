package server

import (
	"fmt"
	"log"
	"os"
)
import "github.com/spf13/viper"

type HttpServerConfig struct {
	Port int
}

type HttpsServerConfig struct {
	Port     int
	KeyFile  string
	CertFile string
}

type ServerConfig struct {
	Http  *HttpServerConfig
	Https *HttpsServerConfig
}

type Config struct {
	Server *ServerConfig
}

func ReadConfig(env string) *Config {
	configFile := fmt.Sprintf("config-%s.yaml", env)
	log.Printf("Loading %s\n", configFile)
	viper.SetConfigName(configFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Unable to read config file: %s\n", err)
	}
	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		log.Fatalf("Unable to umarshalle config file: %s\n", err)
	}
	return config
}

func GetEnvOrElse(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
