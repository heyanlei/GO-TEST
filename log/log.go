package main

import (
	"fmt"
	"log"
	"os"

	"github.com/toolkits/logger"
)

func logFile() {
	fileName := "xxx_debug.log"
	logFile, err := os.Create(fileName)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error !")
	}
	debugLog := log.New(logFile, "[Debug]", log.Llongfile)
	debugLog.Println("A debug message here")
	debugLog.SetPrefix("[Info]")
	debugLog.Println("A Info Message here1 ")
	debugLog.SetFlags(debugLog.Flags() | log.LstdFlags)
	debugLog.Println("A different prefix")

}

// log模块主要提供了3类接口。分别是 “Print 、Panic 、Fatal ”。
func main() {
	logger.Trace("777-%d", 88)
	logFile()
	/*
	   输出中的日期和时间是默认的格式，如果直接调用简单接口，其格式是固定的，可以通过 SetFlags 方法进行修改，
	   同时这里输出 内容的（传个log.Print的内容）前面和时间的后面是空的，这也是默认的行为，
	   我们可以通过添加前缀来表示其是一条"Warnning" 或者是一条"Debug"日志。通过使用 SetPrefix 可以设置该前缀。
	*/
	log.SetPrefix("debug--")
	//Print
	arr := []int{2, 3}
	log.SetFlags(log.Ldate)
	log.Print("Print array -0", arr, "\n")
	log.SetFlags(log.Lmicroseconds)
	log.Println("Println array-1", arr)
	log.SetFlags(log.Ltime)
	log.Printf("Printf array with item [%d,%d]-2\n", arr[0], arr[1])
	log.SetFlags(log.LstdFlags)
	log.Printf("Printf array with item [%d,%d]-3\n", arr[0], arr[1])
	//log.PanicXxx
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Just comming recover")
			fmt.Println("e from recover is :", e)
			fmt.Println("After recover")
		}
	}()
	log.Panic("Print array ", arr, "\n")
	//Fatal
	//log.Fatalf("this is a log %s-%d", "test", 100)
	//log.Fatalln("hello world!")
}
