package raindrops

import "strconv"

const testVersion = 2

func Convert(num int) string {
	res := ""
	plingAble, plangAble, plongAble := false, false, false
	for {
		if num%3 == 0 && !plingAble {
			res += "Pling"
			num /= 3
			plingAble = true
		} else if num%5 == 0 && !plangAble {
			res += "Plang"
			num /= 5
			plangAble = true
		} else if num%7 == 0 && !plongAble {
			res += "Plong"
			num /= 7
			plongAble = true
		} else {
			break
		}
	}

	if res == "" {
		return strconv.Itoa(num)
	}

	return res
}
