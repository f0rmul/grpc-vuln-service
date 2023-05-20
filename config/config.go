package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	GrpcServer GrpcServer `yaml:"grpc_server"`
	Logger     Logger     `yaml:"logger"`
}

type GrpcServer struct {
	Host              string        `yaml:"host"`
	Port              string        `yaml:"port"`
	MaxConnectionIdle time.Duration `yaml:"max_connection_idle"`
	MaxConnectionAge  time.Duration `yaml:"max_connection_age"`
	Timeout           time.Duration `yaml:"timeout"`
	Time              time.Duration `yaml:"time"`
}

type Logger struct {
	Development       bool   `yaml:"development"`
	DisableCaller     bool   `yaml:"disable_caller"`
	DisableStacktrace bool   `yaml:"disable_stacktrace:"`
	Encoding          string `yaml:"encoding"`
	Level             string `yaml:"level"`
}

func NewConfig(cfgFile string) (*Config, error) {

	cfg := &Config{}

	if err := cleanenv.ReadConfig(cfgFile, cfg); err != nil {
		return nil, err
	}
	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
