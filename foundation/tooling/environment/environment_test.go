package environment

import (
	"os"
	"testing"
)

func Test_GetStrEnv(t *testing.T) {
	err := os.Setenv("TEST_ENV", "test")
	if err != nil {
		t.Error("Error setting environment variable")
	}

	got := GetStrEnv("TEST_ENV", "default")

	if got != "test" {
		t.Errorf("GetStrEnv() = %s; want test", got)
	}
}

func Test_GetBoolEnv(t *testing.T) {
	err := os.Setenv("TEST_ENV", "true")
	if err != nil {
		t.Error("Error setting environment variable")
	}

	got := GetBoolEnv("TEST_ENV", false)

	if got != true {
		t.Errorf("GetBoolEnv() = %t; want true", got)
	}
}
