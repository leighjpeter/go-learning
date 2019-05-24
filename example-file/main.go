package main

import (
	"bufio"
	"fmt"
	"io"
	// "io/ioutil"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { // 已存在
		return true, nil
	}
	if os.IsNotExist(err) { //返回的err用IsNotExist判断，如果是true则表示不存在
		return false, nil
	}
	return false, err
}

func main() {
	filePath := "/Users/jessica/Dev/go/src/github.com/leighjpeter/go-learning/example-file/"
	fileName := "aa.txt"

	// 读文件
	file, err := os.Open(filePath + fileName)
	if err != nil {
		fmt.Println("open fail,err=", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	//循环读
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { // io.EOF 表示文件的末尾
			break
		}
		fmt.Println(str)
	}
	//os.O_WRONLY
	//os.O_RDONLY
	//os.O_RDWR
	//os.O_APPEND
	//os.O_CREATE
	//os.O_TRUNC
	//写文件
	file, err = os.OpenFile(filePath+fileName, os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("open fail,err=", err)
	}
	defer file.Close()
	str := "hello world\n"
	//写入时，使用带缓冲的*writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()

	//使用ioutil一次将整个文件读入到内存中
	// content, err := ioutil.ReadFile(filePath + fileName)
	// if err != nil {
	// 	fmt.Println("read fail,err=", err)
	// }
	// fmt.Println(string(content))

	// fmt.Println("read done")

}
