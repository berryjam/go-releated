package main

import (
	"math/rand"
	"time"
	"fmt"
	"github.com/lafikl/consistent"
	"strconv"
)

var PHONE_PREFIX = []int{130, 131, 132, 133, 134, 135, 136, 137, 138, 139,
	150, 151, 152, 153, 155, 156, 157, 158, 159,
	176, 177, 178,
	180, 181, 182, 183, 184, 185, 186, 187, 188, 189}

func phoneNumRandGen(count int) []int {
	rand.Seed(time.Now().Unix())
	res := make([]int, 0)
	for _, phonePrefix := range PHONE_PREFIX {
		for i := 0; i < count/len(PHONE_PREFIX); i++ {
			res = append(res, phonePrefix*10000*10000+rand.Intn(10000)*10000+rand.Intn(10000))
		}
	}
	return res
}

func main() {
	counts := []int{5000, 50000, 500000, 5000000}
	for _, count := range counts {
		phones := phoneNumRandGen(count)
		modDelNeedChangeNum := 0
		modIncrNeedChangeNum := 0
		csDelNeedChangeNum := 0
		csIncrNeedChangeNum := 0
		for _, phone := range phones {
			if phone%1024 != phone%1023 {
				modDelNeedChangeNum++
			}
			if phone%1024 != phone%1025 {
				modIncrNeedChangeNum++
			}
		}
		//fmt.Printf("modNeedChangeNum:=%+v\n", modNeedChangeNum)

		c := consistent.New()
		for i := 0; i < 1024; i++ {
			c.Add(strconv.Itoa(i))
		}

		oriMap := make(map[int]string)
		delMap := make(map[int]string)
		incrMap := make(map[int]string)
		for _, phone := range phones {
			oriHost, err := c.Get(strconv.Itoa(phone))
			if err != nil {
				panic(err)
			}
			oriMap[phone] = oriHost
		}
		c.Remove("1023")
		for _, phone := range phones {
			delHost, err := c.Get(strconv.Itoa(phone))
			if err != nil {
				panic(err)
			}
			delMap[phone] = delHost
		}
		c.Add("1023")
		c.Add("1024")
		for _, phone := range phones {
			incrHost, err := c.Get(strconv.Itoa(phone))
			if err != nil {
				panic(err)
			}
			incrMap[phone] = incrHost
		}

		for _, phone := range phones {
			if oriMap[phone] != delMap[phone] {
				csDelNeedChangeNum++
			}
			if oriMap[phone] != incrMap[phone] {
				csIncrNeedChangeNum++
			}
		}

		fmt.Printf("当phone数量为%d:\n modDelNeedChangeNum=%+v modIncrNeedChangeNum=%+v csDelNeedChangeNum=%+v csIncrNeedChangeNum=%+v\n\n", count, modDelNeedChangeNum, modIncrNeedChangeNum, csDelNeedChangeNum, csIncrNeedChangeNum)
	}

}
