package main

import (
	"io"
	"net/http"
)

func Logging(l *io.Writer) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			l.PrintF("%s: %s %s", r.UserAgent(), r.Method, r.URL)
			return c.Do(r)
		})

	}
}
