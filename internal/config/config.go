package config

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Template struct {
	Path string `toml:"path"`
	Name string `toml:"name"`
}

type Config struct {
	Templates []Template `toml:"templates"`
}

func NewConfig() *Config {
	return &Config{
		Templates: make([]Template, 0),
	}
}

func (c *Config) AddTemplate(template Template) {
	c.Templates = append(c.Templates, template)
}

func (c *Config) RemoveTemplate(template Template) {
	for i, t := range c.Templates {
		if t.Name == template.Name {
			c.Templates = append(c.Templates[:i], c.Templates[i+1:]...)
			break
		}
	}
}

func (c *Config) GetTemplate(name string) *Template {
	for _, t := range c.Templates {
		if t.Name == name {
			return &t
		}
	}

	return nil
}

func (c *Config) Save() error {
	path, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	cfgDir := filepath.Join(path, ".config", "bale")
	err = os.MkdirAll(cfgDir, os.ModePerm)
	if err != nil {
		return err
	}

	cfgPath := filepath.Join(cfgDir, "bale.toml")
	f, err := os.Create(cfgPath)
	if err != nil {
		return err
	}
	defer f.Close()

	err = toml.NewEncoder(f).Encode(c)
	if err != nil {
		return err
	}

	return nil
}

func LoadConfig() (*Config, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	cfgPath := filepath.Join(path, ".config", "bale", "bale.toml")
	_, err = os.Stat(cfgPath)
	if errors.Is(err, os.ErrNotExist) {
		return NewConfig(), nil
	} else if err != nil {
		return nil, err
	}

	var config Config
	_, err = toml.DecodeFile(cfgPath, &config)
	if err != nil {
		return nil, err
	}
	
	return &config, nil
}