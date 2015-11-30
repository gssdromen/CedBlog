package main

import (
	"CedBlog/CedStruct"
	"CedBlog/CedUtils"
	"bytes"
	"fmt"
	"github.com/russross/blackfriday"
	"html"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Go(c *CedStruct.Config) error {
	// 清空原有的public文件夹
	fileNames, err := CedUtils.FilesInDir(c.PublicDir)
	if err != nil {
		return err
	}
	for i := 0; i < len(fileNames); i++ {
		CedUtils.Remove(filepath.Join(c.PublicDir, fileNames[i]))
	}
	// 解析所有markdown到html
	fileNames, err = CedUtils.FilesInDir(c.ContentDir)
	if err != nil {
		return err
	}
	for i := 0; i < len(fileNames); i++ {
		// 解析
		fileName := fileNames[i]
		file, err := os.Open(filepath.Join(c.ContentDir, fileName))
		if err != nil {
			continue
		}
		data, err := ioutil.ReadAll(file)
		output := blackfriday.MarkdownCommon(data)
		file.Close()
		var doc bytes.Buffer
		template, _ := template.ParseFiles(filepath.Join(c.TemplateDir, "index.tmpl"),
			filepath.Join(c.TemplateDir, "root.tmpl"))
		template.Execute(&doc, string(output))
		// 保存
		fmt.Println(doc.String())
		htmlFileName := strings.Split(fileName, ".")[0]
		file, err = os.Create(filepath.Join(c.PublicDir, htmlFileName+".html"))
		if err != nil {
			return err
		}
		file.WriteString(html.UnescapeString(doc.String()))
		file.Close()
	}
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
