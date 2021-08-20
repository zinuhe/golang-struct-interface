package main

import (
    "fmt"
)
type s interface{ 
    getSecond()
}

type As struct{
    s
}

type bs struct{

}

func (ss *As) GetFirst() {
    ss.getSecond()
}

func (ss bs) getSecond() {
    fmt.Println("--get_second--")
}

func main() {
  fmt.Printf("As{bs{}}: %+v \n\n", As{bs{}})

  a1 := bs{}
  a1.getSecond()
}
