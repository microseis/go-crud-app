package utils

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	App struct {
		Port string `yaml:"port" envconfig:"APP_PORT"`
		Host string `yaml:"host" envconfig:"APP_HOST"`
	} `yaml:"app"`
	Database struct {
		Username string `yaml:"db_user" envconfig:"DB_USERNAME"`
		Password string `yaml:"db_password" envconfig:"DB_PASSWORD"`
		Port     string `yaml:"db_port" envconfig:"DB_PORT"`
		Host     string `yaml:"db_host" envconfig:"DB_HOST"`
		Name     string `yaml:"db_name" envconfig:"DB_NAME"`
	} `yaml:"database"`
}

func ProcessError(err error) {
	fmt.Println(err)
	os.Exit(2)
}
//получает настройки приложения из файла config.yml
func ReadFile(cfg *Config) {
	f, err := os.Open("./config.yml")
	if err != nil {
		ProcessError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		ProcessError(err)
	}
}
// считывает переменные окружения.
func ReadEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		ProcessError(err)
	}
}
