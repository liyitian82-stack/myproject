package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("./main.go")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	var strSlice []byte
	var tempSlice = make([]byte, 150)
	for {
		n, err := file.Read(tempSlice)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println("读取失败")
			return
		}
		fmt.Printf("读取到了%v个字节", n)
		strSlice = append(strSlice, tempSlice[:n]...)
	}

	fmt.Println(string(strSlice))
}
