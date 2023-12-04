package encoding

import (
	"encoding/json"

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
	bytes, err := json.Marshal(j)
	if err != nil {
		return err
	}
	str := string(bytes)
	result := struct {
		Data string `yaml:"data"`
	}{}
	err = yaml.Unmarshal([]byte(str), &result)
	if err != nil {
		panic(err)
	}
	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	// Ниже реализуйте метод
	bytes, err := yaml.Marshal(y)
	if err != nil {
		return err
	}
	str := string(bytes)
	result := struct {
		Data string `json:"data"`
	}{}
	err = json.Unmarshal([]byte(str), &result)
	if err != nil {
		panic(err)
	}
	return nil
}
