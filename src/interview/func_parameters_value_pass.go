package main

import (
	"fmt"
	"strconv"
)

func modify(ip *int) {
	fmt.Printf("函数里接收到的指针的内存地址是：%p\n", &ip)
	*ip = 1
}

func TestModify() {
	i := 10
	ip := &i
	fmt.Printf("原始指针的内存地址是：%p\n", &i)
	modify(ip)
	fmt.Println("int值被修改了，新值为：", i)
}

func modifyMap(p map[string]int) {
	fmt.Printf("函数里接收到map的内存地址是：%p\n", &p)
	p["张三"] = 20
}

func TestModifyMap() {
	persons := make(map[string]int)
	persons["张三"] = 19

	mp := &persons

	fmt.Printf("原始map的内存地址是：%p\n", mp)
	modifyMap(persons)
	fmt.Println("map值被修改了，新值为：", persons)
}

type Person struct {
	Name string
}

func modifyStruct(p Person) {
	fmt.Printf("函数里接收到Person的内存地址是：%p\n", &p)
	p.Name = "李四"
}

func TestModifyStruct() {
	p := Person{"张三"}
	fmt.Printf("原始Person的内存地址是：%p\n", &p)
	modifyStruct(p)
	fmt.Println(p)
}

func modifyStructPointer(p *Person) {
	fmt.Printf("函数里接收到Person的内存地址是：%p\n", &p)
	p.Name = "李四"
}

func TestModifyStructPointer() {
	p := &Person{"张三"}
	fmt.Printf("原始Person的内存地址是：%p\n", &p)
	modifyStructPointer(p)
	fmt.Println(*p)
}

func modifySlice(ages []int) {
	fmt.Printf("函数里接收到slice的内存地址是：%p\n", ages)
	ages[0] = 1
}

func TestModifySlice() {
	ages := []int{5, 5, 5}
	fmt.Printf("原始slice的内存地址是：%p\n", ages)
	modifySlice(ages)
	fmt.Println(ages)
}

type PersonWithPointer struct {
	name string
	age  *int
}

func modifyStructWithPointer(p PersonWithPointer) {
	p.name = "李四"
	*p.age = 20
}

func (p PersonWithPointer) String() string {
	return "姓名为：" + p.name + ",年龄为：" + strconv.Itoa(*p.age)
}

func TestModifyStructWithPointer() {
	i := 19
	p := PersonWithPointer{name: "张三", age: &i}
	fmt.Println(p)
	modifyStructWithPointer(p)
	fmt.Println(p)
}

// https://mp.weixin.qq.com/s?__biz=MzAxMzc4Mzk1Mw==&mid=2649838219&idx=1&sn=fae21ba5449083cbf5e2598b3f9f0541&chksm=8398b475b4ef3d63de5b1f54b16e49a2253db0cc2651746062f79300307cd68e78c7c9e0df2f&mpshare=1&scene=1&srcid=0312wx5YZGbtozbZbuQ7QSb5&key=11485469ee3633975d44b372902fbaa87c33000ba8730703acd2a6aa5ca1ff95e9b8f3abba37caa822de3d50015b476d0ae97e9c549dc8e9ed52756d44465b6659747e76f528470a998eafcb3b5d35c6&ascene=0&uin=MjI4MTc4MTU4Mw%3D%3D&devicetype=iMac+MacBookPro11%2C1+OSX+OSX+10.12.6+build(16G29)&version=12010110&nettype=WIFI&lang=zh_CN&fontScale=100&pass_ticket=8BxPkXatcFZshfZOhKVXPorEOXs%2BbcsrX1l0%2BP8ABzDRiaLM9jyuZoAaPLwGUx6y
func main() {
	TestModify()
	TestModifyMap()
	TestModifyStruct()
	TestModifyStructPointer()
	TestModifySlice()
	TestModifyStructWithPointer()
}
