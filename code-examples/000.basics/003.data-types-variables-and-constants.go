package main

import (
	"fmt"
)

const birthDate = "01/01/1999"

func main() {
	name, lastName := "John", "Doe"
	var wallet float64 = 541

	fmt.Println(name, lastName, birthDate, wallet)
}