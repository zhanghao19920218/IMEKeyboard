package IMEKeyboard

import (
	"gopkg.in/yaml.v3"
	"os"
)

// FileToMobileYaml Reading the yaml file, to make it to yaml model
func FileToMobileYaml(filePath string) (model MobileYamlModel, errorType *KeyboardError) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		errorType = &KeyboardError{
			Message:   "Read File Error",
			ErrorType: ReadYamlError,
		}
		return
	}
	err2 := yaml.Unmarshal(file, &model)

	if err2 != nil {
		errorType = &KeyboardError{
			Message:   "Parse Yaml File Error",
			ErrorType: YamlReaderError,
		}
		return
	}
	return
}
