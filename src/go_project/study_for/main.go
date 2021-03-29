package main

import "fmt"

// Go语言函数外的语句必须关键字开头
//for循环

func main() {
	//函数内部定义的变量必须使用
	//基本格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//变种1
	i := 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}
	//变种2
	for i < 10 {
		fmt.Println(i)
		i++
	}
	//for range 循环
	s := "ass删除"
	for i, v := range s {
		fmt.Println(i, string(v))
		fmt.Printf("%d %c\n", i, v)
	}
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("over")

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("over")
}
