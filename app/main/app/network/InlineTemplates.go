package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	tmpl := template.New("main")
	tmpl, _ = tmpl.Parse(
		`<div style="display: inline-block; border: 1px solid #aaa;
border-radius: 3px; padding: 30 px; margin: 20 px;">
{{if ne . "str"}}
no str
{{end}}
<pre>{{.}}</pre>
</div>`)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path

		c := http.Client{}
		resp, err := c.Get("http://artii.herokuapp.com/make?text=" + path)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("error"))
			return
		}

		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		//отдаем шаблон с данными
		tmpl.Execute(writer, string(body))
	})

	http.ListenAndServe(":8081", nil)
}
