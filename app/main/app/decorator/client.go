package main

import "net/http"

type Client interface {
	Do(r *http.Request) (*http.Response, error)
}

// A decorator wraps a client with extra behavior
type Decorator func(Client) Client

// ClientFunc is a function type that implements the Client interface
type ClientFunc func(*http.Request) (*http.Response, error)

// Do does the request
func (f ClientFunc) Do(r *http.Request) (*http.Response, error) {
	return f(r)
}
