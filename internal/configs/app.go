package configs

import (
	"fmt"
	"time"
	"strconv"

	"github.com/nturbo1/reverse-proxy/internal/log"
)

type AppConfigs struct {
	LogLevel string `json:"logLevel"`
	EnvFile string `json:"envFile"`
	RoutesMasterFile string `json:"routesMasterFile"`
	Server ServerConfigs `json:"server"`
}
func (ac AppConfigs) String() string {
	return fmt.Sprintf(
		"{logLevel: %s, envFile: %s, routesMasterFile: %s, server: %s}", 
		ac.LogLevel,
		ac.EnvFile,
		ac.RoutesMasterFile,
		ac.Server,
	)
}

type ServerConfigs struct {
	Port int64 `json:"port"`
	Timeout time.Duration `json:"timeout"` // Expected to be in milliseconds
}
func (sc ServerConfigs) String() string {
	return fmt.Sprintf("{port: %v, timeout: %s}", sc.Port, sc.Timeout)
}

func InitAppConfigs() (*AppConfigs, error) {
	initEnvVars()

	logLevel := log.INFO
	if level, ok := envVars[EnvVarReverseProxyLogLevel.String()]; ok {
		logLevel = level
	}
	routesMasterFile := envVars[EnvVarRoutesConfigsMasterFile.String()]
	var serverPort int64 = 8080
	if portStr, ok := envVars[EnvVarPort.String()]; ok {
		port, err := strconv.ParseInt(portStr, 10, 64)
		if err != nil {
			fmt.Println("Invalid port:", portStr)
			return nil, err
		}
		serverPort = port
	}


	appConfigs := &AppConfigs{
		LogLevel: logLevel,
		RoutesMasterFile: routesMasterFile,
		Server: ServerConfigs{
			Port: serverPort,
			Timeout: 10 * time.Millisecond,
		},
	}

	err := parseAppConfigs(envVars[EnvVarReverseProxyAppConfigFile.String()], appConfigs)
	if err != nil {
		fmt.Println("Failed to parse the app configs")
		return nil, err
	}

	return appConfigs, nil
}
