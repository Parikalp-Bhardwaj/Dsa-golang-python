package main

import (
	"fmt"
	"strings"
)

func main() {
	// path := "/home/"
	path := "/a/./b/../../c/"

	fmt.Println(simplifyPath(path))
}

func simplifyPath(path string) string {
	p1 := strings.Split(path, "/")
	res := make([]string, 0)
	for _, p := range p1 {
		if p == "." || p == "" {
			continue
		}
		if p == ".." {
			if len(res) >= 1 {
				res = res[:len(res)-1]
			}
		} else {
			res = append(res, p)
		}

	}
	return "/" + strings.Join(res, "/")
}
