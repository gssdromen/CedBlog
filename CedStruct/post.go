package CedStruct

import (
	"CedBlog/CedUtils"
	// "fmt"
	// "strings"
	"time"
	"os"
)

type Post struct {
	title       string
	subtitle    string
	create_time time.Time
	author      string
	path        string
	fileName    string
	tags        string
}

func (p *Post) New(name string) (err error) {
	if !CedUtils.Exist("public") {
		os.Mkdir("public", 0755)
	}
	// 创建文件失败error
	file, err := os.OpenFile("public/" +name, os.O_WRONLY|os.O_CREATE, 0755)
	defer file.Close()
	if err != nil {
		return err
	}
	p.path = name
	p.fileName = name
	_, err = file.Write([]byte(p.fileName))
	return nil
}