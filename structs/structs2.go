package structs

import "fmt"

type Customer struct {
}

func (Customer) Write(s string) {
	fmt.Println("s:", s)
}

// func main() {
//   cust := Customer{}
//   cust.Write("-message-")
// }
