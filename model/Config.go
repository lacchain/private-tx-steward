package model

type ApplicationConfig struct {
	NodeURL string `mapstructure:"nodeURL"`
	Port string `mapstructure:"port"`
}

type Config struct {
	Application  ApplicationConfig `mapstructure:"application"`
}