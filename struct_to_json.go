package main

import (
	"encoding/json"
	"fmt"
)

// Book : book
type Book struct {
	Name   string `json:"name"`
	Year   int    `json:"year"`
	Author Author `json:"author"`
}

// Author : author
type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var author1 = Author{
		Name: "nghiahsgs",
		Age:  20,
	}
	listBook := []Book{
		Book{
			Name:   "day la ten cuon sach 1",
			Year:   2015,
			Author: author1,
		},
		Book{
			Name:   "day la ten cuon sach 2",
			Year:   2016,
			Author: author1,
		},
	}
	// or
	// listBook := make([]Book, <size>, <capacity>)
	// listBook := make([]Book, 0, 2)
	// listBook = append(listBook, book1)
	// listBook = append(listBook, book1)

	// fmt.Println(listBook)

	// listBookStr, err := json.Marshal(listBook)
	if listBookStr, err := json.MarshalIndent(listBook, "", "    "); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(listBookStr))
	}
}
