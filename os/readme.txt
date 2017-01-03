os包中实现了平台无关的接口，设计向Unix风格，但是错误处理是go风格，当os包使用时，如果失败之后返回错误类型而不是错误数量．

os包中函数设计方式和Unix类似，下面来看一下．

func Chdir(dir string) error   //chdir将当前工作目录更改为dir目录．

func Getwd() (dir string, err error)    //获取当前目录，类似linux中的pwd
func Chmod(name string, mode FileMode) error     //更改文件的权限（读写执行，分为三类：all-group-owner）
func Chown(name string, uid, gid int) error  //更改文件拥有者owner
func Chtimes(name string, atime time.Time, mtime time.Time) error    //更改文件的访问时间和修改时间，atime表示访问时间，mtime表示更改时间
func Clearenv()    //清除所有环境变量（慎用）
func Environ() []string  //返回所有环境变量
func Exit(code int)     //系统退出，并返回code，其中０表示执行成功并退出，非０表示错误并退出，其中执行Exit后程序会直接退出，defer函数不会执行．
//Expand用mapping 函数指定的规则替换字符串中的${var}或者$var（注：变量之前必须有$符号）。比如，os.ExpandEnv(s)等效于os.Expand(s, os.Getenv)。
func Expand(s string, mapping func(string) string) string   
func Getenv(key string) string  //获取系统key的环境变量，如果没有环境变量就返回空
fmt.Println(os.Getenv("GOROOT")) // /home/software/go
func Geteuid() int  //获取调用者用户id
func Getgid() int   //获取调用者的组id
func Getpagesize() int　　　//获取底层系统内存页的数量
func Getpid() int　　　　//获取进程id
func Getppid() int             //获取调用者进程父id
func Hostname() (name string, err error)    //获取主机名
func IsExist(err error) bool    　　　　　 //返回一个布尔值，它指明err错误是否报告了一个文件或者目录已经存在。它被ErrExist和其它系统调用满足。
func IsNotExist(err error) bool　　　　　//返回一个布尔值，它指明err错误是否报告了一个文件或者目录不存在。它被ErrNotExist 和其它系统调用满足。
func IsPathSeparator(c uint8) bool         //判断c是否是一个路径分割符号，是的话返回true,否则返回false
func IsPermission(err error) bool   //判定err错误是否是权限错误。它被ErrPermission 和其它系统调用满足。
func Lchown(name string, uid, gid int) error　　　//改变了文件的gid和uid。如果文件是一个符号链接，它改变的链接自己。如果出错，则会是*PathError类型。
func Link(oldname, newname string) error       //创建一个从oldname指向newname的硬连接，对一个进行操作，则另外一个也会被修改．
func Mkdir(name string, perm FileMode) error　//创建一个新目录，该目录具有FileMode权限，当创建一个已经存在的目录时会报错
unc MkdirAll(path string, perm FileMode) error　//创建一个新目录，该目录是利用路径（包括绝对路径和相对路径）进行创建的，如果需要创建对应的父目录，也一起进行创建，如果已经有了该目录，则不进行新的创建，当创建一个已经存在的目录时，不会报错.
func NewSyscallError(syscall string, err error) error    //NewSyscallError返回一个SyscallError 错误，带有给出的系统调用名字和详细的错误信息。也就是说，如果err为空，则返回空
func Readlink(name string) (string, error)         //返回符号链接的目标。如果出错，将会是 *PathError类型。
func Remove(name string) error           //删除文件或者目录
func RemoveAll(path string) error　　//删除目录以及其子目录和文件，如果path不存在的话，返回nil
func Rename(oldpath, newpath string) error　　//重命名文件，如果oldpath不存在，则报错no such file or directory
func SameFile(fi1, fi2 FileInfo) bool　　　　　　//查看f1和f2这两个是否是同一个文件，如果再Unix系统，这意味着底层结构的device和inode完全一致，在其他系统上可能是基于文件绝对路径的．SameFile只适用于本文件包stat返回的状态，其他情况下都返回false
func Setenv(key, value string) error           //设定环境变量，经常与Getenv连用，用来设定环境变量的值
func Symlink(oldname, newname string) error　　　//创建一个newname作为oldname的符号连接，这是一个符号连接（Link是硬连接），与link的硬连接不同，利用Link创建的硬连接，则newname和oldname的file互不影响，一个文件删除，另外一个文件不受影响；但是利用SymLink创建的符号连接，其newname只是一个指向oldname文件的符号连接，当oldname　file删除之后，则newname的文件也就不能够继续使用．
func TempDir() string　　　　　　　　//创建临时目录用来存放临时文件，这个临时目录一般为/tmp
func Truncate(name string, size int64) error     //按照指定长度size将文件截断，如果这个文件是一个符号链接，则同时也改变其目标连接的长度，如果有错误，则返回一个错误．
文件操作：
type File
type File struct {
    // contains filtered or unexported fields
}

