package hades

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

// exists returns whether the given file or directory exists
func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

// getPath returns whichever path exists
func getPath(file string, paths []string) string {
	for _, path := range paths {
		filePath := path + "/" + file
		if exists(filePath) {
			return filePath
		}
	}
	panic("Invalid Paths")
}

// GetConfig returns a map of all config in a toml file
func GetConfig(file string, paths []string) Config {

	filePath := getPath(file, paths)

	content, _ := ioutil.ReadFile(filePath)

	config := map[string]interface{}{}
	yaml.Unmarshal(content, &config)

	return Config{Data: config}
}
