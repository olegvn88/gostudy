package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//Мапинг наших ресурсов.

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//w http.ResponseWriter - структура, которая вернется пользователю в виде ответа(response)
		//r *http.Request - тот запрос, коотрый к нам пришел.
		path := r.URL.Path

		//  Создаем http клиент. В структуру можно передать таймаут, куки и прочую инфу о запросе

		c := http.Client{}
		resp, err := c.Get("http://artii.herokuapp.com/make?text=" + path)
		if err != nil {
			log.Println(err)
		}
		//Нужно закрыть тело, когда прочитаем, что нужно
		defer resp.Body.Close() // если не закрыть body, го не позволит keep alive
		body, _ := ioutil.ReadAll(resp.Body)

		// status - OK
		w.WriteHeader(200)
		w.Write(body)
	})

	http.ListenAndServe(":8081", nil)
}