写程序离不了文件操作，这里总结下go语言文件操作。

一、建立与打开

建立文件函数：

func Create(name string) (file *File, err Error)
func NewFile(fd int, name string) *File

打开文件函数：

func Open(name string) (file *File, err Error)
func OpenFile(name string, flag int, perm uint32) (file *File, err Error)


二、写文件

写文件函数：

func (file *File) Write(b []byte) (n int, err Error)
func (file *File) WriteAt(b []byte, off int64) (n int, err Error)
func (file *File) WriteString(s string) (ret int, err Error)

三、读文件

读文件函数：

func (file *File) Read(b []byte) (n int, err Error)
func (file *File) ReadAt(b []byte, off int64) (n int, err Error)
四、删除文件

函数：

func Remove(name string) Error


File表示打开的文件描述符
func Create(name string) (file *File, err error)　　//创建一个文件，文件mode是0666(读写权限)，如果文件已经存在，则重新创建一个，原文件被覆盖，创建的新文件具有读写权限，能够备用与i/o操作．其相当于OpenFile的快捷操作，等同于OpenFile(name string, O_CREATE,0666)
func NewFile(fd uintptr, name string) *File　　　//根据文件描述符和名字创建一个新的文件
unc OpenFile(name string, flag int, perm FileMode) (file *File, err error)　//指定文件权限和打开方式打开name文件或者create文件，其中flag标志如下:

打开标记：
O_RDONLY：只读模式(read-only)
O_WRONLY：只写模式(write-only)
O_RDWR：读写模式(read-write)
O_APPEND：追加模式(append)
O_CREATE：文件不存在就创建(create a new file if none exists.)
O_EXCL：与 O_CREATE 一起用，构成一个新建文件的功能，它要求文件必须不存在(used with O_CREATE, file must not exist)
O_SYNC：同步方式打开，即不使用缓存，直接写入硬盘
O_TRUNC：打开并清空文件
至于操作权限perm，除非创建文件时才需要指定，不需要创建新文件时可以将其设定为０.虽然go语言给perm权限设定了很多的常量，但是习惯上也可以直接使用数字，如0666(具体含义和Unix系统的一致).
func Pipe() (r *File, w *File, err error)        //返回一对连接的文件，从r中读取写入w中的数据，即首先向w中写入数据，此时从r中变能够读取到写入w中的数据，Pipe()函数返回文件和该过程中产生的错误．
func (f *File) Chmod(mode FileMode) error　　　//更改文件权限，其等价与os.Chmod(name string,mode FileMode)error
func (f *File) Chown(uid, gid int) error                     //更改文件所有者，与os.Chown等价
func (f *File) Close() error　　　　　　　　　　//关闭文件，使其不能够再进行i/o操作，其经常和defer一起使用，用在创建或者打开某个文件之后，这样在程序退出前变能够自己关闭响应的已经打开的文件．
func (f *File) Fd() uintptr　　　//返回系统文件描述符，也叫做文件句柄．
func (f *File) Name() string　　//返回文件名字，与file.Stat().Name()等价
func (f *File) Read(b []byte) (n int, err error)　//将len(b)的数据从f中读取到b中，如果无错，则返回n和nil,否则返回读取的字节数n以及响应的错误
func (f *File) ReadAt(b []byte, off int64) (n int, err error)　//和Read类似，不过ReadAt指定开始读取的位置offset
func (f *File) Readdir(n int) (fi []FileInfo, err error)            
Readdir读取file指定的目录的内容，然后返回一个切片，它最多包含n个FileInfo值，这些值可能是按照目录顺序的Lstat返回的。接下来调用相同的文件会产生更多的FileInfos。

