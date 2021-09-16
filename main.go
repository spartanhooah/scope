package main

import (
	"fmt"
	"myapp/packageone"
)

var one = "One"

func main() {
	fmt.Println(one)
	myFunc()

	newString := packageone.PublicVar
	fmt.Println("From packageone:", newString)

	packageone.Exported()
}

func myFunc() {
	fmt.Println(one)
}