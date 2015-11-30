package CedUtils

import (
	"os"
)

// 检查文件或目录是否存在
// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func Remove(filename string) error {
	if Exist(filename) {
		err := os.Remove(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func FilesInDir(dirPath string) ([]string, error) {
	if !Exist(dirPath) {
		return nil, nil
	}
	file, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}
	var result []string
	files, err := file.Readdir(-1)
	for i := 0; i < len(files); i++ {
		aFile := files[i]
		if !aFile.IsDir() {
			result = append(result, aFile.Name())
		}
	}
	return result, nil
}
