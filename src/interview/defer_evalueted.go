package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func deferAnnoy() int {
	var i int
	defer func() {
		i++
		fmt.Println("deferAnnoy defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("deferAnnoy defer1:", i)
	}()
	return i
}

func deferNamed() (i int) {
	defer func() {
		i++
		fmt.Println("deferNamed defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("deferNamed defer1:", i)
	}()

	return i
}

func deferVerify() *int {
	var i int
	defer func() {
		i++
		fmt.Println("deferVerify defer2:", i, &i)
	}()
	defer func() {
		i++
		fmt.Println("deferVerify defer1:", i, &i)
	}()
	return &i
}

func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func main() {
	//a := 1
	//b := 2
	//defer calc("1", a, calc("10", a, b))
	//a = 0
	//defer calc("2", a, calc("20", a, b))
	//b = 1

	//fmt.Println("deferAnnoy return:", deferAnnoy())
	//fmt.Println("deferNamed return:", deferNamed())
	//deferVerify := deferVerify()
	//fmt.Println("deferVerify return:", *deferVerify, deferVerify)

	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
}
