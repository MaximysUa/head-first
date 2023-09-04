package kata6kyu

import (
	"fmt"
	"regexp"
)

func Arrange(s []int) []int {
	n := make([]int, 0)
	for k, j := 0, len(s)-1; k <= j; k, j = k+1, j-1 {
		if k == j && len(s)%2 == 1 {
			n = append(n, s[k])
		} else {
			if k%2 == 0 {
				n = append(n, s[k], s[j])
			} else {
				n = append(n, s[j], s[k])
			}
		}
	}
	return n
}
func FindMissingLetter(chars []rune) rune {
	j := chars[0]
	var result rune
	for i, k := range chars {
		fmt.Println(j, k)
		if k != j {
			return chars[i] - 1
		}
		j++
	}
	return result
}
func AscendDescend(length, minimum, maximum int) string {
	var res string
	counter := minimum
	var isUp bool = true
	if maximum < minimum || length == 0 {
		return ""
	}
	if maximum == minimum {
		for i := 0; i < length; i++ {
			res += fmt.Sprintf("%d", counter)
		}
		return res[0:length]
	}
	for i := 0; i < length; i++ {
		if counter <= maximum && isUp {
			res += fmt.Sprintf("%d", counter)
			if counter == maximum {
				isUp = false
				counter--
				continue
			} else {
				counter++
			}
		}

		if counter >= minimum && !isUp {
			res += fmt.Sprintf("%d", counter)
			if counter == minimum {
				isUp = true
				counter++
				continue
			} else {
				counter--
			}
		}

	}
	return res[0:length]
}
func FoldArray(arr []int, runs int) []int {
	var l int

	res := make([]int, len(arr))
	copy(res, arr)
	for runs != 0 {
		var isOdd bool
		l = len(res) / 2
		if len(res)%2 != 0 {
			isOdd = true
		}
		temp := make([]int, l)
		if isOdd {
			temp = append(temp, res[l])
		}
		for i := 0; i < l; i++ {
			temp[i] = res[i] + res[len(res)-1-i]
		}
		res = make([]int, len(temp))
		copy(res, temp)
		runs--
	}
	return res
}
func ConsonantsSubstringCounter(str string) int {
	compile, _ := regexp.Compile("[aeiou]")
	split := compile.Split(str, -1)
	var counter int
	for _, j := range split {
		var temp int
		for _, r := range j {

			temp += int(r - 'a' + 1)
		}
		if temp > counter {
			counter = temp
		}
	}
	return counter
}
func MinDistance(n int) int {
	var list []int
	for i := 1; i < n; i++ {
		if n%i == 0 {
			list = append(list, i)
		}
	}
	list = append(list, n)
	minNum := n
	for i := 1; i < len(list); i++ {
		temp := list[i] - list[i-1]
		if temp < minNum {
			minNum = temp
		}
	}
	return minNum
}
