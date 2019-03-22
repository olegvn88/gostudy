package tests

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestSum(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
	//В Go нет встроенной функции assert, все проверки происходят так
	expected := 4
	actual := Sum(1, 3)
	if actual != expected {
		t.Error("Achtung!", "expected", expected, "got", actual)
	}
}

func Sum(a int, b int) int {
	return a + b
}

func TestDifference(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
	//В Go нет встроенной функции assert, все проверки происходят так
	expected := 2
	actual := Difference(5, 3)
	if actual != expected {
		t.Error("Achtung!", "expected", expected, "got", actual)
	}
	t.Log("I am log")
	t.Skip()
}

func Difference(a int, b int) int {
	return a - b

}

type User struct {
	Name string
	Age  int
}

func TestAddUser(t *testing.T) {
	t.Parallel()
	users := []User{}
	//Then pass the slice as a reference to some function
	AddUser(&users, "Vasya", 22)

	if len(users) == 0 {
		//Fatal не продолжит выполнение кейса
		t.Fatal("Empty slice")
	}

	expected := []User{
		{
			Name: "Vasya",
			Age:  22,
		},
	}
	// Для сравнения слайсов и структур нужно использовать DeepEqual
	if !reflect.DeepEqual(users, expected) {
		t.Errorf("expected %+v, got %+v", expected, users)
	}
}

func AddUser(user *[]User, name string, age int) {
	//*[]User would be a pointer to a slice of User
	*user = []User{{name, age}}

	fmt.Println(user)
}
