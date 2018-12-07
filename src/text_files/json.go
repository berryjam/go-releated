package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
)

type Server struct {
	// ID will not be outputed.
	ID         int `json:"-"`

	ServerName string `json:"serverName"`
	ServerIP   string `json:"serverIP"`
}

type Serverslice struct {
	Servers []Server `json:"servers"`
}

/**
Go只会反序列化已在struct定义的字段，所以即便json串很大，如果我们只关心某些字段并只定义某些字段时，Go就会忽略掉没定义的字段，方便使用

In the above example, we defined a corresponding structs in Go for our JSON,
using slice for an array of JSON objects and field name as our JSON keys.
But how does Go know which JSON object corresponds to which specific struct filed?
Suppose we have a key called Foo in JSON. How do we find its corresponding field?

1.First, Go tries to find the (capitalised) exported field whose tag contains Foo.

2.If no match can be found, look for the field whose name is Foo.

3.If there are still not matches look for something like FOO or FoO, ignoring case sensitivity.
 */
func ParseToStruct() {
	fmt.Println("=============ParseToStruct=============")
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.02"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s, "\n")
}

func ParseToInterface() {
	fmt.Println("=============ParseToInterface=============")
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	/**
	The f stores a map, where keys are strings and values are interface{}'s'.
	So, how do we access this data? Type assertion
	 */
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for idx, val := range vv {
				fmt.Println(idx, val)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
	fmt.Println("\n")
}

func SimpleJson() {
	fmt.Println("=============SimpleJson=============")
	js, err := simplejson.NewJson([]byte(`{
    "test": {
        "array": [1, "2", 3],
        "int": 10,
        "float": 5.150,
        "bignum": 9223372036854775807,
        "string": "simplejson",
        "bool": true
    }
}`))
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	arr, _ := js.Get("test").Get("array").Array()
	i, _ := js.Get("test").Get("int").Int()
	ms := js.Get("test").Get("string").MustString()

	fmt.Printf("array : %v,int : %v,string : %v\n", arr, i, ms)
	fmt.Println("\n")
}

/**
Here are some points you need to keep in mind when trying to produce JSON:

1.Field tags containing "-" will not be outputted.

2.If a tag contains a customized name, Go uses this instead of the field name, like serverName in the above example.

3.If a tag contains omitempty, this field will not be outputted if it is zero-value.

4.If the field type is bool, string, int, int64, etc, and its tag contains ",string",
Go converts this field to its corresponding JSON type.
 */
func ProducingJson() {
	fmt.Println("=============ProducingJson=============")
	var s Serverslice
	s.Servers = append(s.Servers, Server{ID: 3, ServerName:"Shanghai_VPN", ServerIP:"127.0.0.1"})
	s.Servers = append(s.Servers, Server{ID: 6, ServerName:"Beijing_VPN", ServerIP:"127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
		return
	}
	fmt.Println(string(b))
}

func main() {
	ParseToStruct()
	ParseToInterface()
	SimpleJson()
	ProducingJson()
}
