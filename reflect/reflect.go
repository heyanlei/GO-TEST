package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"
)

type MyStruct struct {
	name string
}
type Home struct {
	i int `nljb:"100"`
}

//
type Person struct {
	Name string
	Id   int32
	Age  int32
}

func (person Person) Eat(food string) {
	fmt.Println("eat...", food)
}
func getInterface() {
	person := Person{Name: "777", Id: 555, Age: 87}
	// Type
	t := reflect.TypeOf(person)
	// output: struct Person 3 1
	fmt.Println(t.Kind(), t.Name(), t.NumField(), t.NumMethod())

	// Value
	v1 := reflect.ValueOf(person)
	v1.MethodByName("Eat").Call([]reflect.Value{reflect.ValueOf("666")})
	//v1.FieldByName("Age").SetInt(120)---ERROR
	// output: struct {zhulx 1 28} 3 1 false
	fmt.Println(v1.Kind(), v1.Interface(), v1.NumField(), v1.NumMethod(), v1.CanSet())

	// 指针的Value
	v2 := reflect.ValueOf(&person)
	// Interface方法用于获取具体的值，CanSet表示是否可以修改原始类型的值，如果返回false，则不能调用CanXXX()类型的方法用来设置字段的值，也不能调用Call方法用来调用方法。
	// output: ptr &{zhulx 1 28} 3 1 false true
	fmt.Println(v2.Kind(), v2.Interface(), v2.Elem().NumField(), v2.NumMethod(), v2.CanSet(), v2.Elem().CanSet())
	// 可以通过传入指针类型对象来对原始对象的字段进行修改，或者调用方法，如
	// Field Set
	v3 := reflect.ValueOf(&person).Elem()
	v3.FieldByName("Age").SetInt(30)
	// output: 30
	fmt.Println(person.Age)
	// 也可以调用Call方法来调用具体某个方法，如
	// Method Call
	v4 := reflect.ValueOf(&person).Elem()
	// output: eat... apple
	v4.MethodByName("Eat").Call([]reflect.Value{reflect.ValueOf("apple")})
}

//获取 Struct 对象的 Tag
func getStructTag() {
	home := new(Home)
	home.i = 5
	// ValueOf用于获取Value信息
	rcvr := reflect.ValueOf(home)
	home1 := Home{i: 7777}

	t := reflect.Indirect(reflect.ValueOf(home1)).Type()
	fmt.Println("t type is " + t.String())
	//Indirec针对指针
	typ := reflect.Indirect(rcvr).Type()

	//ind()方法用于表示是哪种类型的，如struct, int, string等。Name方法表示类名，NumField表示字段的个数，NumMethod表示方法的个数。
	fmt.Println(t.Kind(), t.Name(), t.NumField(), t.NumMethod())

	t1 := reflect.ValueOf(home1).Type()
	fmt.Println(t1.Kind(), t1.Name(), t.NumField(), t1.NumMethod())
	//对于value的函数操作
	fmt.Println(rcvr.Kind(), rcvr.Elem().NumField(), rcvr.NumMethod())

	fmt.Println(typ.Kind().String())
	x := typ.NumField()
	for i := 0; i < x; i++ {
		nljb := typ.Field(0).Tag.Get("nljb")
		fmt.Println(nljb)
	}
}
func (this *MyStruct) GetName(str string) string {
	this.name = str
	return this.name
}

////////////////////////////////
func Info(o interface{}) {
	t := reflect.TypeOf(o)         //获取接口的类型
	fmt.Println("Type:", t.Name()) //t.Name() 获取接口的名称
	fmt.Println(t.Kind())
	if t.Kind() != reflect.Struct { //通过Kind()来判断反射出的类型是否为需要的类型
		fmt.Println("err: type invalid!")
		return
	}

	v := reflect.ValueOf(o) //获取接口的值类型
	//判断是否接口
	/*
		Interfaces don't have fields,
		they only define a method set of the value they contain.
		When reflecting on an interface, you can extract the value with Value.Elem().or reflect.Indirect()
	*/
	/**var vf interface{}
	var tf = reflect.ValueOf(&vf).Type().Elem()
	fmt.Println(tf.Kind() == reflect.Interface)
	*/
	//
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ { //NumField取出这个接口所有的字段数量
		f := t.Field(i) //取得结构体的第i个字段
		var tf = reflect.ValueOf(&f).Type()
		fmt.Println(tf.Kind() == reflect.Interface)
		if v.Field(i).CanInterface() {
			val := v.Field(i).Interface()

			fmt.Printf("interface false, %6s: %v = %v\n", f.Name, f.Type, val) //第i个字段的名称,类型,值
		} else {
			val := "99"
			fmt.Printf("interface false,%6s: %v = %v\n", f.Name, f.Type, val) //第i个字段的名称,类型,值
		}
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type) //获取方法的名称和类型
	}
}

