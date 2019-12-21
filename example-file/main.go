package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	// "strconv"
	// "strings"
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

type Page struct {
	Title string
	Body  []byte
}

func (page *Page) save() (err error) {
	return ioutil.WriteFile(page.Title, page.Body, 0666)
}

func (page *Page) load(title string) (err error) {
	page.Title = title
	page.Body, err = ioutil.ReadFile(page.Title)
	return err
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

	/*
		请给这个结构编写一个 save 方法，将 Title 作为文件名、Body作为文件内容，写入到文本文件中。
		再编写一个 load 函数，接收的参数是字符串 title，该函数读取出与 title 对应的文本文件。
		请使用 *Page 做为参数，因为这个结构可能相当巨大，
	*/
	page := Page{
		"Page.md",
		[]byte("# Page\n## Section1\nThis is section1."),
	}

	page.save()
	var new_page Page
	new_page.load("Page.md")
	fmt.Println(string(new_page.Body))
}
