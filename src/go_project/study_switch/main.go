package main

import "fmt"

// switch
// 简化大量的判断

func main() {

	switch n := 3; n {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("大拇指")
	case 3:
		fmt.Println("大拇指")
	default:
		fmt.Println("无效的数值")

	}
}
