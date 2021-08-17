// https://gobyexample.com/structs
// https://medium.com/rungo/structures-in-go-76377cc106a2

package structs

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func newPerson(fName string) *person {
	p := person{firstName: fName}
	p.age = 42
	return &p
}

func DisplayStructs() {
	fmt.Println("1:", person{firstName: "Alice", age: 30})

	fmt.Println(person{"John", "Snow", 30})

	fmt.Println(person{firstName: "Fred"})

	fmt.Println(&person{firstName: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := person{firstName: "Sean", age: 50}
	fmt.Println(s.firstName)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
