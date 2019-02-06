package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

//func main() {
//	//WordCount("My name is Oleg. What is you name? My name is Anastasia")
//	WordCount("I ate a donut . Then I ate another donut a a.")
//}
func WordCount(s string) map[string]int {
	var r = strings.Split(s, " ")
	countMap := map[string]int{}
	for i := 0; i < len(r); i++ {
		var count = 0
		for j := 0; j < len(r); j++ {
			if r[i] == r[j] {
				count++
				countMap[r[i]] = count
			}
		}
	}
	return countMap
}

func main() {
	wc.Test(WordCount)
}
