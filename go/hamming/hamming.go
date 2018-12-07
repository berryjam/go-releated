package hamming

import "fmt"

const testVersion = 5

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, fmt.Errorf("")
	}

	if len(a) == 0 && len(b) == 0 {
		return 0, nil
	}

	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}

func getDistance(distances [][] int, i, j int) int {
	if i == -1 || j == -1 {
		return 0
	} else {
		return distances[i][j]
	}
}
