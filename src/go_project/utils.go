package go_project

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

package _021最新最系统区块链详解

import (
"bytes"
"encoding/binary"
"fmt"
"os"
)

func IntToByte(num int64) []byte{
	var buffer bytes.Buffer
	//func Write(w io.Writer, order ByteOrder, data interface{}) error
	err := binary.Write(&buffer, binary.BigEndian, num)
	CheckErr(pos: "IntToByte", err)
	return buffer.Bytes()
}

func CheckErr(pos string, err error){
	if err != nil{
		fmt.Println(a: "error,pos:",pos,err)
		os.Exit(code: 1)
	}
}

