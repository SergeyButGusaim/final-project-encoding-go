package encoding

import (
	"encoding/json"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	// ниже реализуйте метод
	jsonFile, err := os.ReadFile(j.FileInput)
	if err != nil {
		return err
	}
	var dockerCompose models.DockerCompose
	err = json.Unmarshal(jsonFile, &dockerCompose)
	if err != nil {
		return err
	}
	yamlData, err := yaml.Marshal(dockerCompose)
	if err != nil {
		return err
	}
	err = os.WriteFile(j.FileOutput, yamlData, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	// Ниже реализуйте метод
	yamlFile, err := os.ReadFile(y.FileInput)
	if err != nil {
		return err
	}
	var dockerCompose models.DockerCompose
	err = yaml.Unmarshal(yamlFile, &dockerCompose)
	if err != nil {
		return err
	}
	jsonData, err := json.Marshal(dockerCompose)
	if err != nil {
		return err
	}
	err = os.WriteFile(y.FileOutput, jsonData, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
