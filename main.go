package main

import (
	"CedBlog/CedStruct"
	"fmt"
)

func main() {
	fmt.Println("in")
	config := new(CedStruct.Config)
	err := config.Read()
	if err != nil {
		panic(err)
	}
}
