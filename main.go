package main

import (
	"CedBlog/CedStruct"
	"fmt"
)

func main() {
	config := new(CedStruct.Config)
	config = config.Read()
	fmt.Println(config.BaseUrl)
}
