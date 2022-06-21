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

	setConfigFile(v, file)

	v.SetDefault("store.path", "list.txt")

	err := v.ReadInConfig()
	err = createFileIfNeeded(v, err, file)

	if err != nil {
		return nil, err
	}

	return &Config{v: v}, nil
}

func setConfigFile(v *viper.Viper, file string) {
	ext := filepath.Ext(file)

	v.AddConfigPath(filepath.Dir(file))
	v.SetConfigName(strings.TrimSuffix(filepath.Base(file), ext))
	v.SetConfigType(ext[1:])
}

func createFileIfNeeded(v *viper.Viper, err error, file string) error {
	if err == nil {
		return nil
	}

	if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		return err
	}

	err = os.MkdirAll(filepath.Dir(file), os.ModePerm)
	if err != nil {
		return err
	}

	err = v.SafeWriteConfigAs(file)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) SetStorePath(path string) {
	c.v.Set("store.path", path)
}

func (c *Config) Write() error {
	return c.v.WriteConfig()
}

func (c *Config) GetStorePath() string {
	return c.v.GetString("store.path")
}
