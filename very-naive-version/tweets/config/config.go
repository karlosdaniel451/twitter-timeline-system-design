package config

import (
	"fmt"
	"strings"
)

type AppConfig struct {
	ListenerPort       int
	AppEnvironmentType AppEnvironmentType
	DatabaseHost       string
	DatabaseUser       string
	DatabasePort       int
	DatabaseName       string
	DatabasePassword   string
}

func NewAppConfig(listenerPort int, appEnvironment AppEnvironmentType) *AppConfig {
	return &AppConfig{
		ListenerPort:       listenerPort,
		AppEnvironmentType: appEnvironment,
	}
}

func NewEmptyAppConfig() *AppConfig {
	return &AppConfig{}
}

func (a AppConfig) String() string {
	return fmt.Sprintf(
		"ListenerPort: %d\nAppEnvironment: %s",
		a.ListenerPort, a.AppEnvironmentType.String(),
	)
}

type AppEnvironmentType int

const (
	Development AppEnvironmentType = iota
	Debugging
	Testing
	Production
)

// ParseAppEnvironmentType parses an AppEnvironmentType enum value from its string
// representation. If an empty string is passed, then it returns Development.
func ParseAppEnvironmentType(envTypeUnparsed string) (AppEnvironmentType, error) {
	switch strings.ToLower(envTypeUnparsed) {
	case "":
		return Development, nil
	case "development":
		return Development, nil
	case "debugging":
		return Debugging, nil
	case "testing":
		return Testing, nil
	case "production":
		return Production, nil
	default:
		return 0, fmt.Errorf("invalid app environment")
	}
}

func (appEnv AppEnvironmentType) String() string {
	switch appEnv {
	case Development:
		return "Development"
	case Debugging:
		return "Debugging"
	case Testing:
		return "Testing"
	case Production:
		return "Production"
	default:
		return ""
	}
}
