package main

import (
	"fmt"
	"strings"
)

func t() (str string) {
	str = "asdasdasscxcxcxcxcxdfdeee"
	return str
}
func main() {
	str := "dsdd/dsdds/dsdsds/sdsd"
	i := strings.LastIndex(str, "/")
	str2 := str[i+1:]
	fmt.Println(str2)

	s := t()
	print(s)
}
