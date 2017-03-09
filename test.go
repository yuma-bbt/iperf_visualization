package main

import (
	"fmt"
)
func main(){

	var s []int
	for  i:=0;  i<2000; i++{
		s =append(s,i)
		fmt.Println(s[i])
	}
}

