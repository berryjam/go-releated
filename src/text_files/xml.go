package main

import (
	"encoding/xml"
	"os"
	"fmt"
	"io/ioutil"
)

/**
字段后面的`xml:"servers"`称作struct field tag，用于序列化
 */
type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

/**
Here are some rules when using the xml package to parse XML documents to structs:
1.If the field type is a string or []byte with the tag ",innerxml", Unmarshal will assign raw XML data to it,
like Description in the above example:
Shanghai_VPN127.0.0.1Beijing_VPN127.0.0.2

2.If a field is called XMLName and its type is xml.Name, then it gets the element name, like servers in above example.

3.If a field's tag contains the corresponding element name, then it gets the element name as well,
like servername and serverip in the above example.

4.If a field's tag contains ",attr", then it gets the corresponding element's attribute, like version in above example.

5.If a field's tag contains something like "a>b>c", it gets the value of the element c of node b of node a.

6.If a field's tag contains "=", then it gets nothing.

7.If a field's tag contains ",any", then it gets all child elements which do not fit the other rules.

8.If the XML elements have one or more comments, all of these comments will be added to the first field that
has the tag that contains ",comments". This field type can be a string or []byte.
If this kind of field does not exist, all comments are discarded.

注意：所有字段都要首字母大写，因为需要导出
 */
func ParseXML() {
	file, err := os.Open("/Users/berryjam/Documents/go_workspace/src/go-releated/src/text_files/servers.xml") // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v)
}

/**
Here we can see that Marshal also receives a v parameter of type interface{}.
So what are the rules when marshalling to an XML document?

1.If v is an array or slice, it prints all elements like a value.

2.If v is a pointer, it prints the content that v is pointing to, printing nothing when v is nil.

3.If v is a interface, it deal with the interface as well.

4.If v is one of the other types, it prints the value of that type.

So how does xml.Marshal decide the elements' name? It follows the ensuing rules:

1.If v is a struct, it defines the name in the tag of XMLName.

2.The field name is XMLName and the type is xml.Name.

3.Field tag in struct.

4.Field name in struct.

5.Type name of marshal.

Then we need to figure out how to set tags in orderapi to produce the final XML document.

1.XMLName will not be printed.

2.Fields that have tags containing "-" will not be printed.

3.If a tag contains "name,attr", it uses name as the attribute name and the field value as the value,
like version in the above example.

4.If a tag contains ",attr", it uses the field's name as the attribute name and the field value as its value.

5.If a tag contains ",chardata", it prints character data instead of element.

6.If a tag contains ",innerxml", it prints the raw value.

7.If a tag contains ",comment", it prints it as a comment without escaping, so you cannot have "--" in its value.

8.If a tag contains "omitempty", it omits this field if its value is zero-value, including false, 0,
nil pointer or nil interface, zero length of array, slice, map and string.

9.If a tag contains "a>b>c", it prints three elements where a contains b and b contains c, like in the following code:

FirstName string xml:"name>first" LastName string xml:"name>last"
 */
func ProduceXML() {
	v := &Recurlyservers{Version: "1"}
	v.Svs = append(v.Svs, server{ServerName:"Shanghai_VPN", ServerIP:"127.0.0.1"})
	v.Svs = append(v.Svs, server{ServerName:"Beijing_VPN", ServerIP:"127.0.0.2"})
	output, err := xml.MarshalIndent(v, " ", "  ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write([]byte(xml.Header))

	os.Stdout.Write(output)
}

func main() {
	ParseXML()
	ProduceXML()
}
