package main

func main() {

	s := "ababcbacadefegdehijhklij"
	println(partitionLabels(s))
}

func partitionLabels(s string) []int {
	mk := make(map[string]int)
	for i, s := range s {
		mk[string(s)] = i
	}
	var m int
	var arr []int
	var res int
	for i, char := range s {
		m = maxNum(mk[string(char)], m)
		if i == m {
			arr = append(arr, m+1-res)
			res = m + 1
		}
	}
	return arr
}
func maxNum(i, j int) int {
	if i > j {
		return i
	}
	return j
}
