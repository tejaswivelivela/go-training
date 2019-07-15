package main

import (
	"fmt"
)

func main() {
	var r rune
	fmt.Print("Enter the ASCII Character: ")
	fmt.Scanln(&r)
	ascii := int(r)
	fmt.Println("Converted ASCII Value:", ascii)
}
