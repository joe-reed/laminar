package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	v *viper.Viper
}

func New() *Config {
	return &Config{v: viper.New()}
}

func Load(file string) (*Config, error) {
	v := viper.New()

	ext := filepath.Ext(file)

	v.AddConfigPath(filepath.Dir(file))
	v.SetConfigName(strings.TrimSuffix(filepath.Base(file), ext))
	v.SetConfigType(ext[1:])

	v.SetDefault("store.path", "list.txt")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err := os.MkdirAll(filepath.Dir(file), os.ModePerm)
			if err != nil {
				return nil, err
			}

			err = v.SafeWriteConfigAs(file)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return &Config{v: v}, nil
}

func (c *Config) SetStorePath(path string) error {
	c.v.Set("store.path", path)
	return c.v.WriteConfig()
}

func (c *Config) GetStorePath() string {
	return c.v.GetString("store.path")
}
