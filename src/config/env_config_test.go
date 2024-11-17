package config

import (
	"os"
	"testing"
)

func TestGetEnvAsString(t *testing.T) {
	os.Setenv("GRIMOIRE_TEST_STRING", "test_value")
	defer os.Unsetenv("GRIMOIRE_TEST_STRING")

	value := getEnvAsString("TEST_STRING", "default_value")
	if value != "test_value" {
		t.Errorf("Expected 'test_value', got '%s'", value)
	}

	value = getEnvAsString("NON_EXISTENT_KEY", "default_value")
	if value != "default_value" {
		t.Errorf("Expected 'default_value', got '%s'", value)
	}
}

func TestGetEnvAsApiKey(t *testing.T) {
	os.Setenv("GRIMOIRE_TEST_API_KEY", "api_key_value")
	defer os.Unsetenv("GRIMOIRE_TEST_API_KEY")

	value, err := getEnvAsApiKey("TEST_API_KEY")
	if err != nil || value != "api_key_value" {
		t.Errorf("Expected 'api_key_value', got '%s' with error '%v'", value, err)
	}

	value, err = getEnvAsApiKey("NON_EXISTENT_KEY")
	if err == nil || value != "" {
		t.Errorf("Expected error and empty value, got '%s' with error '%v'", value, err)
	}
}

func TestGetEnvAsBool(t *testing.T) {
	os.Setenv("GRIMOIRE_TEST_BOOL", "true")
	defer os.Unsetenv("GRIMOIRE_TEST_BOOL")

	value := getEnvAsBool("TEST_BOOL", false)
	if value != true {
		t.Errorf("Expected 'true', got '%v'", value)
	}

	value = getEnvAsBool("NON_EXISTENT_KEY", false)
	if value != false {
		t.Errorf("Expected 'false', got '%v'", value)
	}
}
