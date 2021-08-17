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

<h1>Interfaces</h1>

A variable of type interface can hold any value which implements the interface. This property of interfaces is used to achieve polymorphism in Go.

For reference [Go by example: Interfaces](https://gobyexample.com/interfaces)


<h1>Methods</h1>

