package main

import (
	"fmt"
	"github.com/zinuhe/golang-struct-interface/polymorphism"
	"github.com/zinuhe/golang-struct-interface/structs"
)

func main() {
	fmt.Println("Main function")

	fmt.Println("STRUCTS examples")
	structs.DisplayStructs()

	fmt.Println("POLYMORPHISM example")
	// polymorphism.Polymorphism()

	polymorphism.PurchaseCalculator()
}
