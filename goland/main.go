package main

import (
	"fmt"
	"strings"
)

type String *string

func main() {
	var rstr String
	var pstr string

	fmt.Println(rstr)
	fmt.Println(pstr)

	pstr = "go turkıye"
	rstr = &pstr

	*rstr = "Go turkıye kampı"

	//	fmt.Println(pstr)
	//	fmt.Println(rstr)

	//	replaceStr(pstr)

	fmt.Println(*rstr)
	fmt.Println(pstr)
}

//func replaceStr(s string) {
//	s = strings.Replace(s, "go", "Go", 1)
//}

func replaceStr(s String) {
	*s = strings.Replace(*s, "go", "Go", 1)
}
