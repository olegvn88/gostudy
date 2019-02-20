package main

import (
	"fmt"
)

//docs for main method
func main() {

	var symbol rune = 'a'
	var autoSymbol = 'a'
	uautoSymbol := '\u0098'
	fmt.Println(symbol, autoSymbol, uautoSymbol)

	str1 := "Hello world ৺ ↲"
	fmt.Println("en:", str1, len(str1))

	for index, runeValue := range str1 {
		fmt.Printf("%#U at position %d\n", runeValue, index)
	}

	str2 := "привет мир ৺ ↲"
	fmt.Println("en:", str2, len(str2))

	for index, runeValue := range str2 {
		fmt.Printf("%#U at position %d\n", runeValue, index)
	}

	bin := []byte(str2)
	fmt.Println(bin, len(bin))
	for index, runeValue := range bin {
		fmt.Printf(" raw binary idx: %v, oct: %v, hex: %x\n", index, runeValue, runeValue)
	}
}
