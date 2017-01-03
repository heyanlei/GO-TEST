package main

//flag包实现了命令行参数的解析。
import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"
)

//2.如果你喜欢，也可以将flag绑定到一个变量，使用Var系列函数.通过flag.XxxVar()方法将flag绑定到一个变量，该种方式返回值类型：
var flagvar int

func test() {
	flag.IntVar(&flagvar, "name", 555, "help message for flagname")
}

//3 或者你可以自定义一个用于flag的类型（满足Value接口）并将该类型用于flag解析，
//只需要实现Value接口，但实际上，如果需要取值的话，需要实现Getter接口如下：
//对这种flag，默认值就是该变量的初始值。
/*
type Getter interface {
  Value
  Get(string) interface{}
}
type Value interface {
  String() string
  Set(string) error
}
*/
type interval []time.Duration

//实现String接口
func (i *interval) String() string {
	return fmt.Sprintf("%v", *i)
}

//实现Set接口,Set接口决定了如何解析flag的值
func (i *interval) Set(value string) error {
	//此处决定命令行是否可以设置多次-deltaT
	if len(*i) > 0 {
		return errors.New("interval flag already set")
	}
	for _, dt := range strings.Split(value, ",") {
		duration, err := time.ParseDuration(dt)
		if err != nil {
			return err
		}
		*i = append(*i, duration)
	}
	return nil
}

var intervalFlag interval

func init() {
	flag.Var(&intervalFlag, "deltaT", "comma-separated list of intervals to use between events")
}

//
func main() {
	//1. 使用flag.String(), Bool(), Int()等flag.Xxx()方法，该种方式返回一个相应的指针等函数注册flag，声明了一个整数flag，解析结果保存在*int指针ip里：
	ip := flag.Int("flagname", 888, "help message for flagname")

	test()
	flag.Parse()
	//1
	fmt.Println(*ip)
	//2
	fmt.Println(flagvar)
	//3
	fmt.Println(intervalFlag)
	//-deltaT 61m,72h,80s
}

/*
命令行参数的格式可以是：

-flag xxx （使用空格，一个 - 符号）
--flag xxx （使用空格，两个 - 符号）
-flag=xxx （使用等号，一个 - 符号）
--flag=xxx （使用等号，两个 - 符号）
*/
