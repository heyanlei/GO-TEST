package main

import (
	"file"
	"fmt"
	"strconv"
)

/*go关键字
break    default      func    interface    select
case     defer        go      map          struct
chan     else         goto    package      switch
const    fallthrough  if      range        type
continue for          import  return       var
*/
var (
	g bool
	// 和声明array一样，只是少了长度
	fslice []inf
	// 声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用make初始化
	numbers map[string]int
)

//常量声明
const (
	first = iota
	second
)

type printInf interface {
	printPerson()
}
type person struct {
	name string
	age  int
}

// 通过这个方法 Human 实现了 fmt.Stringer
func (t *person) String() string {
	return "❰" + t.name + " - " + strconv.Itoa(t.age) + "❱"
}
func (t *person) printPerson() {
	fmt.Println(*t)
}

//element.(type)语法不能在switch外的任何逻辑里面使用，fallthrough不能在switch type中使用
func testInf(inf interface{}) {
	switch inf.(type) {
	case int:
		b, ok := inf.(int)
		if ok {
			fmt.Printf(" is an int--%d\n", b)
		}

	case string:
		s, ok := inf.(string)
		if ok {
			fmt.Printf("is a string--%s\n", s)
		}
	case person:
		p, ok := inf.(person)
		if ok {
			fmt.Println(p)
		}
	default:
		fmt.Println("is of a different type")
	}
	/*
		使用comma-ok
		Go语言里面有一个语法，可以直接判断是否是该类型的变量：
		value, ok = element.(T)，这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。
		如果element里面确实存储了T类型的数值，那么ok返回true，否则返回false。
	*/
	/*
		if _, ok := inf.(int); ok {
			fmt.Printf(" is an int\n")
		} else if _, ok := inf.(string); ok {
			fmt.Printf("is a string\n")
		} else if _, ok := inf.(person); ok {
			fmt.Printf("is a Person \n")
		} else {
			fmt.Println("is of a different type")
		}
	*/
}
func changeValue(value *bool) {
	*value = !*value
}
func close() {
	fmt.Println("end!")
}

type inf interface{}
type list []inf

func configExists(cfg string) bool {
	if !file.IsExist(cfg) {
		return false
	}
	return true
}
func main() {
	if configExists("c:/go/tmp.txt") {
		fmt.Println("c:/gotmp.txt--true")
	} else {
		fmt.Println("c:/go/tmp.txt--false")
	}
	//new使用
	pInt := new(int)
	*pInt = 999
	testInf(pInt)
	v := person{"JSON", 66}
	fmt.Println(v)
	fmt.Println("start!")
	g = true
	defer close()
	//
	list := make(list, 3)
	list[0] = second  //an int
	list[1] = "Hello" //a string
	list[2] = person{"Dennis", 70}
	//追加数组元素
	list = append(list, 55)
BREAK:
	for {
		if g {
			changeValue(&g)
			fmt.Println("continue")
			continue BREAK
		} else {
			fmt.Println("break")
			break BREAK
		}
	}

GOTO:
	for _, value := range list {
		testInf(value)
		if !g {
			changeValue(&g)
			fmt.Println("goto")
			goto GOTO
		}
	}

	// 指向数组的第1个元素开始，并到第2个元素结束，
	//fslice = list[1:2]
	//slice的默认开始位置是0，ar[:n]等价于ar[0:n]
	fslice = list[:2]
	// 另一种map的声明方式
	numbers := make(map[string]int)
	numbers["one"] = 1  //赋值
	numbers["ten"] = 10 //赋值
	numbers["three"] = 3
	fmt.Println("第三个数字是: ", numbers["three"])
	// 读取数据
	// 打印出来如:第三个数字是: 3
	testInf(fslice)
	/*
		channel通过操作符<-来接收和发送数据
		ch <- v    // 发送v到channel ch.
		v := <-ch  // 从ch中接收数据，并赋值给v
	*/
	cs := make(chan string, 10)
	go recvChan(cs)
	go sendChan(cs)
	select {
	case i := <-cs:
		// use i
		fmt.Println(i)
	default:
		// 当cs阻塞的时候执行这里
		fmt.Println("default")
	}
}
func sendChan(c chan string) {
	for i := 0; i < 2; i++ {
		c <- "hello chan"
	}

}
func recvChan(c chan string) {
	for i := 0; i < 2; i++ {
		recv := <-c
		fmt.Println(recv)
	}

}
