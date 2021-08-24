package interfaces

type StandardCalculator interface {
	sum(a int, b int) int
	subtraction(a int, b int) int
	multiplication(a int, b int) int
}

type SimpleOperations struct{}

func (o SimpleOperations) Sum(a int, b int) int {
	return a + b
}
func (o SimpleOperations) Subtraction(a int, b int) int {
	return a - b
}
func (o SimpleOperations) multiplication(a int, b int) int {
	return a * b
}

func Calculator(sc StandardCalculator, a int, b int) int {
	return sc.sum(a, b)
}

// func main(){
//     var x SimpleOperations
//     fmt.Println("Sum(5+5):", x.sum(5, 5))
// }
