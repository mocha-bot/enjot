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

type Configuration struct {
	Logger Logger `envconfig:"logger"`
	Server Server `envconfig:"server"`
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
