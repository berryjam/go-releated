package moretypes

import "fmt"

// The zero value of a map is nil. A nil map has no keys, nor can keys be added.
var m map[string]Vertex

func TestMaps() {
	// m 需要make函数初始化后才能使用
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
