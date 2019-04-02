package main

import (
	"fmt"
)

type Animal struct {
	Name string
	Type string
	Age  int
}

type AnimalSkills interface {
	makeVoice()
	DoSleep()
	ToEat()
}

func (a Animal) SetName(newName string) {
	a.Name = newName
}

func (a Animal) SetType(newType string) {
	a.Type = newType
}

func (a Animal) SetAge(newAge int) {
	a.Age = newAge
}

func (a Animal) makeVoice() {
	fmt.Println(a.Name, "making voice now")
	fmt.Println(a.Name, "has type", a.Type)
	fmt.Println(a.Name, "has age", a.Age)
}

func (a Animal) DoSleep() {
	fmt.Println(a.Name, "is sleeping now")
}

func (a Animal) ToEat() {
	fmt.Println(a.Name, "is eating now")
}

func describeAnimal(skills AnimalSkills) {
	skills.DoSleep()
	skills.makeVoice()
	skills.ToEat()
}

func main() {
	r := Animal{"Markiza", "cat", 4}

	describeAnimal(r)
}
