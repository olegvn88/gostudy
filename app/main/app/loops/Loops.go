package loomainps

import (
	"fmt"
)

func main() {
	//for {
	//	fmt.Println("Loop while(true)")
	//}

	sl1 := []int{3, 4, 5, 6, 7, 8}
	value := 0
	idx := 0

	for idx < 4 {
		if idx < 2 {
			idx++
			continue
		}
		value = sl1[idx]
		idx++
		fmt.Println("while-style loop, idx:", idx, "value:", value)
	}

	for i := 0; i < len(sl1); i++ {
		fmt.Println("c-style loop", i, sl1[i])
	}

	//конструкция range
	for _, value := range sl1 { // цикл по индексам
		fmt.Println(value)
	}

	mm := map[string]string{
		"firstName": "Oleg",
		"lastName":  "Nesterov",
	}

	for key, val := range mm {
		fmt.Println(key, val)
	}

	switch mm["firstName"] {
	case "Oleg", "Dima":
		fmt.Println("firstName is Oleg or Dima")
	case "Petr":
		fmt.Println("fisrtName is Petr")
		fallthrough
	default:
		fmt.Println("Swich some other name")
	}

	fmt.Println("=======")
MyLoop:
	for key, val := range mm {
		fmt.Println(key, val)
		switch {
		case key == "firstName" && val == "Oleg":
			fmt.Println("switch - break loop here")
			break MyLoop
		}
	}
}
