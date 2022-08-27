package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	n := 218
	fmt.Println(reorderedPowerOf2(n))
}

func reorderedPowerOf2(n int) bool {
	r := reverse(strconv.Itoa(n))
	fmt.Println(r)
	for i := 0; i <= 30; i++ {
		res := 1 << i
		if r == reverse(strconv.Itoa(res)) {
			return true
		}
	}
	return false

}

func reverse(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
