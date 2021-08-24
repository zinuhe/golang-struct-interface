package main

import (
	"fmt"
	"github.com/zinuhe/golang-struct-interface/interfaces"
	"github.com/zinuhe/golang-struct-interface/polymorphism"
	"github.com/zinuhe/golang-struct-interface/structs"
)

func main() {
	fmt.Println("Main function")

	// From structs/structs.go
	fmt.Println("\n\nSTRUCTS example 1")
	structs.DisplayStructs()

	// From polymorphism/polymorphism.go
	fmt.Println("\n\nPOLYMORPHISM example 1")
	// polymorphism.Polymorphism()
	polymorphism.PurchaseCalculator()

	// From interfaces/interfaces2.go
	fmt.Println("\n\nINTERFACES example 2")
	var x interfaces.Operators
	fmt.Println("Sum(5+5):", x.Sum(5, 5))
}
