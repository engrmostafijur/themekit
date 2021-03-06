package themekit

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v1"
)

// DefaultEnvironment ... TODO
const DefaultEnvironment string = "development"

// Environments ... TODO
type Environments map[string]Configuration

// LoadEnvironments ... TODO
func LoadEnvironments(contents []byte) (envs Environments, err error) {
	envs = make(Environments)
	err = yaml.Unmarshal(contents, &envs)
	if err == nil {
		for key, conf := range envs {
			environmentConfig, err := conf.Initialize()
			if err != nil {
				return nil, fmt.Errorf("could not load environment \"%s\": %s", key, err)
			}
			envs[key] = environmentConfig
		}
	}
	return
}

// SetConfiguration ... TODO
func (e Environments) SetConfiguration(environmentName string, conf Configuration) {
	e[environmentName] = conf
}

// GetConfiguration ... TODO
func (e Environments) GetConfiguration(environmentName string) (conf Configuration, err error) {
	conf, exists := e[environmentName]
	if !exists {
		err = fmt.Errorf("%s does not exist in this environments list", environmentName)
	}
	return
}

func (e Environments) Write(w io.Writer) error {
	bytes, err := yaml.Marshal(e)
	if err == nil {
		_, err = w.Write(bytes)
	}
	return err
}

func (e Environments) String() string {
	bytes, err := yaml.Marshal(e)
	if err != nil {
		return "environments: cannot serialize"
	}
	return string(bytes)
}

// Save ... TODO
func (e Environments) Save(location string) error {
	file, err := os.OpenFile(location, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer file.Close()
	if err == nil {
		err = e.Write(file)
	}
	return err
}

// LoadEnvironmentsFromFile ... TODO
func LoadEnvironmentsFromFile(location string) (env Environments, err error) {
	contents, err := ioutil.ReadFile(location)
	if err == nil {
		return LoadEnvironments(contents)
	}
	return
}
