package http

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path

		c := http.Client{}

		resp, err := c.Get("http://artii.herokuapp.com/make?text=" + path)

		if err != nil {
			log.Println(err)
		}

		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		writer.WriteHeader(http.StatusOK)
		writer.Write(body)
	})

	http.ListenAndServe(":8081", nil)
}
