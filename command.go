package main

import (
	"CedBlog/CedStruct"
	"CedBlog/CedUtils"
	"fmt"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"os"
)

func Go(c *CedStruct.Config) error {
	fileNames, err := CedUtils.FilesInDir(c.PublicDir)
	for i := 0; i < len(fileNames); i++ {
		fileName := fileNames[i]
		file, err := os.Open(c.PublicDir + fileName)
		if err != nil {
			continue
		}
		data, err := ioutil.ReadAll(file)
		output := blackfriday.MarkdownBasic(data)
		file.Close()
		fmt.Println("File Open")
		fmt.Println(string(output))
		fmt.Println("File Close")
	}
	// file, err := os.Open(c.PublicDir + )
	fmt.Println(fileNames)
	return err
}

func NewPost(name string) error {
	post := new(CedStruct.Post)
	err := post.New(name)
	if err != nil {
		return err
	}
	return nil
}

func NewPage(name string) error {
	return nil
}
