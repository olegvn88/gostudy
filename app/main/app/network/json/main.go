package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Joke struct {
	ID   uint32 `json:"id"`
	Joke string `json:"joke"`
}

type JokeResponse struct {
	Type  string `json:"type"`
	Value Joke   `json:"value"`
}

func main() {
	GetJoke()
}

func GetJoke() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		defer request.Body.Close()

		c := http.Client{}
		resp, err := c.Get("http://api.icndb.com/jokes/random?limitTo=[nerdy]")
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		joke := JokeResponse{}

		err = json.Unmarshal(body, &joke)
		if err != nil {
			log.Fatal(err)
		}

		//writer.Write([]byte(joke.Value.Joke))

		writer.Write([]byte(joke.Value.Joke))
	})

	http.ListenAndServe(":8081", nil)
}
