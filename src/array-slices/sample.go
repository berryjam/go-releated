package main

import (
	"fmt"
	//"go-releated/src/basics/moretypes"
	"regexp"
	"io/ioutil"
)

func Arrays() {
	/**
	An array's size is fixed;
	its length is part of its type ([4]int and [5]int are distinct, incompatible types).
	 */
	var a [4]int
	a[0] = 1 // a[2] == 0,the zero value of the int type
	i := a[0]
	fmt.Printf("i = %d\n", i)

	// An array literal can be specified like so:
	b := [2]string{"Penn", "Teller"}
	fmt.Printf("b[0] = %s\n", b[0])

	// Or, you can have the compiler count the array elements for you:
	b = [...]string{"Penn", "Teller"}
}

func Slices() {
	/**
	The type specification for a slice is []T, where T is the type of the elements of the slice.
	Unlike an array type, a slice type has no specified length.
	A slice literal is declared just like an array literal, except you leave out the element count:
	 */
	letters := []string{"a", "b", "c", "d"}
	for idx, _ := range letters {
		fmt.Printf("letters[%d] = %s\n", idx, letters[idx])
	}

	/**
	 A slice can be created with the built-in function called make, which has the signature
	 func make([]T, len, cap) []T
	 where T stands for the element type of the slice to be created. The make function takes a type, a length,
	 and an optional capacity. When called, make allocates an array and returns a slice that refers to that array.
	  */
	var s []byte
	s = make([]byte, 5, 5)
	// s == []byte{0,0,0,0,0}
	for idx, _ := range s {
		fmt.Printf("s[%d] = %d\n", idx, s[idx])
	}

	/**
	When the capacity argument is omitted, it defaults to the specified length.
	Here's a more succinct version of the same code:
	 */
	s = make([]byte, 5)

	/**
	The length and capacity of a slice can be inspected using the built-in len and cap functions.
	 */
	fmt.Printf("The length of s is %d\n", len(s))
	fmt.Printf("The capacity of is %d\n", cap(s))

	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	// b[1:4] == []byte{'o', 'l', 'a'}, sharing the same storage as b
	c := b[1:4]
	for idx, val := range c {
		fmt.Printf("c[%d] = %q\n", idx, val)
	}

	x := [3]string{"Лайка", "Белка", "Стрелка"}
	xslice := x[:] // a slice referencing the storage of x
	for idx, val := range xslice {
		fmt.Printf("xslice[%d] = %s\n", idx, val)
	}
}

func SliceInternals() {
	/**
	Slicing does not copy the slice's data. It creates a new slice value that points to the original array.
	This makes slice operations as efficient as manipulating array indices.
	Therefore, modifying the elements (not the slice itself) of a re-slice modifies the elements of the original slice:
	 */
	d := []byte{'r', 'o', 'a', 'd'}
	e := d[2:]
	// e == []byte{'a', 'd'}
	e[1] = 'm'
	// e == []byte{'a', 'm'}
	// d == []byte{'r', 'o', 'a', 'm'}

	/**
	The length is the number of elements referred to by the slice.
	The capacity is the number of elements in the underlying array (beginning at the element referred to by the slice pointer).
	The distinction between length and capacity will be made clear as we walk through the next few examples.
	As we slice s, observe the changes in the slice data structure and their relation to the underlying array:
	 */
	s := make([]byte, 5)
	s = s[2:4] // Hence,the length of s is 2,the capacity of s is 3!

	// Earlier we sliced s to a length shorter than its capacity. We can grow s to its capacity by slicing it again:
	s = s[:cap(s)]
	/**
	A slice cannot be grown beyond its capacity. Attempting to do so will cause a runtime panic,
	just as when indexing outside the bounds of a slice or array. Similarly,
	slices cannot be re-sliced below zero to access earlier elements in the array.
	slice 不能往前引用返回之前数组的元素
	 */
	// s = s[:6] // panic
	//s = s[-1:] // panic
}

/**
To increase the capacity of a slice one must create a new,
larger slice and copy the contents of the original slice into it.
This technique is how dynamic array implementations from other languages work behind the scenes.
The next example doubles the capacity of s by making a new slice, t, copying the contents of s into t,
and then assigning the slice value t to s:
 */
func GrowingSlices() {
	s := make([]byte, 5)
	t := make([]byte, len(s), (cap(s)+1)*2) // +1 in case cap(s) == 0
	//for i := range s {
	//	t[i] = s[i]
	//}
	//s = t

	// Using copy, we can simplify the code snippet above:
	copy(t, s)
	s = t

	/**
	A common operation is to append data to the end of a slice.
	This function appends byte elements to a slice of bytes,
	growing the slice if necessary, and returns the updated slice value:
	 */
	p := []byte{2, 3, 5}
	p = AppendByte(p, 7, 11, 13)
	// p == []byte{2,3,5,7,11,13}
	for idx, val := range p {
		fmt.Printf("p[%d] = %d\n", idx, val)
	}

	/**
	But most programs don't need complete control,
	so Go provides a built-in append function that's good for most purposes; it has the signature
	func append(s []T, x ...T) []T
	 */
	a := make([]int, 1)
	// a == []int{0}
	a = append(a, 1, 2, 3) // The append function appends the elements x to the end of the slice s, and grows the slice if a greater capacity is needed.
	// a == []int{0, 1, 2, 3}

	/**
	To append one slice to another, use ... to expand the second argument to a list of arguments.
	将一个slice append到另外一个slice，第二个参数后面加上...即可
	 */
	c := []string{"John", "Paul"}
	d := []string{"George", "Ringo", "Pete"}
	c = append(c, d...) // equivalent to "append(c, d[0], d[1], d[2])"
	// c == []string{"John","Paul","George", "Ringo", "Pete"}

	/**
	Since the zero value of a slice (nil) acts like a zero-length slice,
	 you can declare a slice variable and then append to it in a loop:
	 */
}

// Filter returns a new slice holding only
// the elements of s that satisfy f()
func Filter(s []int, fn func(int) bool) []int {
	var p []int // == nil
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary. reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

func PossibleGotcha() {
	/**
	As mentioned earlier, re-slicing a slice doesn't make a copy of the underlying array.
	The full array will be kept in memory until it is no longer referenced.
	Occasionally this can cause the program to hold all the data in memory when only a small piece of it is needed.
	For example, this FindDigits function loads a file into memory and searches it for the first group of consecutive numeric digits,
	returning them as a new slice.
	 */
	FindDigits("./sample.go")
}

func FindDigits(filename string) []byte {
	var digitRegexp = regexp.MustCompile("[0-9]+")
	b, _ := ioutil.ReadFile(filename)
	return digitRegexp.Find(b) // 返回值是b的一个slice，所以整个文件内容b被返回值引用，如果返回值slice一直保持，那么b既没有存在必要，但也不能被垃圾回收掉，浪费内存
}

/**
To fix this problem one can copy the interesting data to a new slice before returning it:
可以使用以下函数来解决上面函数出现的问题
 */
func CopyDigits(filename string) []byte {
	var digitRegexp = regexp.MustCompile("[0-9]+")
	b, _ := ioutil.ReadFile(filename)
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	// A more concise version of this function could be constructed by using append.
	//c = append([]byte{}, b...)
	return c
}

func main() {
	Arrays()
	Slices()
	SliceInternals()
	GrowingSlices()
	//moretypes.TestSliceBounds()
}
