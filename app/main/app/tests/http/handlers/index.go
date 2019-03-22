package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
)

func (h Handler) HandleIndex(w http.ResponseWriter, r *http.Request) {
	fileContents, err := ioutil.ReadFile("index.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Write(fileContents)
}
