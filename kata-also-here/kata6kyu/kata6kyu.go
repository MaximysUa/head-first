package kata6kyu

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
