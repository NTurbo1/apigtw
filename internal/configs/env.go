package configs

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type EnvVar int

const (
	EnvVarPort EnvVar = iota
	EnvVarReverseProxyAppConfigFile
	EnvVarRoutesConfigsMasterFile
	EnvVarReverseProxyLogLevel
)

var envVarName = map[EnvVar]string{
	EnvVarPort:                      "PORT",
	EnvVarReverseProxyAppConfigFile: "REVERSE_PROXY_APP_CONFIG_FILE",
	EnvVarRoutesConfigsMasterFile:   "ROUTES_CONFIGS_MASTER_FILE",
	EnvVarReverseProxyLogLevel:      "REVERSE_PROXY_LOG_LEVEL",
}

func (ev EnvVar) String() string {
	return envVarName[ev]
}

var envVars = map[string]string{
	EnvVarPort.String():                      "8080",
	EnvVarReverseProxyAppConfigFile.String(): "app.configs.json",
	EnvVarRoutesConfigsMasterFile.String():   "routes.master.json",
	EnvVarReverseProxyLogLevel.String():      "REVERSE_PROXY_LOG_LEVEL",
}

func initEnvVars() {
	envKVs := os.Environ()
	for _, envKV := range envKVs {
		kv := strings.Split(envKV, "=")
		envVars[kv[0]] = kv[1]
	}
}

type Environment struct {
	Variables map[string]string `json:"variables"`
}

func (e Environment) String() string {
	return fmt.Sprintf("{variables: %s}", e.Variables)
}

func GetEnv(appConfigs *AppConfigs) (*Environment, error) {
	env, err := parseEnv(appConfigs.EnvFile)
	if err != nil {
		fmt.Printf("Failed to get the environment data due to: %s\n", err)
		return nil, err
	}

	return env, nil
}

func parseEnv(envFilename string) (*Environment, error) {
	envContent, err := os.ReadFile(envFilename)
	if err != nil {
		fmt.Printf("Failed to read from the env file '%s' due to: %s\n", envFilename, err)
		return nil, err
	}

	var env Environment
	err = json.Unmarshal(envContent, &env)
	if err != nil {
		fmt.Printf("Failed to json unmarshal the env file '%s' content\n", envFilename)
		return nil, err
	}

	err = ReplaceEnvVarsInConfigs(&env, envVars)
	if err != nil {
		fmt.Println(
			"Failed to replace system environment variables in the application environment configs",
		)
		return nil, err
	}

	return &env, nil
}
