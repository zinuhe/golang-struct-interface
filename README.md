# golang-struct-interface

Examples of the use of structs, interfaces and methods.

go mod used
```
go mod init github.com/zinuhe/golang-struct-interface
```

1. [Structs](#struct)
2. [Interfaces](#interface)
3. [Methods](#method)
4. [Mocks](#mock)

<a name="struct"></a>
<h1>1. Structs</h1>
Go's structs are typed collections of fields. They are useful to group data together to form records

It’s idiomatic to encapsulate new struct creation in constructor functions

Structs are mutable

Ways to declare and use structs<br>
```go
type person struct {
    name string
    age  int
}
```

```go
var p person = person{         
		name: "John",         
		age: 30,     
	}
```

```go
var p person
```

```go
c := new(person)
// This allocates memory for all the fields, sets each of them to their zero value 
// and returns a pointer *person
```

```go
p := person{name: "Sean", age: 50}
```

```go
p := person{"Sean", 50}
```

```go
p := person{name: "jimmy"}
p.age = 42
```

```go
// An '&' prefix yields a pointer to the struct.
fmt.Println(&person{name: "Ann", age: 40})
```

```go
func newPerson(name string) *person {
    p := person{name: name}
    p.age = 42
    return &p
}

p := newPerson("Jon")
```


For reference<br>
[Go by example: Structs](https://gobyexample.com/structs)<br>
[golang-book](https://www.golang-book.com/books/intro/9)


<br>
<a name="interface"></a>
<h1>2. Interfaces</h1>

An interface is two things: 

It is a set of methods (Interfaces are named collections of method signatures), 
but it is also a type


To implement an interface in Go, we just need to implement all the methods in the interface.

If a variable has an interface type, then we can call methods that are in the named interface. 
A variable of type interface can hold any value which implements the interface. This property of interfaces is used to achieve polymorphism in Go.



The interface{} type

The interface{} type is the interface that has no methods. Since there is no implements keyword, all types implement at least zero methods, and satisfying an interface is done automatically, all types satisfy the empty interface. That means that if you write a function that takes an interface{} value as a parameter, you can supply that function with any value.

```go
package main

import (
    "fmt"
)

func PrintAll(vals []interface{}) {
    for _, val := range vals {
        fmt.Println(val)
    }
}

func main() {
    names := []string{"stanley", "david", "oscar"}
    vals := make([]interface{}, len(names))
    for i, v := range names {
        vals[i] = v
    }
    PrintAll(vals)
}
```
That’s pretty ugly.

<br><br>
```go
type Stringer interface {
    String() string
}

type Book struct {
    Title  string
    Author string
}

func (b Book) String() string {
    return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}
```
We say that something satisfies this interface (or implements this interface) if it has a method with the exact signature *String() string*.




For reference<br>
[Go by example: Interfaces](https://gobyexample.com/interfaces)<br>
[How to use interfaces in Go](https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go)<br>
[interfaces-explained](https://www.alexedwards.net/blog/interfaces-explained)


<br>
<a name="method"></a>
<h1>3. Methods</h1>

Go supports methods defined on struct types.

This area method has a receiver type of *rect.

```go
type rect struct {
    width, height int
}

func (r *rect) area() int {
    return r.width * r.height
}
```

Methods can be defined for either pointer or value receiver types. Here’s an example of a value receiver.
```go
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}
```

For reference [Go by example: Methods](https://gobyexample.com/methods)


<br>
<a name="mock"></a>
<h1>4. Mocking</h1>
[Check this](https://stackoverflow.com/questions/68100402/how-to-write-unit-test-in-golang-usng-echo-for-end-point-url-using-go-mock-gen-m)<br>
[Check](https://dev.to/techschoolguru/mock-db-for-testing-http-api-in-go-and-achieve-100-coverage-4pa9)<br>

