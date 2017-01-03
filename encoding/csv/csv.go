package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func read() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(record) // record has the type []string
	}

}
func write() {
	f, err := os.Create("test.txt") //创建文件
	if err != nil {
		panic(err)
	}
	defer f.Close()
	w := csv.NewWriter(f) //创建一个新的写入文件流
	data := [][]string{
		{"1", "中国", "23"},
		{"2", "美国", "23"},
		{"3", "bb", "23"},
		{"4", "bb", "23"},
		{"5", "bb", "23"},
	}
	w.WriteAll(data) //写入数据
	w.Flush()
}
func main() {
	write()
	read()
}
