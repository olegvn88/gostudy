package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 3.4
	fmt.Println("reflect.Type:", reflect.TypeOf(x))

	// reflect.Value != значению, переданному на вход
	fmt.Println("reflect.Type:", reflect.ValueOf(x).String())

	v := reflect.ValueOf(x)

	fmt.Println("Тип Type:", v.Type())
	fmt.Println("Тип float64:", v.Kind() == reflect.Float64)
	fmt.Println("Значение", v.Float())

	type MyInt int
	var c MyInt = 7
	v = reflect.ValueOf(c)
	fmt.Println("kind is int: ", v.Kind() == reflect.Int)
	access()
}

func access() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println(v)

	//v.SetFloat(7.1)

	p := reflect.ValueOf(&x)
	fmt.Println("Type of p:", p.Type())
	fmt.Println("Settability of p: ", p.CanSet())

	//Теперь, испоьзуя  Elem, получим Value, лежащее по ссылке
	v = p.Elem()
	fmt.Println("settability of v:", v.CanSet())

	v.SetFloat(7.1)

	fmt.Println(v.Interface())

}
