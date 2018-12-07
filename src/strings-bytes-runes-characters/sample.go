package main

import (
	"fmt"
	"unicode/utf8"
)

func PrintingStrings() {
	// In Go, a string is in effect a read-only slice of bytes.
	var sample = []byte("\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")

	fmt.Println("Println:")
	fmt.Println(sample)

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		/**
		A shorter way to generate presentable output for a messy string is to use the %x (hexadecimal) format
		verb of fmt.Printf. It just dumps out the sequential bytes of the string as hexadecimal digits, two per byte.
		 */
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("\n")

	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sample)

	/**
	A nice trick is to use the "space" flag in that format, putting a space between the % and the x.
	Compare the format string used here to the one above,
	 */
	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample)

	/**
	There's more. The %q (quoted) verb will escape any non-printable byte sequences
	in a string so the output is unambiguous.
	 */
	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample)

	/**
	If we are unfamiliar or confused by strange values in the string, we can use the "plus" flag to the %q verb.
	This flag causes the output to escape not only non-printable sequences, but also any non-ASCII bytes,
	 all while interpreting UTF-8.
	 The result is that it exposes the Unicode values of properly formatted UTF-8 that represents non-ASCII data in the string:
	 */
	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample) // the Unicode value of the Swedish symbol shows up as a \u escape:

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%q ", sample[i])
	}
	fmt.Printf("\n")
}

func UTF8AndStringLiterals() {
	// 反引号不转义任何字符串，所以是一个"raw string"，双引号会包含转义字符
	const placeOfInterest = `⌘` // a "raw string", enclosed by back quotes, so it can contain only literal text.

	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", placeOfInterest) // an ASCII-only quoted string
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i]) // individual bytes in hexadecimal
	}
	fmt.Printf("\n")
}

/**
To summarize, here are the salient points:

1.Go source code is always UTF-8.
2.A string holds arbitrary bytes.
3.A string literal, absent byte-level escapes, always holds valid UTF-8 sequences.
4.Those sequences represent Unicode code points, called runes.
5.No guarantee is made in Go that characters in strings are normalized.
 */
func CodePointsCharactersAndRunes() {

}

func RangeLoops() {
	const nihongo = "日本語"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
}

func Libraries() {
	const nihongo = "日本語"
	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
}

func iterString() {
	/**
	1.golang的string是以UTF-8编码的,而UTF-8是一种1-4字节的可变长字符集，每个字符可用1-4字节来表示
	2.使用下标方式s[i]访问字符串s，s[i]是UTF-8编码后的一个字节(uint8)，即按字节遍历
	3.使用for i,v := range s 方式访问s，i是字符串下标编号，v是对应的字符值(int32=rune)，即按字符遍历
	4.使用fmt.Printf打印时，%c占位符打印的是字符
	5.如果希望以随机方式访问字符串s的每个字符，可以先转为[]rune数组，再以下标访问
	 */
	s := "我是中国人"

	fmt.Println(len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U 起始于字位置%d %x\n", s[i], i, s[i])
	}

	fmt.Printf("% x\n", s)

	fmt.Printf("\n")

	for index, runeValue := range s {
		fmt.Printf("%#U 起始于字位置%d %+v %x\n", runeValue, index, runeValue, runeValue)
	}

	fmt.Print("\n")

	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U 起始于字位置 %d\n", runeValue, i)
		w = width
	}
}

func testString() {
	s := "我是中国人"
	fmt.Printf("%c", s[0])
}

func testSample() {
	const sample = "\xbd\xb2\x3d\x20\xe2\x8c\x98"
	fmt.Println(sample)
}

func main() {
	//PrintingStrings()
	//UTF8AndStringLiterals()
	//CodePointsCharactersAndRunes()
	//RangeLoops()
	//Libraries()
	//iterString()
	//testString()
	testSample()
}
