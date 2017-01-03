package main

import "fmt"

func main() {
	arraySliceTest0201()
}

//数组初始化的各种方式
func arraySliceTest0201() {
	//创建数组(声明长度)
	var array1 = [5]int{1, 2, 3}
	fmt.Printf("array1--- type:%T \n", array1)
	rangeIntPrint(array1[:])

	//创建数组(不声明长度)
	var array2 = [...]int{6, 7, 8}
	fmt.Printf("array2--- type:%T \n", array2)
	rangeIntPrint(array2[:])

	//创建数组切片
	var array3 = []int{9, 10, 11, 12}
	fmt.Printf("array3--- type:%T \n", array3)
	rangeIntPrint(array3)

	//创建数组(声明长度)，并仅初始化其中的部分元素
	var array4 = [5]string{3: "Chris", 4: "Ron"}
	fmt.Printf("array4--- type:%T \n", array4)
	rangeObjPrint(array4[:])

	//创建数组(不声明长度)，并仅初始化其中的部分元素，数组的长度将根据初始化的元素确定
	var array5 = [...]string{3: "Tom", 2: "Alice"}
	fmt.Printf("array5--- type:%T \n", array5)
	rangeObjPrint(array5[:])

	//创建数组切片，并仅初始化其中的部分元素，数组切片的len将根据初始化的元素确定
	var array6 = []string{4: "Smith", 2: "Alice"}
	fmt.Printf("array6--- type:%T \n", array6)
	rangeObjPrint(array6)
	//创建slice
	array7 := make([]int, 1024)
	for i := range array7 {
		array7[i] = i
	}
	//new	可以用new关键字申明
	p := new([10]int)
	p[9] = 12
	fmt.Println(p)
	//数组指针和指针数组
	a := [3]int{1, 2, 3}
	//这种是数组指针 我们看到可以直接输出指向数组的指针
	var p1 *[3]int = &a
	fmt.Println(p1)
	x, y := 1, 3
	b := [...]*int{&x, &y}
	//输出这样[0xc080000068 0xc080000070]的地址 这就是指针数组
	fmt.Println(b)
}

//输出整型数组切片
func rangeIntPrint(array []int) {
	for i, v := range array {
		fmt.Printf("index:%d  value:%d\n", i, v)
	}
}

//输出字符串数组切片
func rangeObjPrint(array []string) {
	for i, v := range array {
		fmt.Printf("index:%d  value:%s\n", i, v)
	}
}
