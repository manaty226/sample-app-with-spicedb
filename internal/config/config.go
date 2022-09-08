package config

import (
	"encoding/csv"
	"io/ioutil"
	"os"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	SpiceHost     string `env:"SPICE_DB_HOST" envDefault:"127.0.0.1"`
	SpicePort     int    `env:"SPICE_DB_PORT" envDefault:"3001"`
	NameSpaceFile string `env:"NAME_SPACE_FILE" envDefault:".config/namespace.txt"`
	NameSpace     string
	AuthnFile     string `env:"AUTHN_FILE" envDefault:".config/user.csv"`
	AuthnUsers    map[string]string
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	if err := cfg.setAuthnUsers(); err != nil {
		return nil, err
	}

	if err := cfg.setNameSpace(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) setAuthnUsers() error {
	// Set User list who can use APIs
	users, err := readCsv(c.AuthnFile)
	if err != nil {
		return err
	}
	authnUsers := map[string]string{}
	for _, user := range users {
		authnUsers[user[0]] = user[1]
	}
	c.AuthnUsers = authnUsers
	return nil
}

func (c *Config) setNameSpace() error {
	nameSpace, err := readText(c.NameSpaceFile)
	if err != nil {
		return err
	}
	c.NameSpace = nameSpace
	return nil
}

func readText(path string) (string, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

func readCsv(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := csv.NewReader(file)
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return rows, nil
}
