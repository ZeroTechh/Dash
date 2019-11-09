package dash

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
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

	content, err := ioutil.ReadFile(filePath)

	config := map[interface{}]interface{}{}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		panic(fmt.Sprintf("Cant Load Yaml Config Because Of Error %+v", err))
	}

	return Config{data: config}
}
