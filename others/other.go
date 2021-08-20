package others

import "fmt"

type Student struct {
	ID   string
	Name string
}

func main() {
	item := Student{"1", "Uno"}
	doSomethingWithThisParam(&item)
	fmt.Printf("%+v", item)
}

func doSomethingWithThisParam(item interface{}) {
	origin := item.(*Student)
	*origin = Student{"123", "John Doe"}
	item = origin
}
