package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Service represents a single service configuration.
type Service struct {
	Name           string                 `yaml:"name"`
	AllowedMethods []string               `yaml:"allowed_methods"`
	URLs           []string               `yaml:"urls"` // UPDATED: Supports multiple URLs for load balancing
	Prefix         string                 `yaml:"prefix"`
	Filter         map[string]interface{} `yaml:"filter"`
}

// ServiceLeader holds high-level configuration about the API Gateway.
type ServiceLeader struct {
	Name        string                 `yaml:"name"`
	Version     string                 `yaml:"version"`
	Environment string                 `yaml:"environment"`
	Description string                 `yaml:"description"`
	Filter      map[string]interface{} `yaml:"filter"`   // FIXED: Use map instead of slice
	Services    []Service              `yaml:"services"` // FIXED: Now supports multiple services correctly
	Logging     Logging                `yaml:"logging"`
	Env         []EnvVar               `yaml:"env"`
}

// Logging holds logging configuration.
type Logging struct {
	LogRequestBody  bool   `yaml:"log_request_body"`
	LogResponseBody bool   `yaml:"log_response_body"`
	MinLevel        string `yaml:"min_level"`
}

// EnvVar represents an environment variable configuration.
type EnvVar struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

// Config is the top-level configuration structure.
type Config struct {
	ServiceLeader ServiceLeader `yaml:"service_leader"` // FIXED: Match YAML key
}

// GetEnvValue retrieves an environment variable's value from the config.
func GetEnvValue(envs []EnvVar, key, defaultValue string) string {
	for _, env := range envs {
		if env.Name == key {
			return env.Value
		}
	}
	return defaultValue
}

// LoadConfig loads and parses the YAML config file.
func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error opening config file: ", err)
		return nil, err
	}
	defer file.Close()

	var config Config
	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Fatal("Error decoding YAML: ", err)
		return nil, err
	}

	return &config, nil
}
