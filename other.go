package main

import (
    "fmt"
    "encoding/json"
)

type response struct {
    SessionState string `json:"session_state"`
    Scope        string `json:"scope"`
}


func main() {
    data := &response {
            SessionState: "__sessionState",
            Scope: "__scope",
    }

    jsonData, _ := json.Marshal(data)
    fmt.Println("*data:", *data)
    fmt.Println("jsonData:", jsonData)
    fmt.Println("string(jsonData):", string(jsonData))

    fmt.Println("===============================")

    out := response{}
    json.Unmarshal(jsonData, &out)
    fmt.Println("out:", out)
}
