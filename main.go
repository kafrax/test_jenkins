package main

import (
	"fmt"
	"errors"
	"sync"
	"reflect"
)
type modelInfo struct {
	typ          reflect.Type
	name         string
	secondary    string
	expire       string
	fieldIndex   map[string][]int
	uniqueLength int
	indexLength  int
	defaultVal   string
	sync.RWMutex
}
func main() {
	var m=&modelInfo{}
	fmt.Println(reflect.TypeOf(m).Elem())
	fmt.Println(test())
}

func test() (ret int, err error) {
	defer func() {
		ret, ret1, err := test1()
		fmt.Printf("%p\n   %p\n   %p\n", &ret, &ret1, &err)
		return
	}()
	ret=123
	err=errors.New("some err")
	return
}

func test1() (ret, ret1 int, err error) {
	fmt.Printf("%p\n   %p\n   %p\n", &ret, &ret1, &err)
	return 0, 0, nil
}

