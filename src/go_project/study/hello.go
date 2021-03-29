package main

//导入语句
import "fmt"

// 函数外只能放置标识符（变量、常量、函数以及类型）的声明
var name string
var age int
var isok bool

//声明常量
//定义了常量之后不能改变
//在程序运行期间不会改变的量
//const pi = 3.1415926
//批量声明常量时，如果某一行后没有赋值，默认就和上一行一致
const (
	n1 = 100
	n2
	n3
)

// iota在
const (
	a1 = iota //0
	a2        //1
	a3        //2
)

//插队
const (
	c1 = iota //0
	c2 = 100
	c3 = iota //2
	c4        //3
)

//批量声明变量
var (
	bi string
	dd int
)

// mian函数是go语言程序的入口
func main() {
	name = "wqh"
	age = 17
	isok = true
	//go语言中非全局声明变量必须使用，不使用就编译不过去
	fmt.Println("Hello World!")
	fmt.Printf("name:%s", name) //%s占位符，使用name这个变量的值去替代占位符
	fmt.Print(isok)             //不换行
	fmt.Println(age)            //换行
	//声明变量并同时赋值
	var d2 string = "whb"
	fmt.Println(d2)
	var d1 = "rf"
	fmt.Println(d1)
	//简短变量声明 只能在函数里面有用
	s3 := "ghaag"
	fmt.Println(s3)
	//匿名变量 用一个_代替
	//同一个作用域中不能重复声明同名的变量（{}）
}
