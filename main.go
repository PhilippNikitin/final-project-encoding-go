package main

import (
	"fmt"

	"github.com/Yandex-Practicum/final-project-encoding-go/encoding" // импортируем внутренний пакет encoding
	"github.com/Yandex-Practicum/final-project-encoding-go/utils"    // импортируем внутренний пакет utils
)

// Функция Encode
// принимает один параметр (data) интерфейсного типа MyEncoder (импортирован из пакета encoding)
// применяет метод Encoding к полученному аргументу
// возвращает ошибку
func Encode(data encoding.MyEncoder) error {
	return data.Encoding()
}

func main() {
	// создаем JSON файл
	utils.CreateJSONFile()
	// создаем YAML файл
	utils.CreateYAMLFile()

	// объявляем переменную jsonData структурного типа JSONData
	jsonData := encoding.JSONData{FileInput: "jsonInput.json", FileOutput: "yamlOutput.yml"}
	// применяем функцию Encode к инициализированной переменной - проводим конвертацию из JSON в YAML
	err := Encode(&jsonData)
	// проверяем, если была сгенерирована ошибка, если да, то выводим сообщение об ошибке
	if err != nil {
		fmt.Printf("ошибка при перекодировании данных из JSON в YAML: %s", err.Error())
	}

	// объявляем переменную yamlData структурного типа YAMLData
	yamlData := encoding.YAMLData{FileInput: "yamlInput.yml", FileOutput: "jsonOutput.json"}
	// применяем функцию Encode к инициализированной переменной - проводим конвертацию из YAML в JSON
	err = Encode(&yamlData)
	// проверяем, если была сгенерирована ошибка, если да, то выводим сообщение об ошибке
	if err != nil {
		fmt.Printf("ошибка при перекодировании данных из YAML в JSON: %s", err.Error())
	}
}
