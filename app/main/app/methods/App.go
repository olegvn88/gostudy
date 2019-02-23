package main

import "fmt"

func main() {
	fmt.Println(WithFiles)
	fmt.Println(WithoutFiles)

	fmt.Println(Mytype{})
}

type Mytype struct {
	Num  int
	Name string
}

type MyInt int

type withFiles bool

var WithFiles = withFiles(true)

var WithoutFiles = withFiles(false)

func (m MyInt) showYourSelf() {
	fmt.Printf("%T %v\n", m, m)
}

func (m *MyInt) add(i MyInt) {
	*m = *m + MyInt(i)
}

//type mySliceStruct []myStruct
