package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

/// The config package is how Gosh loads, defines, and alters standard
/// behavior with the shell. By default, Gosh has a default config.yaml
/// in $HOME/.config/gosh/config.yaml, which can be modified by the user.
///
/// Main features:
///
/// GoshConfig struct => A struct that holds all of the default + user configs.

var goshConfigLocation string

func init() {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		panic("No user specified!")
	}

	goshConfigLocation = homeDir + "/.config/gosh/config.yaml"

}

type GoshConfig struct {
	Aliases map[string]string `yaml:"aliases"`
}

func NewConfig() *GoshConfig {
	aliases := make(map[string]string, 0)
	return &GoshConfig{aliases}
}

// FindAlias looks up a user's input in (*c).Aliases and attempts to replace it
// with what it finds (if any).
func (c *GoshConfig) FindAlias(text string) (string, bool) {
	al, ok := c.Aliases[text]
	return al, ok
}

func FromConfigFile() (*GoshConfig, error) {
	// fmt.Printf("Looking for history file in %s\n", goshConfigLocation)
	c := NewConfig()
	content, err := ioutil.ReadFile(goshConfigLocation)

	if err != nil {
		return nil, fmt.Errorf("No config file found!")
	}

	err = yaml.Unmarshal(content, &c)
	return c, err
}

// In the case that a user enters input starting with "alias",
// this func runs, parsing the input.
// func (c *GoshConfig) AddUserAlias(text string) error {
// }
