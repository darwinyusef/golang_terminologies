package main

import (
	"fmt"
)

func main() {
	var a [5]int
	a[0] = 2
	a[1] = 5
	a[2] = 22
	a[3] = 10
	a[4] = 8
	// b = [22,10]
	b := [...]interface{}{1, "Hola", 5, a[2:4], 5, "❤", "Mundo"}

	fmt.Println("Info A:", a[2:4])
	fmt.Println("Info B:", b[5:6])
	fmt.Println("slice a en B:", b[3])
}
