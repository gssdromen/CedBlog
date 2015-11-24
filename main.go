package main

import (
	"CedBlog/CedStruct"
	"fmt"
)

func main() {
	fmt.Println("hello")
	post := new(CedStruct.Post)
	err := post.New("bishisisi.md")
	if err != nil {
		panic(err)
	}
}
