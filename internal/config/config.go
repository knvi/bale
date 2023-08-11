package config

import (
	"github.com/spf13/viper"
)

type Template struct {
	Path string
	Name string
}

type Config struct {
	Templates []Template
	viper *viper.Viper
}

func NewConfig() (*Config, error) {
	viper := viper.New()
	viper.SetConfigName("bale")
	viper.AddConfigPath("$HOME/.bale")
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{
		Templates: []Template{},
		viper: viper,
	}, nil
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
	return c.viper.WriteConfig()
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("bale")
	viper.AddConfigPath("$HOME/.bale")
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}