package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// User : Comment for exported type
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (user *User) greeting() string {
	return "user name: " + user.Name + " and user age:" + strconv.Itoa(user.Age)
}

func main() {
	// parse user
	// userStr := `{"name":"nghia","age":20}`
	// var user User
	// err := json.Unmarshal([]byte(userStr), &user)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(user)

	// parse list user
	listUserStr := `[{"name":"nghia","age":23},{"name":"nguyen","age":25}]`
	var listUser []User
	if err := json.Unmarshal([]byte(listUserStr), &listUser); err != nil {
		fmt.Println(err)
	}
	for _, user := range listUser {
		fmt.Println(user.greeting())
	}

}
