package config

import (
	"os"
	"testing"
)

func init() {
	currPath, _ := os.Getwd()
	goshConfigLocation = currPath + "/testConfig.yaml"
}

func mockConfig() *GoshConfig {
	c := NewConfig()
	c.Aliases["a"] = "b"

	return c
}

func TestNewConfig(t *testing.T) {
	c := mockConfig()

	if al := c.Aliases["a"]; al != "b" {
		t.Fatalf("Have alias %s, want \"b\"", al)
	}
}

func TestFindAlias(t *testing.T) {
	c := mockConfig()
	_, ok := c.FindAlias("a")

	if !ok {
		t.Fatalf("Unable to find alias \"%s\" for config", "a")
	}
}

func TestFromConfigFile(t *testing.T) {
	config, err := FromConfigFile()

	if err != nil {
		t.Fatalf("%v", err)
	}

	if config == nil {
		t.Fatalf("Unable to load config for location %s", goshConfigLocation)
	}

	if l := len(config.Aliases); l != 1 {
		t.Fatalf("Have len %d, want len 1", l)
	}
}

func TestFromConfigNoLocation(t *testing.T) {
	goshConfigLocation = "/tmp/somerandomconfig.Yaml"

	config, err := FromConfigFile()

	if config != nil || err == nil {
		t.Fatalf("Unexpected config file found for location %s", goshConfigLocation)
	}
}
