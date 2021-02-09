package main

import (
	"fmt"
)

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}

	for _, stu := range stus {
		fmt.Printf("stu:%+v &stu:%+v\n", stu, &stu)
		m[stu.name] = &stu
	}
	fmt.Printf("%+v\n", m)
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}
