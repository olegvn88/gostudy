package main

import (
	"fmt"
)

type Flyer interface {
	fly()
}

type Voice interface {
	makeVoice()
}

type Bird struct {
	Name string
}

type Animal struct {
	Name string
}

func (a Animal) makeVoice() {
	fmt.Println(a.Name + " makes noise")
}

func (b Bird) fly() {
	fmt.Println(b.Name + " is flying now")
}

func doFly(f Flyer) {
	f.fly()
}

func doVoice(voice Voice) {
	voice.makeVoice()
}

type typeOfVoice struct {
	cow, dog, cat string
}

func main() {
	huwk := Bird{"Hawk"}
	doFly(huwk)

	//voice := typeOfVoice{"mooooo", "bark", "myau"}

	dogsVoice := Animal{"Bob"}
	doVoice(dogsVoice)
	GoFly(huwk)
}

//Type Assertion

func GoFly(f Flyer) {
	f.fly()
	if b, ok := f.(Bird); ok {
		fmt.Println(b.Name)
	}
}
