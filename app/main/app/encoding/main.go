package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := []byte(`
{
"id": "121212",
"title": "My awesome app",
"type": "baloon",
"person": {
		"name": "Oleg",
		"job": "go_developer"
         }
}
`)

	var app App
	err := json.Unmarshal(data, &app)
	if err != nil {
		panic(err)
	}

	fmt.Printf("App: %+v\n", app)

	fmt.Println(app.Person.Name)
	fmt.Println(app.Person.job)
	//================================
	var app2 Baloon

	err = json.Unmarshal(data, &app2)
	if err != nil {
		panic(err)
	}
	//var per Person
	fmt.Printf("App: %+v\n", app2)

	ea := app2.Person2.UnmarshalJSON(data)

	err = json.Unmarshal(data, ea)

	fmt.Println()

}

type App struct {
	Id     string `json:"id"`
	Title  string `json:"title,omitempty"`
	Data   int    `json:"-"`
	Person Person `json:"person"`
	//Все неэкспортируемы поля не учавствуют в маршалинге
	hidden string
}

type Person struct {
	Name string `json:"name"`
	job  string `json:"job"`
}

type Baloon struct {
	Id      string `json:"id"`
	Title   string `json:"title,omitempty"`
	Data    int    `json:"-"`
	hidden  string
	Type    string          `json:"string"`
	Person2 json.RawMessage `json:"person"`
}
