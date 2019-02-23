package main

import (
	"errors"
	"fmt"
)

var ErrNamedType = New("EOF")

type errorString string

func (e errorString) Error() string {
	return string(e)
}

//New creates interface values of type error
func New(text string) error {
	return errorString(text)
}

var ErrStructType = errors.New("EOF")

func main() {

	if ErrNamedType == New("EOF") {
		fmt.Println("Named Type Error")
	}

	if ErrStructType == errors.New("EOF") {

	}
}
