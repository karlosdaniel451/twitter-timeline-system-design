package config

import (
	"fmt"
	"strings"
)

type AppConfig struct {
	ListenerPort   int
	AppEnvironment AppEnvironment
}

func NewAppConfig(listenerPort int, appEnvironment AppEnvironment) *AppConfig {
	return &AppConfig{
		ListenerPort:   listenerPort,
		AppEnvironment: appEnvironment,
	}
}

type AppEnvironment int

const (
	Development AppEnvironment = iota
	Testing
	Production
)

// ParseAppEnvironment parses an AppEnvironment enum value from its string
// representation. If an empty string is passed, then it returns Development.
func ParseAppEnvironment(environmentUnparsed string) (AppEnvironment, error) {
	switch strings.ToLower(environmentUnparsed) {
	case "":
		return Development, nil
	case "development":
		return Development, nil
	case "testing":
		return Testing, nil
	case "production":
		return Production, nil
	default:
		return 0, fmt.Errorf("invalid app environment")
	}
}

func (appEnv AppEnvironment) String() string {
	switch appEnv {
	case Development:
		return "Development"
	case Testing:
		return "Testing"
	case Production:
		return "Production"
	default:
		return ""
	}
}
