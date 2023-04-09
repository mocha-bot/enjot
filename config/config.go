package config

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Logger struct {
	Level string `envconfig:"level" default:"info"`
}

type Server struct {
	Host                string        `envconfig:"host" default:"0.0.0.0"`
	Port                int           `envconfig:"port" default:"8080"`
	WaitGracefulTimeout time.Duration `envconfig:"wait_graceful_timeout" default:"30s"`
	WriteTimeout        time.Duration `envconfig:"write_timeout" default:"15s"`
	ReadTimeout         time.Duration `envconfig:"read_timeout" default:"15s"`
	IdleTimeout         time.Duration `envconfig:"idle_timeout" default:"60s"`
}

func (s Server) Address() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

type Mysql struct {
	Host     string `envconfig:"host"`
	Port     int    `envconfig:"port"`
	Username string `envconfig:"username"`
	Password string `envconfig:"password"`
	DBName   string `envconfig:"db_name"`
}

func (m Mysql) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.DBName,
	)
}

type Database struct {
	Mysql Mysql `envconfig:"mysql"`
}

type Configuration struct {
	Logger   Logger   `envconfig:"logger"`
	Server   Server   `envconfig:"server"`
	Database Database `envconfig:"database"`

	TokenKeySecret string `envconfig:"token_key_secret"`
}

const NAMESPACE = "enjot"

func Get() Configuration {
	var configuration Configuration

	err := envconfig.Process(NAMESPACE, &configuration)

	if err != nil {
		return Configuration{}
	}

	return configuration
}
