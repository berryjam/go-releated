package moretypes

import "fmt"

type StructPointersVertex struct {
	X int
	Y int
}

func TestStructPointers() {
	v := StructPointersVertex{1, 2}
	p := &v
	p.X = 1e9 // 或者 (*p).X = 2e9，但这种方式太笨重，go能直接通过.来访问struct的成员变量
	fmt.Println(v)
}
