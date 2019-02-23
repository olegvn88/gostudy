package main

import "fmt"

type iStuff interface {
	DoStuff()
}

type stuff struct {
	iStuff
	Name string
}

type realStuff string

func (r realStuff) DoStuff() {
	fmt.Println(r)
}

type fakeStuff int

func (f fakeStuff) DoStuff() {
	fmt.Println("It's a trap")
}

func (s stuff) SomeComplex() {
	s.DoStuff()
}
func main() {
	r := realStuff("Hey")
	f := fakeStuff(0)

	rS := stuff{r, "stuff"}
	rS.SomeComplex()
	fS := stuff{f, "fake"}
	fS.SomeComplex()
}
