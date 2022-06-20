package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	V *viper.Viper
}

func New() *Config {
	return &Config{V: viper.New()}
}

func Load() (*Config, error) {
	v := viper.New()

	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(home, ".laminar")
	v.AddConfigPath(path)
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.SetDefault("store.path", "list.txt")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("config file not found")

			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return nil, err
			}

			err = v.SafeWriteConfigAs(filepath.Join(path, "config.yaml"))
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return &Config{V: v}, nil
}

func (c *Config) SetStorePath(path string) error {
	c.V.Set("store.path", path)
	return c.V.WriteConfig()
}

func (c *Config) GetStorePath() string {
	return c.V.GetString("store.path")
}
