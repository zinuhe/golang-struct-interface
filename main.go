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
	var so interfaces.SimpleOperations
	fmt.Println("Sum(5+5):", so.Sum(5, 5))

	sso := interfaces.SimpleOperations{}
	// isc := interfaces.StandardCalculator

	x := interfaces.StandardCalculator{sso}

	interfaces.Calculator(x, 10, 10)
	// fmt.Println(interfaces.Calculator(x, 10, 10))
}
