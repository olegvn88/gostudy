package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Todo struct {
	Name string `json:"name"`
	Done bool   `json:"done"`
}

func main() {
	todos := []Todo{
		{"Выучить Go", false},
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//здесь нужно отдать статический файл, который будет общаться с API браузера
		//открываем файл
		fileContents, err := ioutil.ReadFile("/home/onest/go/src/github.com/olegvn88/gostudy/app/main/app/network/json_api/index.html")

		if err != nil {
			fmt.Println(err)
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		//выводим содержание файла
		writer.Write(fileContents)
	})

	http.HandleFunc("/todos", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("request ", request.URL.Path)
		defer request.Body.Close()

		//разные методы обрабатываются по-разному
		switch request.Method {
		// GET для получения данных
		case http.MethodGet:
			productsJson, _ := json.Marshal(todos)
			writer.Header().Set("Content-type", "application/json")
			writer.WriteHeader(http.StatusOK)
			writer.Write(productsJson)
		//POST для добавления чего-то нового
		case http.MethodPost:
			decoder := json.NewDecoder(request.Body)
			todo := Todo{}
			//преобразуем  json запрос в структуру
			err := decoder.Decode(&todo)
			if err != nil {
				log.Println(err)
				writer.WriteHeader(http.StatusBadRequest)
				return
			}
			todos = append(todos, todo)
		//PUT для обновления существующей информации
		case http.MethodPut:
			id := request.URL.Path[len("/todos/"):]
			index, _ := strconv.ParseInt(id, 10, 0)
			todos[index].Done = true
		default:
			writer.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8081", nil)

}
