package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurlyservers struct {
	XMLName xml.Name `xml:"servers"`
	//xml:"serverName"称为 strcut tag
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}
type Servers struct {
    XMLName xml.Name `xml:"servers"`
    Version string   `xml:"version,attr"`
    Svs     []server `xml:"server"`
}
type xmlServers struct {
    XMLName xml.Name `xml:"servers"`
    Version string   `xml:"version,attr"`
    Svs     []xmlServer `xml:"server"`
}
type xmlServer struct {
    ServerName string `xml:"serverName"`
    ServerIP   string `xml:"serverIP"`
}
/* 
xml包中提供了 Marshal 和 MarshalIndent 两个函数，来满足我们的需求。
这两个函数主要的区别是第二个函数会增加前缀和缩进，函数的定义如下所示：
func Marshal(v interface{}) ([]byte, error)
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
两个函数的第一个参数都是用来生成XML的结构定义类型数据，返回值都是XML数据流。
*/

func marshal() {
    v := &xmlServers{Version: "1"}
    v.Svs = append(v.Svs, xmlServer{"Shanghai_VPN", "127.0.0.1"})
    v.Svs = append(v.Svs, xmlServer{"Beijing_VPN", "127.0.0.2"})
  //  output, err := xml.Marshal(v)
    output, err := xml.MarshalIndent(v, "  ", "    ")
    if err != nil {
        fmt.Printf("error: %v\n", err)
    }
    /*os.Stdout.Write([]byte(xml.Header)) 这句代码的出现，
    是因为xml.MarshalIndent或者xml.Marshal输出的信息都是不带XML头的，
    为了生成正确的xml文件，我们使用了xml包预定义的Header变量。*/
    os.Stdout.Write([]byte(xml.Header))
    os.Stdout.Write(output)
}
func umMarshal(){
    file, err := os.Open("test.xml") // For read access.
	defer file.Close()
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := Recurlyservers{}
    /*
    用xml包的Unmarshal函数解析XML文件。
    func Unmarshal(data []byte, v interface{}) error
    data是接收的xml数据流；Interface（）是要输出的结构。目前只支持struct，slice，string。
    */
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v)
}
func main() {
	umMarshal()
    marshal()
}