/*
reflect包的几个使用场景：
1. 遍历结构体字段名（避免代码的硬编码）
2. 调用结构体方法（自动映射）
3. 获取结构体的tag标记的值（json/xml转换）
*/
type Db struct {
	Port int
	Host string
	pw   string
}

type Conf struct {
	Op       *string `json:"jsonop" xml:"xmlOpName"`
	Charlist *string
	Length   *int
	Num      *int
	Output   *string
	Input    *string
	hidden   *string
	Db
}

func (this Conf) SayOp(subname string) string {
	return *this.Op + subname
}

func (this Conf) getDbConf() Db {
	return this.Db
}
func testGo() {

	// 创建Conf实例
	conf := Conf{}

	opName := "create"
	conf.Op = &opName
	conf.Port = 3308

	fmt.Printf("conf.Port=%d\n\n", conf.Port)

	// 结构信息
	t := reflect.TypeOf(conf)
	// 值信息
	v := reflect.ValueOf(conf)

	printStructField(&t)

	callMethod(&v, "SayOp", []interface{}{" Db"})

	// panic: reflect: Call of unexported method
	//callMethod(&v, "getDbConf", []interface{}{})

	getTag(&t, "Op", "json")
	getTag(&t, "Op", "xml")
	getTag(&t, "nofield", "json")
}

// 场景1：遍历结构体字段名
func printStructField(t *reflect.Type) {
	fieldNum := (*t).NumField()
	for i := 0; i < fieldNum; i++ {
		fmt.Printf("conf's field: %s--type--%s value---%s\n", (*t).Field(i).Name)
	}
	fmt.Println("8888888888888888888")
}

// 场景2：调用结构体方法
func callMethod(v *reflect.Value, method string, params []interface{}) {
	// 字符串方法调用，且能找到实例conf的属性.Op
	f := (*v).MethodByName(method)
	if f.IsValid() {
		args := make([]reflect.Value, len(params))
		for k, param := range params {
			args[k] = reflect.ValueOf(param)
		}
		// 调用
		ret := f.Call(args)
		if ret[0].Kind() == reflect.String {
			fmt.Printf("%s Called result: %s\n", method, ret[0].String())
		}
	} else {
		fmt.Println("can't call " + method)
	}
	fmt.Println("")
}

// 场景3：获取结构体的tag标记
func getTag(t *reflect.Type, field string, tagName string) {
	var (
		tagVal string
		err    error
	)
	fieldVal, ok := (*t).FieldByName(field)
	if ok {
		tagVal = fieldVal.Tag.Get(tagName)
	} else {
		err = errors.New("no field named:" + field)
	}

	fmt.Printf("get struct[%s] tag[%s]: %s, error:%v\n", field, tagName, tagVal, err)
	fmt.Println("")
}

