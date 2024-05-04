package config

import "fmt"

type Config struct{}

func LoadConfig() (*Config, error) {
	fmt.Println("Loading config...")
	return &Config{}, nil
}
