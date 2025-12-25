package configs

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

const (
	RegexEnvVarNameInsideString = "\\$[a-zA-Z0-9_]+"
)

// Function ReplaceEnvVarsIn identifies env variables in a given text string value and replaces each
// of those env variable occurrences with their value found from a given map of env variables.
//
// An env variable occurrence inside a string value is considered to be any substring value that is
// preceded by '$' character and contains one or more alphanumeric and '_' (underscore) characters.
func ReplaceEnvVarsIn(text string, envVars map[string]string) (string, error) {
	re := regexp.MustCompile(RegexEnvVarNameInsideString)
	envVarNames := re.FindAll([]byte(text), -1)

	for _, envVarName := range envVarNames {
		envVarNameStr := string(envVarName)
		if envVarVal, ok := envVars[envVarNameStr[1:]]; ok {
			text = strings.ReplaceAll(text, envVarNameStr, envVarVal)
		} else {
			return "", fmt.Errorf("Environment variable '%s' is not defined!", envVarNameStr)
		}
	}

	return text, nil
}

func ReplaceEnvVarsInConfigs(configs any, envVars map[string]string) error {
	return replaceEnvVarsInValue(reflect.ValueOf(configs), envVars)
}

func replaceEnvVarsInValue(v reflect.Value, envVars map[string]string) error {
    switch v.Kind() {

	case reflect.Ptr:
        if v.IsNil() {
            return nil
        }
        return replaceEnvVarsInValue(v.Elem(), envVars)

    case reflect.String:
        replaced, err := ReplaceEnvVarsIn(v.String(), envVars)
        if err != nil {
            return err
        }
        if v.CanSet() {
            v.SetString(replaced)
        }

    case reflect.Struct:
        for i := 0; i < v.NumField(); i++ {
            if err := replaceEnvVarsInValue(v.Field(i), envVars); err != nil {
                return err
            }
        }

    case reflect.Map:
        for _, key := range v.MapKeys() {
            val := v.MapIndex(key)

            // Create a mutable copy
            newVal := reflect.New(val.Type()).Elem()
            newVal.Set(val)

            if err := replaceEnvVarsInValue(newVal, envVars); err != nil {
                return err
            }

            v.SetMapIndex(key, newVal)
        }
    }

    return nil
}
