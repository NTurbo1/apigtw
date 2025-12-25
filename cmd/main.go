package main

import (
	"fmt"

	"github.com/nturbo1/reverse-proxy/internal/configs"
	"github.com/nturbo1/reverse-proxy/internal/log"
	"github.com/nturbo1/reverse-proxy/internal/server"
)

func main() {
	appConfigs, env, err := setUpConfigsAndEnv()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Initializing a logger...")
	log.InitLogger(appConfigs.LogLevel)
	log.Debug("Configurations: %s", appConfigs)
	log.Debug("Environment: %s", env)
	
	server := server.NewServer(appConfigs, env)
	log.Info("Listening on port %d...", appConfigs.Server.Port)
	err = server.ListenAndServe()
	log.Fatal("%s", err)
}

func setUpConfigsAndEnv() (*configs.AppConfigs, *configs.Environment, error) {
	appConfigs, err := configs.InitAppConfigs()
	if err != nil {
		fmt.Printf("Failed to initialize app configs due to: %s\n", err)
		return nil, nil, err
	}

	env, err := configs.GetEnv(appConfigs)
	if err != nil {
		fmt.Printf("Failed to set up the Environment due to: %s\n", err)
		return nil, nil, err
	}
	fmt.Printf("Environment: %v\n", env)

	fmt.Printf("Raw app configs: %s\n", appConfigs)
	fmt.Println("Replacing variable occurrences with their values in the app configs...")
	err = configs.ReplaceEnvVarsInConfigs(appConfigs, env.Variables)
	if err != nil {
		fmt.Println("Failed to replace env variables occurrences with their values in the app configs")
		return nil, nil, err
	}
	fmt.Printf("Processed app configs: %s\n", appConfigs)

	return appConfigs, env, nil
}
