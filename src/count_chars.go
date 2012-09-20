package main

import (
	"fmt"
	//	"unicode/utf8"
	"reflect"
)

func main() {
	str1 := "HI hello world"
	str2 := "HI helloこん world"
	fmt.Println(str1, len(str1))
	fmt.Println(str2, len(str2))
	for i, b := range str2 { //range可處理UTF-8
		fmt.Printf("%c - pos:%d\n", b, i)
		switch t := reflect.ValueOf(b); t.Interface().(type) {
		case rune:
			fmt.Print("...")
		case byte:
			fmt.Print("###")
		default:
			fmt.Print("***")
		}
		fmt.Println(reflect.TypeOf(b))
	}
	fmt.Println("TypeOf str1:", reflect.TypeOf(str1))
}
