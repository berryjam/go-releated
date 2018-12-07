package main

import (
	"regexp"
	"os"
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
)

/**
he regexp package has 3 functions to match:
if it matches a pattern, then it returns true, returning false otherwise.
 */
func IsIP(ip string) (b bool) {
	if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
		return false
	}

	return true
}

func IsNumber() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: regexp [string]")
		os.Exit(1)
	} else if m, _ := regexp.MatchString("^[0-9]+$", os.Args[1]); m {
		fmt.Println("Number")
	} else {
		fmt.Println("Not number")
	}
}

/**
Match mode can verify content but it cannot cut, filter or collect data from it. If you want to do that,
you have to use the complex mode of Regexp.
 */
func Filter() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("http get error.")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error")
		return
	}

	src := string(body)

	// Convert HTML tags to lower case.
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	// Remove STYLE.
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	// Remove SCRIPT.
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	// Remove all HTML code in angle brackets, and replace with newline.
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	// Remove continuous newline.
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	fmt.Println(strings.TrimSpace(src))
}

func Find() {
	a := "I am learning Go language"

	re, _ := regexp.Compile("[a-z]{2,4}")

	// Find the first match.
	one := re.Find([]byte(a))
	fmt.Println("Find:", string(one))

	// Find all matches and save to a slice, nless than 0 means return all matches,indicated length of slice
	// if it's greater than 0.
	all := re.FindAll([]byte(a), -1)
	for _, v := range all {
		fmt.Println(string(v))
	}
	fmt.Println("FindAll", all)

	// Find index of first match,start and end position. index是左闭右开的
	index := re.FindIndex([]byte(a))
	fmt.Println("FindIndex", index)

	// Find index of all matches,the n does same job as above.
	allindex := re.FindAllIndex([]byte(a), -1)
	fmt.Println("FindAllIndex", allindex)

	re2, _ := regexp.Compile("am(.*)lang(.*)")

	// Find first submatch and return array, the first element contains all elements, the second element contains the result of first (), the third element contains the result of second ().
	// Output:
	// the first element: "am learning Go language"
	// the second element: " learning Go ", notice spaces will be outputed as well.
	// the third element: "uage"
	submatch := re2.FindSubmatch([]byte(a))
	fmt.Println("FindSubmatch", submatch)
	for _, v := range submatch {
		fmt.Println(string(v))
	}

	// Same as FindIndex().
	submatchindex := re2.FindSubmatchIndex([]byte(a))
	fmt.Println(submatchindex)

	// FindAllSubmatch, find all submatches.
	submatch1 := re2.FindAllSubmatch([]byte(a), -1)
	fmt.Println(submatch1)

	// FindAllSubmatchIndex,find index of all submatches.
	submatch1index := re2.FindAllSubmatchIndex([]byte(a), -1)
	fmt.Println(submatch1index)
}

func Expand() {
	src := []byte(`
	call hello alice
	hello bob
	call hello eve
	`)
	pat := regexp.MustCompile(`(?m)(call)\s+(?P<cmd>\w+)\s+(?P<arg>.+)\s*$`)
	res := []byte{}
	for _, s := range pat.FindAllSubmatchIndex(src, -1) {
		res = pat.Expand(res, []byte("$cmd('$arg')\n"), src, s)
	}
	fmt.Println(string(res))
}

func main() {
	//ip := "127.0.0.1"
	//fmt.Printf("%v is ip?%v\n", ip, IsIP(ip))
	//IsNumber()
	//Filter()
	//Find()
	Expand()
}
