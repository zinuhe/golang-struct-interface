package main

import "fmt"

type s interface{ 
    getSecond(msg string)
}

type As struct{
    s
}

type bs struct{
}

func (ss As) GetFirst(msg string) {
    ss.getSecond(msg)
}

func (ss bs) getSecond(msg string) {
    fmt.Println("--get_second:", msg)
}

func main() {
  fmt.Printf("As{bs{}}: %+v \n\n", As{bs{}})

  b1 := bs{}
  b1.getSecond("b1")

  a1 := As{bs{}}
  a1.GetFirst("a1")
}

//https://stackoverflow.com/questions/42839517/how-to-mock-only-specific-method-in-golang
