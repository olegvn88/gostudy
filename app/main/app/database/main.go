package main

import (
	"fmt"
	"net/http"
	"time"
)

var loginFormTmpl = `
	<html>
		<body>
		<form action="/get_cookie" method="post">
			Login: <input type="text" name="login">
			Password: <input type="password" name="password">
			Submit: <input type="submit" value="Login">
		</form>
		</body>
	</html>
`

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		sessionID, err := request.Cookie("session _id")
		if err == http.ErrNoCookie {
			writer.Write([]byte(loginFormTmpl))
			return
		} else if err != nil {
			panic(err)
		}
		fmt.Fprint(writer, "Welcome, "+sessionID.Value)
	})

	http.HandleFunc("/get_cookie", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		inputLogin := request.Form["login"][0]
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "session_id", Value: inputLogin, Expires: expiration}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", http.StatusFound)
	})

	http.ListenAndServe("8081", nil)
}