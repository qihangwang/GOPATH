package main

import "fmt"

func main() {
	// "Hello" ->  'H' 'e' 'l' 'l' 'o'
	s2 := "白萝卜"
	s3 := []rune(s2)
	s3[0] = '红'
	fmt.Println(string(s3))
}
