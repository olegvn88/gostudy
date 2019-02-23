package main

import "fmt"

func main() {
	sa := SecretAgent{Person: Person{"James", "1223432345"}}

	fmt.Println("secret inn", sa.GetName())
}

type Person struct {
	Name string
	inn  string
}

type Stuff struct {
	inn int
}

type SecretAgent struct {
	Person //embedded structure
	Stuff
	//Name int
	LicenseToKill bool
}

func (p Person) GetName() string {
	return p.Name
}

func (s SecretAgent) GetName() string {
	return "CLASSIFIED"
}
