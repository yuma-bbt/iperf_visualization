package main

import (
	"fmt"
	"strings"
)

func main(){
	var variable string
	variable = "hoge,foo"
	var bandwidth []string
	var time []string
	var test []string
	test = strings.Split(variable,",")
	time[0] =append(time[1],test[0])
	bandwidth[0] = append(bandwidth[0],test[1])
	fmt.Println(time[0])
	fmt.Println(bandwidth[0])
}
