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
		Dsn string `yaml:"db_dsn" envconfig:"GOOSE_DBSTRING"`
	} `yaml:"database"`
}

func ProcessError(err error) {
	fmt.Println(err)
	os.Exit(2)
}
//получает настройки приложения из файла config.yml
func ReadFile(cfg *Config) {
	f, err := os.Open("/app/config.yml")
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