如果n>0，Readdir返回最多n个FileInfo结构。在这种情况下，如果Readdir返回一个空的切片，它将会返回一个非空的错误来解释原因。在目录的结尾，错误将会是io.EOF。

如果n<=0，Readdir返回目录的所有的FileInfo，用一个切片表示。在这种情况下，如果Readdir成功（读取直到目录的结尾），它会返回切片和一个空的错误。如果它在目录的结尾前遇到了一个错误，Readdir返回直到当前所读到的FIleInfo和一个非空的错误。


func (f *File) Readdirnames(n int) (names []string, err error)
Readdirnames读取并返回目录f里面的文件的名字切片。

如果n>0，Readdirnames返回最多n个名字。在这种情况下，如果Readdirnames返回一个空的切片，它会返回一个非空的错误来解释原因。在目录的结尾，错误为EOF。

如果n<0，Readdirnames返回目录下所有的文件的名字，用一个切片表示。在这种情况下，如果用一个切片表示成功（读取直到目录结尾），它返回切片和一个空的错误。如果在目录结尾之前遇到了一个错误，Readdirnames返回直到当前所读到的names和一个非空的错误。
func (f *File) Seek(offset int64, whence int) (ret int64, err error)　　　　//Seek设置下一次读或写操作的偏移量offset，根据whence来解析：0意味着相对于文件的原始位置，1意味着相对于当前偏移量，2意味着相对于文件结尾。它返回新的偏移量和错误（如果存在）。
func (f *File) Stat() (fi FileInfo, err error)　　//返回文件描述相关信息，包括大小，名字等．等价于os.Stat(filename string)
func (f *File) Sync() (err error)                        //同步操作，将当前存在内存中的文件内容写入硬盘．
func (f *File) Truncate(size int64) error        //类似  os.Truncate(name, size),，将文件进行截断
func (f *File) Write(b []byte) (n int, err error)　　//将b中的数据写入f文件
func (f *File) WriteAt(b []byte, off int64) (n int, err error)　//将b中数据写入f文件中，写入时从offset位置开始进行写入操作
func (f *File) WriteString(s string) (ret int, err error)　　　//将字符串s写入文件中
type FileInfo

type FileInfo interface {
	Name() string       //文件名字
	Size() int64        // length in bytes for regular files; system-dependent for others，文件大小
	Mode() FileMode     // file mode bits，文件权限
	ModTime() time.Time // modification time　文件更改时间
	IsDir() bool        // abbreviation for Mode().IsDir()　文件是否为目录
	Sys() interface{}   // underlying data source (can return nil)　　基础数据源
}
FileInfo经常被Stat和Lstat返回来描述一个文件

func Lstat(name string) (fi FileInfo, err error)      //返回描述文件的FileInfo信息。如果文件是符号链接，返回的FileInfo描述的符号链接。Lstat不会试着去追溯link。如果出错，将是 *PathError类型。
func Stat(name string) (fi FileInfo, err error)       //返回描述文件的FileInfo信息。如果出错，将是 *PathError类型。


type FileMode

type FileMode uint32

FileMode代表文件的模式和权限标志位。标志位在所有的操作系统有相同的定义，因此文件的信息可以从一个操作系统移动到另外一个操作系统。不是所有的标志位是用所有的系统。唯一要求的标志位是目录的ModeDir。

const (
	// The single letters are the abbreviations
	// used by the String method's formatting.
	ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
	ModeAppend                                     // a: append-only
	ModeExclusive                                  // l: exclusive use
	ModeTemporary                                  // T: temporary file (not backed up)
	ModeSymlink                                    // L: symbolic link
	ModeDevice                                     // D: device file
	ModeNamedPipe                                  // p: named pipe (FIFO)
	ModeSocket                                     // S: Unix domain socket
	ModeSetuid                                     // u: setuid
	ModeSetgid                                     // g: setgid
	ModeCharDevice                                 // c: Unix character device, when ModeDevice is set
	ModeSticky                                     // t: sticky

	// Mask for the type bits. For regular files, none will be set.
	ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice

	ModePerm FileMode = 0777 // permission bits
)

所定义的文件标志位最重要的位是FileMode。9个次重要的位是标准Unix rwxrwxrwx权限。这些位的值应该被认为公开API的一部分，可能用于连接协议或磁盘表示：它们必须不能被改变，尽管新的标志位有可能增加。

