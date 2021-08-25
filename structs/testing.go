package main

import (
	"fmt"
)

type iCalculator interface{
	multiplication(x int, y int) int
}

type sNormalCalc struct{
	id int
	operation iCalculator
}

func (n sNormalCalc) multiplication (x int, y int) int{
	n.id = 77
	return x * y
}
func (n *sNormalCalc) multiplication2 (x int, y int) int{
	n.id = 77
	return x * y
}

func main() {
	//var snc *sNormalCalc
	//res := snc.operation.multiplication2(5, 5)
	res := &sNormalCalc{id: 1, operation: {multiplication(5,5)}}
	fmt.Println("res:", res)

	var snc2 sNormalCalc
	res2 := snc2.multiplication(6,6)
	fmt.Println("res2:", res2)
}

