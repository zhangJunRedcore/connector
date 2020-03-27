package utils

import (
	"io/ioutil"
	"strings"
	"sync"
)

type Yaml struct {
	YamlPath string
}

var mutex sync.Mutex

func (yaml *Yaml) Read() ([]string, error) {
	mutex.Lock()
	defer mutex.Unlock()

	input, err := ioutil.ReadFile(yaml.YamlPath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(input), "\n")
	return lines, nil
}

func (yaml *Yaml) Write(lines []string) error {
	mutex.Lock()
	defer mutex.Unlock()

	output := strings.Join(lines, "\n")
	err := ioutil.WriteFile(yaml.YamlPath, []byte(output), 0644)
	if err != nil {
		return err
	}
	return nil
}

func (yaml *Yaml) modifyAll(lines []string, key string, value string) []string {
	for i, line := range lines {
		if strings.Contains(line, key) {
			values := strings.Split(line, ":")
			lines[i] = values[0] + ": " + value
		}
	}

	return lines
}

func (yaml *Yaml) Modify(key string, value string) error {
	lines, err := yaml.Read()
	if err != nil {
		return err
	}
	lines = yaml.modifyAll(lines, key, value)
	return yaml.Write(lines)
}
