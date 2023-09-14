package kata6kyu

import (
	"fmt"
	"regexp"
	"slices"
	"sync"
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
func Race(v1, v2, g int) [3]int {
	if v1 >= v2 {
		return [3]int{-1, -1, -1}
	}

	t := int(float64(g) / float64(v2-v1) * 3600.0)

	h := t / 3600
	t = t % 3600

	m := t / 60
	s := t % 60

	return [3]int{h, m, s}
}

// TODO разобраться с этим
func MergeMultiple(c ...chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(c))
	for _, ch := range c {
		go func(ch <-chan string) {
			for v := range ch {
				out <- v
			}
			wg.Done()
		}(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
func Merge(a <-chan string, b <-chan string) <-chan string {
	c := make(chan string)
	done := make(chan string)

	go func() {
		// consume all messages from a
		for m := range a {
			c <- m
		}
		done <- ""
	}()

	go func() {
		// consume all messages from b
		for m := range b {
			c <- m
		}
		done <- ""
	}()

	go func() {
		// wait until a and b are both closed, and then close c
		<-done
		<-done
		close(c)
	}()

	return c
}

func CreateIterator(fn func(int) int, n int) func(int) int {
	return func(i int) int {
		count := n
		for count > 0 {
			i = fn(i)
			count--
		}
		return i
	}
}
func QueueTime(customers []int, n int) int {
	if len(customers) == 0 {
		return 0
	}
	checkout := make([]int, n)
	for _, i := range customers {
		slices.Sort(checkout)
		checkout[0] += i

	}
	slices.Sort(checkout)
	return checkout[n-1]
}
