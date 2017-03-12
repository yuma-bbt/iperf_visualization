package main

import (
	"fmt"
	"reflect"
	"strings"
	)

func main(){

	var hoge string
	hoge = "hoge,foo"
	var time []string
	var bandwidth []string
	var test  []string

	test = strings.Split(hoge,",")


	fmt.Println(reflect.TypeOf(hoge))
	fmt.Println(reflect.TypeOf(time))
	fmt.Println(reflect.TypeOf(bandwidth))
	fmt.Println(reflect.TypeOf(test))
	fmt.Println(reflect.TypeOf(test[0]))
	fmt.Println(reflect.TypeOf(test[1]))
	fmt.Println("appended")
	bandwidth = append(bandwidth,test[0])
	time = append(time,test[1])
	fmt.Println(time)
	fmt.Println(bandwidth)
}