FileModel的方法主要用来进行判断和输出权限
func (m FileMode) IsDir() bool              //判断m是否是目录，也就是检查文件是否有设置的ModeDir位
func (m FileMode) IsRegular() bool　　//判断m是否是普通文件，也就是说检查m中是否有设置mode type
func (m FileMode) Perm() FileMode　　//返回m的权限位
func (m FileMode) String() string　　　　//返回m的字符串表示
type LinkError
type LinkError struct {
    Op  string
    Old string
    New string
    Err error
}
func (e *LinkError) Error() string　　　　//LinkError记录了一个在链接或者符号连接或者重命名的系统调用中发生的错误和引起错误的文件的路径。

type PathError

type PathError struct {
Op string
Path string
Err error
}

func (e *PathError) Error() string　　//返回一个有操作者，路径以及错误组成的字符串形式

进程相关操作：

type ProcAttr
type ProcAttr struct {
    Dir   string               //如果dir不是空，子进程在创建之前先进入该目录
    Env   []string　　　//如果Env不是空，则将里面的内容赋值给新进程的环境变量，如果他为空，则使用默认的环境变量
    Files []*File　　　  //Files指定新进程打开文件，前三个实体分别为标准输入，标准输出和标准错误输出，可以添加额外的实体，这取决于底层的操作系统，当进程开始时，文
　　　　　　　　　　//件对应的空实体将被关闭
    Sys   *syscall.SysProcAttr  //操作系统特定进程的属性，设置该值也许会导致你的程序在某些操作系统上无法运行或者编译
}

ProcAttr包含属性，这些属性将会被应用在被StartProcess启动的新进程上type Process
Process存储了通过StartProcess创建的进程信息。
type Process struct {
    Pid int
     handle uintptr 　　//处理指针
      isdone uint32        // 如果进程正在等待则该值非０，没有等待该值为０
 }
func FindProcess(pid int) (p *Process, err error)　　　　//通过进程pid查找运行的进程，返回相关进程信息及在该过程中遇到的errorfunc StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)  //StartProcess启动一个新的进程，其传入的name、argv和addr指定了程序、参数和属性；StartProcess是一个低层次的接口。os/exec包提供了高层次的接口；如果出错，将会是*PathError错误。func (p *Process) Kill() error　　　　　　　　　　　//杀死进程并直接退出func (p *Process) Release() error　　　　　　　　//释放进程p的所有资源，之后进程p便不能够再被使用，只有Wati没有被调用时，才需要调用Release释放资源
func (p *Process) Signal(sig Signal) error　　　　//发送一个信号给进程p, 在windows中没有实现发送中断interrupt
func (p *Process) Wait() (*ProcessState, error)　　//Wait等待进程退出，并返回进程状态和错误。Wait释放进程相关的资源。在大多数的系统上，进程p必须是当前进程的子进程，否则会返回一个错误。
type ProcessState     //ProcessState存储了Wait函数报告的进程信息
type ProcessState struct {
	pid    int                
	status syscall.WaitStatus 
	rusage *syscall.Rusage
}

func (p *ProcessState) Exited() bool　　　　　　// 判断程序是否已经退出
func (p *ProcessState) Pid() int                                //返回退出进程的pid
func (p *ProcessState) String() string                     //获取进程状态的字符串表示
func (p *ProcessState) Success() bool                    //判断程序是否成功退出，而Exited则仅仅只是判断其是否退出
func (p *ProcessState) Sys() interface{}                //返回有关进程的系统独立的退出信息。并将它转换为恰当的底层类型（比如Unix上的syscall.WaitStatus），主要用来获取进程退出相关资源。
func (p *ProcessState) SysUsage() interface{}         //SysUsage返回关于退出进程的系统独立的资源使用信息。主要用来获取进程使用系统资源
func (p *ProcessState) SystemTime() time.Duration      //获取退出进程的内核态cpu使用时间
func (p *ProcessState) UserTime() time.Duration     //返回退出进程和子进程的用户态CPU使用时间
type Signal

type Signal interface {
    String() string
    Signal() // 同其他字符串做区别
}
代表操作系统的信号．底层的实现是操作系统独立的：在Unix中是syscal.Signal．
var (
	Interrupt Signal = syscall.SIGINT
	Kill      Signal = syscall.SIGKILL
)
在所有系统中都能够使用的是interrupe,给进程发送一个信号，强制杀死该进程


