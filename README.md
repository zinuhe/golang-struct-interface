# golang-struct-interface

Examples of the use of structs, interfaces and methods.

go mod used
```
go mod init github.com/zinuhe/golang-struct-interface
```

<h1>Structs</h1>
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


For reference [Go by example: Structs](https://gobyexample.com/structs)

<br>
<h1>Interfaces</h1>

Interfaces are named collections of method signatures.

To implement an interface in Go, we just need to implement all the methods in the interface.

If a variable has an interface type, then we can call methods that are in the named interface. 
A variable of type interface can hold any value which implements the interface. This property of interfaces is used to achieve polymorphism in Go.

For reference [Go by example: Interfaces](https://gobyexample.com/interfaces)
[How to use interfaces in Go](https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go)


<br>
<h1>Methods</h1>

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
