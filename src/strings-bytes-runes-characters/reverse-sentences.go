package main

import "fmt"

func reverseSentences(sentences string) {
	fmt.Printf("before:%v\n", sentences)
	runeArr := []rune(sentences)
	for i := 0; i < len(runeArr)/2; i++ {
		runeArr[i], runeArr[len(runeArr)-1-i] = runeArr[len(runeArr)-1-i], runeArr[i]
	}
	fmt.Printf("after:%v\n", string(runeArr))
}

func main() {
	reverseSentences("Hello World")
}