type SyscallError　//SyscallError记录了一个特定系统调用的错误，主要用来返回SyscallError的字符串表示
type SyscallError struct {
	Syscall string
	Err     error
}

func (e *SyscallError) Error() string　　　//返回SyscallError的字符串表示
//

///go 黑魔法


今天我要教大家一些无用技能，也可以叫它奇技淫巧或者黑魔法。用得好可以提升性能，用得不好就会招来恶魔，嘿嘿。

黑魔法导论

为了让大家在学习了基础黑魔法之后能有所悟，在必要的时候能创造出本文传授之外的属于自己的魔法，这里需要先给大家打好基础。

学习Go语言黑魔法之前，需要先看清Go世界的本质，你才能获得像Neo一样的能力。

在Go语言中，Slice本质是什么呢？是一个reflect.SliceHeader结构体和这个结构体中Data字段所指向的内存。String本质是什么呢？是一个reflect.StringHeader结构体和这个结构体所指向的内存。

在Go语言中，指针的本质是什么呢？是unsafe.Pointer和uintptr。

当你清楚了它们的本质之后，你就可以随意的玩弄它们，嘿嘿嘿。

第一式 - 获得Slice和String的内存数据

让我小试身手，你有一个CGO接口要调用，需要你把一个字符串数据或者字节数组数据从Go这边传递到C那边，比如像这个：mysql/conn.go at master · funny/mysql · GitHub

查了各种教程和文档，它们都告诉你要用C.GoString或C.GoBytes来转换数据。

但是，当你调用这两个函数的时候，发生了什么事情呢？这时候Go复制了一份数据，然后再把新数据的地址传给C，因为Go不想冒任何风险。

你的C程序只是想一次性的用一下这些数据，也不得不做一次数据复制，这对于一个性能癖来说是多麽可怕的一个事实！

这时候我们就需要一个黑魔法，来做到不拷贝数据又能把指针地址传递给C。

// returns &s[0], which is not allowed in go
func stringPointer(s string) unsafe.Pointer {
	p := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return unsafe.Pointer(p.Data)
}

// returns &b[0], which is not allowed in go
func bytePointer(b []byte) unsafe.Pointer {
	p := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	return unsafe.Pointer(p.Data)
}
以上就是黑魔法第一式，我们先去到Go字符串的指针，它本质上是一个*reflect.StringHeader，但是Go告诉我们这是一个*string，我们告诉Go它同时也是一个unsafe.Pointer，Go说好吧它是，于是你得到了unsafe.Pointer，接着你就躲过了Go的监视，偷偷的把unsafe.Pointer转成了*reflect.StringHeader。

有了*reflect.StringHeader，你很快就取到了Data字段指向的内存地址，它就是Go保护着不想给你看到的隐秘所在，你把这个地址偷偷告诉给了C，于是C就愉快的偷看了Go的隐私。

第二式 - 把[]byte转成string

你肯定要笑，要把[]byte转成string还不简单？Go语言初学者都会的类型转换语法：string(b)。

但是你知道这么做的代价吗？既然我们能随意的玩弄SliceHeader和StringHeader，为什么我们不能造个string给Go呢？Go的内部会不会就是这么做的呢？

先上个实验吧：

package labs28

import "testing"
import "unsafe"

func Test_ByteString(t *testing.T) {
	var x = []byte("Hello World!")
	var y = *(*string)(unsafe.Pointer(&x))
	var z = string(x)

	if y != z {
		t.Fail()
	}
}

func Benchmark_Normal(b *testing.B) {
	var x = []byte("Hello World!")
	for i := 0; i < b.N; i ++ {
		_ = string(x)
	}
}

func Benchmark_ByteString(b *testing.B) {
	var x = []byte("Hello World!")
	for i := 0; i < b.N; i ++ {
		_ = *(*string)(unsafe.Pointer(&x))
	}
}
这个实验先证明了我们可以用[]byte的数据造个string给Go。接着做了两组Benchmark，分别测试了普通的类型转换和伪造string的效率。

结果如下：

