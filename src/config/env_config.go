package config

import "os"

const EnvPrefix string = "GRIMOIRE_"

type EnvError struct {
	Key string
}

func (e *EnvError) Error() string {
	return "environment variable not found: " + e.Key
}

func getEnvAsString(key, defaultValue string) string {
	if value, exists := os.LookupEnv(EnvPrefix + key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsApiKey(key string) (string, error) {
	if value, exists := os.LookupEnv(EnvPrefix + key); exists {
		return value, nil
	}
	return "", &EnvError{key}
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if value, exists := os.LookupEnv(EnvPrefix + key); exists {
		return value == "true"
	}
	return defaultValue
}
