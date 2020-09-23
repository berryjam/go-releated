package main
//
//import (
//	"fmt"
//	"github.com/youmark/pkcs8"
//	"hash/crc32"
//	"io/ioutil"
//	"os"
//	"reflect"
//	"runtime/debug"
//)
//
//type User struct {
//	Id   int
//	Name string
//	Age  int
//}
//
//type Manager struct{}
//
//func (m Manager) IntSwap(a, b int) (int, int) {
//	return a, b
//}
//
//func (m Manager) FloatSwap(a, b float64) (float64, float64) {
//	return a, b
//}
//
//func (u User) ReflectCallFuncHasArgs(name string, age int) {
//	fmt.Println("ReflectCallFuncHasArgs name: ", name, ", age:", age, "and origal User.Name:", u.Name)
//}
//
//func (u User) ReflectCallFuncNoArgs() {
//	fmt.Println("ReflectCallFuncNoArgs")
//}
//
//type Tmp struct {
//	Tags []interface{}
//}
//
//func main() {
//	//fmt.Printf("%v\n", 281475003522658%16)
//	//ReflectUsage()
//	//var tmp Tmp
//	//tmp.Tags =nil
//	//tmp.Tags = append(tmp.Tags,1)
//	//tmp.Tags = append(tmp.Tags,2)
//	//fmt.Println(len(tmp.Tags))
//	//a()
//	//fmt.Println("normally returned from main")
//	key := `-----BEGIN ENCRYPTED PRIVATE KEY-----
//MIHFMCgGCiqGSIb3DQEMAQMwGgQUjgUAk/I02VEqIlIM92zFpJuJJSECAggABIGY
//xMBRK+m6nEM85XrviGNO3bn051ALM7nVPq81f8NyBqbwtY5y+U3hHv9iiOUM+BWA
//d6EbIWHi6M0STRsJ/UGx+/UvCcr2szZhP20ynzVQMlrPkKbmfL4STrBuKCZoiWzk
//I+WkO6vPoFP6FM/0ztd/u97rA9yapWAwEjYRkO4EcKqe2QHQtQ6Dd+2bU4PvVPnF
//dgtQCiDHw2c=
//-----END ENCRYPTED PRIVATE KEY-----`
//	pwd := `Pwd12345^`
//	_,err := pkcs8.ParsePKCS8PrivateKeyECDSA([]byte(key),[]byte(pwd))
//	if err != nil {
//		panic(err)
//	}
//}
//
//func r() {
//	if r := recover(); r != nil {
//		fmt.Println("Recovered", r)
//		debug.PrintStack()
//	}
//}
//
//func a() {
//	defer r()
//	n := []int{5, 7, 4}
//	fmt.Println(n[3])
//	fmt.Println("normally returned from a")
//}
//
//func ReflectUsage() {
//	swap := func(in []reflect.Value) []reflect.Value {
//		return []reflect.Value{in[0].Index(1), in[0].Index(0)}
//	}
//
//	makeSwap := func(fptr interface{}) {
//		fn := reflect.ValueOf(fptr).Elem()
//		v := reflect.MakeFunc(fn.Type(), swap)
//		fn.Set(v)
//	}
//
//	var intSwap func(...int) (int, int)
//	makeSwap(&intSwap)
//	fmt.Println(intSwap(0, 1))
//}
//
//func printSqlTemplate(file string, shardCount int) {
//	reader, _ := os.Open(file)
//	bytes, _ := ioutil.ReadAll(reader)
//	for i := 0; i < shardCount; i++ {
//		fmt.Printf(string(bytes)+"\n\n", i)
//	}
//}
//
//func uid2shard(uid int64, shard uint32) {
//	res := crc32.ChecksumIEEE([]byte(fmt.Sprintf("%v", uid))) % shard
//	fmt.Printf("%v\n", res)
//}
//
//func orderId2shard(orderId int64, shard uint32) {
//	res := orderId & int64(shard-1)
//	fmt.Printf("%v\n", res)
//}
