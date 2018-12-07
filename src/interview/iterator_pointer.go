package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		// https://stackoverflow.com/questions/15945030/change-values-while-iterating-in-golang
		// stu只是stus的值的一个副本，并且stu的地址会重用		
		m[stu.Name] = &stu
	}

	x := make([]int, 3)
	x[0], x[1], x[2] = 1, 2, 3

	for i, val := range x {
		fmt.Println(&x[i], "vs.", &val)
	}

	for k, v := range m {
		fmt.Printf("Name:%s Age:%v\n", k, *v)
	}
}

func main() {
	pase_student()
}
