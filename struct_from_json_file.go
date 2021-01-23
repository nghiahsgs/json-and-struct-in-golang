package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// User : Comment for exported type
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var listUser []User

	if data, err := ioutil.ReadFile("data.json"); err != nil {
		log.Fatal(err)
	} else {
		if err := json.Unmarshal(data, &listUser); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(listUser)
	fmt.Println(len(listUser))
}
