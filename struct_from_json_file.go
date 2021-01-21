package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var listUser []User

	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(data, &listUser)

	fmt.Println(listUser)
	fmt.Println(len(listUser))
}
