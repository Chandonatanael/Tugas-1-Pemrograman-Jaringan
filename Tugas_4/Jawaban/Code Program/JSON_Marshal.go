package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastname"`
}

func main() {
	bytes, err := json.Marshal(Person{
		FirstName: "Jhon",
		LastName:  "Dow",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