$ go test -bench="."
PASS
Benchmark_Normal    20000000            63.4 ns/op
Benchmark_ByteString    2000000000           0.55 ns/op
ok      github.com/idada/go-labs/labs28 2.486s
哟西，显然Go这次又为了稳定性做了些复制数据之类的事情了！这让性能癖怎么能忍受！

我现在手头有个[]byte，但是我想用strconv.Atoi()把它转成字面含义对应的整数值，竟然需要发生一次数据拷贝把它转成string，比如像这样：mysql/types.go at master · funny/mysql · GitHub，这实在不能忍啊！

出招：

// convert b to string without copy
func byteString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
我们取到[]byte的指针，这次Go又告诉你它是*byte不是*string，你告诉它滚犊子这是unsafe.Pointer，Go这下又老实了，接着你很自在的把*byte转成了*string，因为你知道reflect.StringHeader和reflect.SliceHeader的结构体只相差末尾一个字段，两者的内存是对齐的，没必要再取Data字段了，直接转吧。

于是，世界终于安宁了，嘿嘿。

第三式 - 结构体和[]byte互转

有一天，你想把一个简单的结构体转成二进制数据保存起来，这时候你想到了encoding/gob和encoding/json，做了一下性能测试，你想到效率有没有可能更高点？

于是你又试了encoding/binady，性能也还可以，但是你还不满意。但是瓶颈在哪里呢？你恍然大悟，最高效的办法就是完全不解析数据也不产生数据啊！

怎么做？是时候使用这个黑魔法了：

type MyStruct struct {
	A int
	B int
}

var sizeOfMyStruct = int(unsafe.Sizeof(MyStruct{}))

func MyStructToBytes(s *MyStruct) []byte {
	var x reflect.SliceHeader
	x.Len = sizeOfMyStruct
	x.Cap = sizeOfMyStruct
	x.Data = uintptr(unsafe.Pointer(s))
	return *(*[]byte)(unsafe.Pointer(&x))
}

func BytesToMyStruct(b []byte) *MyStruct {
	return (*MyStruct)(unsafe.Pointer(
		(*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
	))
}
这是个曲折但又熟悉的故事。你造了一个SliceHeader，想把它的Data字段指向你的结构体，但是Go又告诉你不可以，你像往常那样把Go提到一边，你得到了unsafe.Pointer，但是这次Go有不死心，它告诉你Data是uintptr，unsafe.Pointer不是uintptr，你大脚把它踢开，怒吼道：unsafe.Pointer就是uintptr，你少拿这些概念糊弄我，Go屁颠屁颠的跑开了，现在你一马平川的来到了函数的出口，Go竟然已经在哪里等着你了！你上前三下五除二把它踢得远远的，顺利的把手头的SliceHeader转成了[]byte。

过了一阵子，你拿到了一个[]byte，你知道需要把它转成MyStruct来读取其中的数据。Go这时候已经完全不是你的对手了，它已经洗好屁股在函数入口等你，你一行代码就解决了它。

第四式 - 用CGO优化GC

你已经是Go世界的Neo，Go跟本没办法拿你怎么样。但是有一天Go的GC突然抽风了，原来这货是不管对象怎么用的，每次GC都给来一遍人口普查，导致系统暂停时间很长。

可是你是个性能癖，你把一堆数据都放在内存里方便快速访问，你这时候很想再踢Go的屁股，但是你没办法，毕竟你还在Go的世界里，你现在得替它擦屁股了，你似乎看到Go躲在一旁偷笑。

你想到你手头有CGO，可以轻易的用C申请到Go世界外的内存，Go的GC不会扫描这部分内存。

你还想到你可以用unsafe.Pointer将C的指针转成Go的结构体指针。于是一大批常驻内存对象被你用这种方式转成了Go世界的黑户，Go的GC一下子轻松了下来。

但是你手头还有很多Slice，于是你就利用C申请内存给SliceHeader来构造自己的Slice，于是你旗下的Slice纷纷转成了Go世界的黑户，Go的GC终于平静了。

但好景总是不长久，有一天Go世界突然崩溃了，只留下一句话：Segmentation Fault。你一下怂了，怎么段错误了？

经过一个通宵排查，你发现你管辖的黑户对象竟然偷偷的跟Go世界的其它合法居民搞在一起，当Go世界以为某个居民已经消亡时，用GC回收了它的住所，但是你的地下世界却认为它还活着，还继续访问它