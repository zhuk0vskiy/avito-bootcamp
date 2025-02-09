package config

import (
	// "flag"
	// "os"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	// "cleanenv"
)

type Config struct {
	Env      string        `yaml:"env" env-default:"local"`
	DB       DB            `yaml:"db"`
	SSO      SSO           `yaml:"sso"`
	Logger   Logger        `yaml:"logger"`
	TokenTTL time.Duration `yaml:"tokenTtl" env-default:"1h"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Driver   string `yaml:"driver"`
}

type SSO struct {
	GRPC GRPC `yaml:"grpc"`
}

type Sqlite struct {
	StoragePath    string `yaml:"storagePath" env-required:"true"`
	MigrationPath  string `yaml:"migrationPath"`
	MigrationTable string `yaml:"migrationTable"`
}

type DB struct {
	Postgres Postgres `yaml:"postgres"`
	Sqlite   Sqlite   `yaml:"sqlite"`
	ExecRetryNumber int `yaml:"execRetryNumber"`
	SleepTimeMilsec time.Duration `yaml:"sleepTimeMilsec"`
}

type GRPC struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type Logger struct {
	Level string `yaml:"level"`
	File  string `yaml:"file"`
}

func New() (config *Config, err error) {

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file %s does not exist", configPath)
	}

	viper.SetConfigFile(configPath)

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, err
}
