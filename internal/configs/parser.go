package configs

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func parseAppConfigs(appConfigFile string, appConfigs *AppConfigs) error {
	appConfigFileBytes, err := os.ReadFile(appConfigFile)
	if err != nil {
		fmt.Printf("Failed to read from file '%s'\n", appConfigFile)
		return err
	}

	err = json.Unmarshal(appConfigFileBytes, appConfigs)
	if err != nil {
		fmt.Printf("Failed to json unmarshal the bytes of file '%s'\n", appConfigFile)
		return err
	}

	appConfigs.Server.Timeout *= time.Millisecond

	return nil
}
