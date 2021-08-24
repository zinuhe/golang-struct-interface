package interfaces

type Calculator interface {
	sum(a int, b int) int
	subtraction(a int, b int) int
	multiplication(a int, b int) int
}

type Operators struct{}

func (o Operators) Sum(a int, b int) int {
	return a + b
}

func (o Operators) subtraction(a int, b int) int {
	return a - b
}

func (o Operators) multiplication(a int, b int) int {
	return a * b
}

// func main(){
//     var x Operators
//     fmt.Println("Sum(5+5):", x.sum(5, 5))
// }
