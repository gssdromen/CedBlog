package main

import (
	"CedBlog/CedStruct"
	"flag"
	"fmt"
)

func main() {
	config := new(CedStruct.Config)
	config = config.Read()

	flag.Parse()
	fmt.Print("参数为:  ")
	fmt.Println(flag.Args())
	if len(flag.Args()) > 0 {
		if flag.Args()[0] == "new_post" {
			NewPost(flag.Args()[1])
		}
		if flag.Args()[0] == "go" {
			err := Go(config)
			if err != nil {
				panic(err)
			}
		}
	}
}
