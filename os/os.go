package main

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"unsafe"
)

func b2S(buf []byte) string {
	return *(*string)(unsafe.Pointer(&buf))
}

func s2B(s *string) []byte {
	return *(*[]byte)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(s))))
}
func envOperations() {
	//hostname
	name, err := os.Hostname()
	if err != nil {
		fmt.Println("error!")
		return
	}
	fmt.Println(name)
	//返回所有环境变量
	strEnv := os.Environ()
	for k, v := range strEnv {
		fmt.Printf("key is %d,value is %s\n", k, v)
	}
	//
	envValue := os.Getenv("GOROOT")
	fmt.Println(envValue)
	//
	//更改当前目录
	strPath := "c:\\go"
	os.Chdir(strPath)
	//获得当前目录
	dir, _ := os.Getwd()
	fmt.Println(dir)
	//
	mapping := func(key string) string {
		m := make(map[string]string)
		m = map[string]string{
			"world": "kitty",
			"hello": "hi",
		}
		if m[key] != "" {
			return m[key]
		}
		return key
	}
	//Expand用mapping 函数指定的规则替换字符串中的${var}或者$var（注：变量之前必须有$符号）。
	//  hello,world，由于hello world之前没有$符号，则无法利用map规则进行转换
	s := "hello,world"
	//  hi,kitty finish，finish没有在map规则中，所以还是返回原来的值
	s1 := "$hello,$world $finish"
	fmt.Println(os.Expand(s, mapping))
	fmt.Println(os.Expand(s1, mapping))
	//判断c是否是一个路径分割符号，是的话返回true,否则返回false
	fmt.Println(os.IsPathSeparator('/'))
	e := os.Setenv("goenv", "go environment")
	a := os.Getenv("goenv")
	fmt.Println(e, a) //  <nil> go environment

}
func dirOperations() {
	var path string
	if os.IsPathSeparator('\\') {
		path = "\\"
	} else {
		path = "/"
	}
	pwd, _ := os.Getwd()
	err := os.Mkdir(pwd+path+"tmp", os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	/*
	   创建一个新目录，该目录是利用路径（包括绝对路径和相对路径）进行创建的，如果需要创建对应的父目录，
	   也一起进行创建，如果已经有了该目录，则不进行新的创建，当创建一个已经存在的目录时，不会报错.
	*/
	os.MkdirAll(pwd+path+"tmp"+path+"del", os.ModePerm)
	/*
		func Remove(name string) error           //删除文件或者目录
		func RemoveAll(path string) error　　//删除目录以及其子目录和文件，如果path不存在的话，返回nil
	*/
}

/*
建立文件函数：

func Create(name string) (file *File, err Error)
func NewFile(fd int, name string) *File

打开文件函数：

func Open(name string) (file *File, err Error)
func OpenFile(name string, flag int, perm uint32) (file *File, err Error)
*/
func fileOperations() {
	createFile()
	open()
	openFile()
}
func open() {
	file, err := os.Open("tmp.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
	content := [100]byte{}
	/*读文件函数：
	func (file *File) Read(b []byte) (n int, err Error)
	func (file *File) ReadAt(b []byte, off int64) (n int, err Error)
	*/
	len, err := file.Read(content[:])
	if err != nil {
		fmt.Println("error!")
	} else {
		fmt.Printf("len is -%d,content is -%s", len, b2S(content[:]))
	}
}
func createFile() {
	/*
		//创建一个文件，文件mode是0666(读写权限)，
		如果文件已经存在，则重新创建一个，原文件被覆盖，创建的新文件具有读写权限，
		能够备用与i/o操作．其相当于OpenFile的快捷操作，
		等同于OpenFile(name string, O_CREATE,0666)
	*/
	file, err := os.Create("tmpCreate.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
}

/*
func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
*/
func openFile() {
	file, err := os.OpenFile("tmp1.txt", os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		写文件函数：
		func (file *File) Write(b []byte) (n int, err Error)
		func (file *File) WriteAt(b []byte, off int64) (n int, err Error)
		func (file *File) WriteString(s string) (ret int, err Error)
	*/
	file.WriteString("Just a test!\r\n")
	file.Write([]byte("Just a test!\r\n"))
	defer file.Close()
}
func main() {
	output, err := exec.Command("ipconfig", "-all").Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(output))
	envOperations()
	//创建一个新目录，该目录具有FileMode权限，当创建一个已经存在的目录时会报错
	//dirOperations()
	fileOperations()
}
