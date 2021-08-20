package main

import (
    "fmt"
)
type s interface{ 
    getSecond()
}

type A struct{
    s
}

type a struct{

}

func (s *A) GetFirst() {
    s.getSecond()
}

func (s a) getSecond() {
    fmt.Println("getSecond")
}

func main() {
  //Use a 
  fmt.Printf("A{a{}}: %+v", A{a{}})
}
