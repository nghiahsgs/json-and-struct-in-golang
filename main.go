package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Name   string `json:"name"`
	Year   int    `json:"year"`
	Author Author `json:"author"`
}
type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var listBook []Book
	var author1 = Author{Name: "nghiahsgs", Age: 20}

	book1 := Book{Name: "day la ten cuon sach 1", Year: 2015, Author: author1}
	listBook = append(listBook, book1)

	book2 := Book{Name: "day la ten cuon sach 2", Year: 2016, Author: author1}
	listBook = append(listBook, book2)

	// fmt.Println(listBook)

	// listBookStr, err := json.Marshal(listBook)
	listBookStr, err := json.MarshalIndent(listBook, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(listBookStr))
}
