package moretypes

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)

	tmp := strings.Fields(s)

	for _, value := range tmp {
		result[value]++
	}

	return result
}

func TestExerciseMaps() {
	wc.Test(WordCount)
}
