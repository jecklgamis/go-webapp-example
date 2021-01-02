package server

import (
	"fmt"
	"log"
	"os"
)
import "github.com/spf13/viper"

// HTTPServerConfig stores HTTP server specific config
type HTTPServerConfig struct {
	Port int
}

// HTTPSServerConfig stores HTTPS server specific config
type HTTPSServerConfig struct {
	Port     int
	KeyFile  string
	CertFile string
}

// ListenerConfig stores HTTP or HTTPS server config
type ListenerConfig struct {
	HTTP  *HTTPServerConfig
	HTTPS *HTTPSServerConfig
}

// Config store the main application configuration
type Config struct {
	Server   *ListenerConfig
	Metadata map[string]string
}

// ReadConfig loads the environment-specific configuration in config directory
func ReadConfig(env string) *Config {
	configFile := fmt.Sprintf("config-%s.yaml", env)
	log.Printf("Loading %s\n", configFile)
	viper.SetConfigName(configFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Unable to read config file: %s\n", err)
	}
	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		log.Fatalf("Unable to umarshall config file: %s\n", err)
	}
	return config
}

// GetEnvOrElse retrieves the value of environment variable or fallback if it does not exist
func GetEnvOrElse(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
