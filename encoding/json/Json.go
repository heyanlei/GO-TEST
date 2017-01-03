package main

//Decoder从输入流解码json对象
import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const jsonStream = `
		{"Name": "Ed", "Text": "Knock knock."}
		{"Name": "Sam", "Text": "Who's there?"}
		{"Name": "Ed", "Text": "Go fmt."}
		{"Name": "Sam", "Text": "Go fmt who?"}
		{"Name": "Ed", "Text": "Go fmt yourself!"}
	`

//
type message struct {
	Name, Text string
}

func testDecoder() {
	//NewDecoder创建一个从r读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}

}

//其中把对象转换为JSON的方法（函数）为 json.Marshal()，其函数原型如下
//　　　　func Marshal(v  interface{}) ([]byte, error)
//func Marshal(v interface{}) ([]byte, error) Marshal函数返回v的json编码。
/*
结构体的值编码为json对象。每一个导出字段变成该对象的一个成员，除非：

- 字段的标签是"-"
- 字段是空值，而其标签指定了omitempty选项
空值是false、0、""、nil指针、nil接口、长度为0的数组、切片、映射。对象默认键字符串是结构体的字段名，但可以在结构体字段的标签里指定。结构体标签值里的"json"键为键名，后跟可选的逗号和选项，举例如下：

// 字段被本包忽略
Field int `json:"-"`
// 字段在json里的键为"myName"
Field int `json:"myName"`
// 字段在json里的键为"myName"且如果字段为空值将在对象中省略掉
Field int `json:"myName,omitempty"`
// 字段在json里的键为"Field"（默认值），但如果字段为空值会跳过；注意前导的逗号
Field int `json:",omitempty"`
"string"选项标记一个字段在编码json时应编码为字符串。它只适用于字符串、浮点数、整数类型的字段。这个额外水平的编码选项有时候会用于和javascript程序交互：

Int64String int64 `json:",string"`
如果键名是只含有unicode字符、数字、美元符号、百分号、连字符、下划线和斜杠的非空字符串，将使用它代替字段名。
*/
type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

//对象转换为JSON格式
func testMarshal() {
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}

/*
把 JSON 转换回对象的方法（函数）为 json.Unmarshal()，其函数原型如下
func Unmarshal(data [] byte, v interface{}) error
func Unmarshal
func Unmarshal(data []byte, v interface{}) error
Unmarshal函数解析json编码的数据并将结果存入v指向的值。

Unmarshal和Marshal做相反的操作，必要时申请映射、切片或指针，有如下的附加规则：

要将json数据解码写入一个指针，Unmarshal函数首先处理json数据是json字面值null的情况。此时，函数将指针设为nil；否则，函数将json数据解码写入指针指向的值；如果指针本身是nil，函数会先申请一个值并使指针指向它。

要将json数据解码写入一个结构体，函数会匹配输入对象的键和Marshal使用的键（结构体字段名或者它的标签指定的键名），优先选择精确的匹配，但也接受大小写不敏感的匹配。

要将json数据解码写入一个接口类型值，函数会将数据解码为如下类型写入接口：

Bool                   对应JSON布尔类型
float64                对应JSON数字类型
string                 对应JSON字符串类型
[]interface{}          对应JSON数组
map[string]interface{} 对应JSON对象
nil                    对应JSON的null
*/
var jsonBlob = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	]`)

type Animal struct {
	Name  string
	Order string
}

//json转换为json
func testUnMarshal() {
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
}

func main() {
	testDecoder()
	testMarshal()
	fmt.Println("")
	testUnMarshal()
}
