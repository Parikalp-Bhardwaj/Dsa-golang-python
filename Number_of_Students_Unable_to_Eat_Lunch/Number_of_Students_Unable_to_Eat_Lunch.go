package main

import "fmt"

func main() {
	// students := []int{1, 1, 0, 0}
	// sandwiches := []int{0, 1, 0, 1}
	students := []int{1, 1, 1, 0, 0, 1}
	sandwiches := []int{1, 0, 0, 0, 1, 1}
	fmt.Println(countStudents(students, sandwiches))
	
}

func countStudents(students []int, sandwiches []int) int {
	count := 0
	for len(students) > count {
		if students[0] == sandwiches[0] {
			sandwiches = append(sandwiches[:0], sandwiches[1:]...)
			count = 0
		} else {
			students = append(students, students[0])
			count++
		}
		students = append(students[:0], students[1:]...)
	}
	return len(students)
}