////////////////////////////////////////////////////////
func main() {
	//panic处理
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容
		}
		os.Exit(1)
	}()
	conf := Conf{}
	opName := "create"
	conf.Op = &opName
	conf.Port = 3308
	Info(conf)
	return
	//
	testGo()

	//reflect server case
	// server := NewServer()
	// fmt.Println(server.Register(new(Hello)))
	// server.Start(":8080")
	//get struct tag
	getInterface()
	fmt.Println("/////////////////////////////////////")
	getStructTag()
	// 备注: reflect.Indirect -> 如果是指针则返回 Elem()
	// 首先，reflect包有两个数据类型我们必须知道，一个是Type，一个是Value。
	// Type就是定义的类型的一个数据类型，Value是值的类型
	// 对象
	s := "this is string"
	// 获取对象类型 (string)
	fmt.Println(reflect.TypeOf(s))
	// 获取对象值 (this is string)
	fmt.Println(reflect.ValueOf(s))
	// 对象
	var x float64 = 3.4
	// 获取对象值 (<float64 Value>)
	fmt.Println(reflect.ValueOf(x))
	// 对象
	a := &MyStruct{name: "nljb"}
	// 返回对象的方法的数量 (1)
	fmt.Println(reflect.TypeOf(a).NumMethod())
	// 遍历对象中的方法
	for m := 0; m < reflect.TypeOf(a).NumMethod(); m++ {
		method := reflect.TypeOf(a).Method(m)
		fmt.Println(method.Type)         // func(*main.MyStruct) string
		fmt.Println(method.Name)         // GetName
		fmt.Println(method.Type.NumIn()) // 参数个数
		fmt.Println(method.Type.In(1))   // 参数类型
	}
	// 获取对象值 (<*main.MyStruct Value>)
	fmt.Println(reflect.ValueOf(a))
	// 获取对象名称
	fmt.Println(reflect.Indirect(reflect.ValueOf(a)).Type().Name())
	// 参数
	i := "Hello"
	v := make([]reflect.Value, 0)
	v = append(v, reflect.ValueOf(i))
	// 通过对象值中的方法名称调用方法 ([nljb]) (返回数组因为Go支持多值返回)
	fmt.Println(reflect.ValueOf(a).MethodByName("GetName").Call(v))
	// 通过对值中的子对象名称获取值 (nljb)
	fmt.Println(reflect.Indirect(reflect.ValueOf(a)).FieldByName("name"))
	// 是否可以改变这个值 (false)
	fmt.Println(reflect.Indirect(reflect.ValueOf(a)).FieldByName("name").CanSet())
	// 是否可以改变这个值 (true)
	fmt.Println(reflect.Indirect(reflect.ValueOf(&(a.name))).CanSet())
	// 不可改变 (false)
	fmt.Println(reflect.Indirect(reflect.ValueOf(s)).CanSet())
	// 可以改变
	// reflect.Indirect(reflect.ValueOf(&s)).SetString("jbnl")
	fmt.Println(reflect.Indirect(reflect.ValueOf(&s)).CanSet())
}

//reflect case
type Server struct {
	name    string
	rcvr    reflect.Value
	typ     reflect.Type
	methods map[string]*Method
}
type Method struct {
	method reflect.Method
	json   bool
}

func NewServer() *Server {
	server := new(Server)
	server.methods = make(map[string]*Method)
	return server
}
func (this *Server) Start(port string) error {
	return http.ListenAndServe(port, this)
}
func (this *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for mname, mmethod := range this.methods {
		if strings.ToLower("/"+this.name+"."+mname) == r.URL.Path {
			if mmethod.json {
				returnValues := mmethod.method.Func.Call(
					[]reflect.Value{this.rcvr, reflect.ValueOf(w), reflect.ValueOf(r)})
				content := returnValues[0].Interface()
				if content != nil {
					w.WriteHeader(500)
				}
			} else {
				mmethod.method.Func.Call(
					[]reflect.Value{this.rcvr, reflect.ValueOf(w), reflect.ValueOf(r)})
			}
		}
	}
}

/*
   func (this *Hello) JsonHello(r *http.Request) {}
   func (this *Hello) Hello(w http.ResponseWriter, r *http.Request) {}
*/
func (this *Server) Register(rcvr interface{}) error {
	this.typ = reflect.TypeOf(rcvr)
	this.rcvr = reflect.ValueOf(rcvr)
	this.name = reflect.Indirect(this.rcvr).Type().Name()
	if this.name == "" {
		return fmt.Errorf("no service name for type ", this.typ.String())
	}
	for m := 0; m < this.typ.NumMethod(); m++ {
		method := this.typ.Method(m)
		mtype := method.Type
		mname := method.Name
		if strings.HasPrefix(mname, "Json") {
			if mtype.NumIn() != 2 {
				return fmt.Errorf("method %s has wrong number of ins: %d", mname, mtype.NumIn())
			}
			arg := mtype.In(1)
			if arg.String() != "*http.Request" {
				return fmt.Errorf("%s argument type not exported: %s", mname, arg)
			}
			this.methods[mname] = &Method{method, true}
		} else {
			if mtype.NumIn() != 3 {
				return fmt.Errorf("method %s has wrong number of ins: %d", mname, mtype.NumIn())
			}
			reply := mtype.In(1)
			if reply.String() != "http.ResponseWriter" {
				return fmt.Errorf("%s argument type not exported: %s", mname, reply)
			}
			arg := mtype.In(2)
			if arg.String() != "*http.Request" {
				return fmt.Errorf("%s argument type not exported: %s", mname, arg)
			}
			this.methods[mname] = &Method{method, false}
		}
	}
	return nil
}

// ... //
type Hello struct {
}

func (this *Hello) Print(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	w.Write([]byte("print"))
	return nil
}
func (this *Hello) Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
func (this *Hello) JsonHello(r *http.Request) {
}
