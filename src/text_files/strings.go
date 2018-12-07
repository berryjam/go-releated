package main

import (
	"fmt"
	"strings"
	"strconv"
)

func Strings() {
	/**
	1.func Contains(s, substr string) bool
	Check if string s contains string substr, returns a boolean value.
	 */
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))

	/**
	2.func Join(a []string, sep string) string
	Combine strings from slice with separator sep.
	 */
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ","))

	/**
	3.func Index(s, sep string) int
	Find index of sep in string s, returns -1 if it's not found.
	 */
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicke", "dmr"))

	/**
	4.func Repeat(s string, count int) string
	Repeat string s count times.
	 */
	fmt.Println("ba" + strings.Repeat("na", 2))

	/**
	5.func Replace(s, old, new string, n int) string
	Replace string old with string new in string s. n is the number of replacements. If n is less than 0, replace all instances.
	 */
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))

	/**
	6.func Split(s, sep string) []string
	Split string s with separator sep into a slice.
	 */
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))

	/**
	7.func Trim(s string, cutset string) string
	Remove cutset of string s if it's leftmost or rightmost.
	 */
	fmt.Printf("[%q]\n", strings.Trim(" !!! Achtung !!! ", "! "))

	/**
	8.func Fields(s string) []string
	Remove space items and split string with space into a slice.
	 */
	fmt.Printf("Fields are: %q\n", strings.Fields(" foo bar baz "))
}

func Strconv() {
	/**
	1.Append series, convert data to string, and append to current byte slice.
	 */
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))

	/**
	2.Format series, convert other data types into string.
	 */
	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)

	/**
	3.Parse series, convert strings to other types.
	 */
	aBool, aErr := strconv.ParseBool("false")
	if aErr != nil {
		fmt.Println(aErr)
	}
	bFloat, bErr := strconv.ParseFloat("123.23", 64)
	if bErr != nil {
		fmt.Println(bErr)
	}
	cInt, cErr := strconv.ParseInt("1234", 10, 64)
	if cErr != nil {
		fmt.Println(cErr)
	}
	dUint, dErr := strconv.ParseUint("12345", 10, 64)
	if dErr != nil {
		fmt.Println(dErr)
	}
	eInt, eErr := strconv.Atoi("1023")
	if eErr != nil {
		fmt.Println(eErr)
	}
	fmt.Println(aBool, bFloat, cInt, dUint, eInt)
}

func main() {
	Strings()
	Strconv()
}
