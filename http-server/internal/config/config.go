package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type httpConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type dbConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Basename string `yaml:"basename"`
}

type Config struct {
	Http httpConfig `yaml:"http"`
	Db   dbConfig   `yaml:"db"`
}

func Load() (*Config, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("read config.yaml: %w", err)
	}

	cfg := &Config{}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("unmarshal yaml: %w", err)
	}

	return cfg, nil
}

func (c *Config) GetAddr() string {
	return fmt.Sprintf("%s:%d", c.Http.Host, c.Http.Port)
}

func (c *Config) GetPgUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		c.Db.Username,
		c.Db.Password,
		c.Db.Host,
		c.Db.Port,
		c.Db.Basename,
	)
}
