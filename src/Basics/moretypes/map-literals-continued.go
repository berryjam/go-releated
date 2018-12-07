package moretypes

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// Map literals are like struct literals, but the keys are required.
//var m = map[string]Vertex{
//	"Bell Labs": {40.68433, -74.39967},
//	"Google":    {37.42202, -122.08408},
//}

func TestMapLiteralsContinued() {
	fmt.Println(m)
}
