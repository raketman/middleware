package middleware

import (
	"encoding/json"
	"os"
)

type DefaultAvailableClientResolver struct {
	FilePath string
}

// Конфигурация
type configuration struct {
	Clients []Client
}



func (d DefaultAvailableClientResolver) GetClients()  []Client {
	file, _ := os.Open(d.FilePath)
	decoder := json.NewDecoder(file)
	configuration := configuration{}
	err := decoder.Decode(&configuration)

	if err != nil {
		panic(err) // генерим панику, иначе скрипт будет работать некорректно
	}

	file.Close()

	return configuration.Clients
}