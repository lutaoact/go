package main

import (
	"fmt"
)

func IsBigEndian() bool {
	var i int32 = 0x12345678
	var b byte = byte(i)
	if b == 0x12 {
		return true
	}

	return false
}

func mainIsBigEndian() {
	if IsBigEndian() {
		fmt.Println("大端序")
	} else {
		fmt.Println("小端序")
	}
}

func main() {
	fmt.Printf("[]byte(true) = %+v\n", []byte(true))
}
