package encoding

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3" // импортируем пакет для работы с YAML

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

// Метод Encoding структуры JSONData перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	// ниже реализуйте метод

	// шаг 1: читаем файл и записываем данные в виде слайса байт в переменную jsonFile
	jsonFile, err := os.ReadFile(j.FileInput) // в переменную jsonFile будет записан слайс байт
	if err != nil {
		fmt.Printf("ошибка при чтении файла: %s", err.Error())
		return err
	}

	// шаг 2: десериализуем данные из jsonFile обратно в структуру DockerCompose
	// инициализируем новую структуру типа DockerCompose
	var dockercompose models.DockerCompose

	err = json.Unmarshal(jsonFile, &dockercompose)
	if err != nil {
		fmt.Printf("ошибка при десериализации из jsonFile: %s", err.Error())
		return err
	}

	// теперь мы сохранили данные из JSONData в структуру dockercompose

	// шаг 3: сериализуем данные в формат YAML
	yamlData, err := yaml.Marshal(&dockercompose)
	if err != nil {
		fmt.Printf("ошибка при сериализации в yaml: %s", err.Error())
		return err
	}

	//  шаг 4: создаем файл с названием, как в FileOutput
	f, err := os.Create(j.FileOutput)
	if err != nil {
		fmt.Printf("ошибка при создании файла: %s", err.Error())
		return err
	}
	// когда программа завершается, закрываем дескриптор файла
	defer f.Close()

	// записываем слайс байт в файл
	_, err = f.Write(yamlData)
	if err != nil {
		fmt.Printf("ошибка при записи данных в файл: %s", err.Error())
		return err
	}

	return nil
}

// Метод Encoding структуры YAMLData перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	// Ниже реализуйте метод
	// шаг 1: читаем файл и записываем данные в виде слайса байт в переменную yamlFile
	yamlFile, err := os.ReadFile(y.FileInput) // в переменную jsonFile будет записан слайс байт
	if err != nil {
		fmt.Printf("ошибка при чтении файла: %s", err.Error())
		return err
	}

	// шаг 2: десериализуем данные из yamlFile в структуру DockerCompose
	// инициализируем новую структуру типа DockerCompose
	var dockercompose models.DockerCompose

	err = yaml.Unmarshal(yamlFile, &dockercompose)
	if err != nil {
		fmt.Printf("ошибка при десериализации из yamlFile: %s", err.Error())
		return err
	}

	// теперь мы сохранили данные из YAMLData в структуру dockercompose

	// шаг 3: сериализуем данные в формат JSON
	jsonData, err := json.Marshal(&dockercompose)
	if err != nil {
		fmt.Printf("ошибка при сериализации в json: %s", err.Error())
		return err
	}

	//  шаг 4: создаем файл с названием, как в FileOutput
	f, err := os.Create(y.FileOutput)
	if err != nil {
		fmt.Printf("ошибка при создании файла: %s", err.Error())
		return err
	}
	// когда программа завершается, закрываем дескриптор файла
	defer f.Close()

	// записываем слайс байт в файл
	_, err = f.Write(jsonData)
	if err != nil {
		fmt.Printf("ошибка при записи данных в файл: %s", err.Error())
		return err
	}

	return nil
}
