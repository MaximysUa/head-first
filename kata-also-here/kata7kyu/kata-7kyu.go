package kata7kyu

import (
	"math"
)

func BreakChocolate(n, m int) int {
	if n == 0 || m == 0 {
		return 0
	}
	if n <= 1 && m <= 1 {
		return 0
	}

	return (n * m) - 1
}

// How to improve my func
// func IsLeapYear(year int) bool {
// return year % 400 == 0 || year % 100 != 0 && year % 4 == 0
// }
func IsLeapYear(year int) bool {
	if year < 1 {
		return false
	}
	if year%400 == 0 {
		return true
	}
	if year%4 == 0 {
		if year%100 == 0 {
			return false
		}
		return true
	}
	return false
}
func FunWithBinary(n, b int) []int {
	var res []int
	if b == 0 || n == 0 {
		return res
	}
	for i := 1; i <= int(math.Pow(2.0, (float64(n)))-1); i++ {
		if b&i > 0 {
			res = append(res, i)
		}
	}
	return res
}
func PickGrains(grains <-chan string) (good int, bad int) {
	for i := range grains {
		if i == "good" {
			good++
		} else {
			bad++
		}
	}

	return
}
func Dominator(a []int) int {
	result := make(map[int]int)
	for _, i := range a {
		result[i]++
	}
	for i, j := range result {
		if j > len(a)/2 {
			return i
		}
	}
	return -1
}
