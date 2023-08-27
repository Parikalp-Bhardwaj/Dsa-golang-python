package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	orders := [][]string{{"David", "3", "Ceviche"},
		{"Corina", "10", "Beef Burrito"},
		{"David", "3", "Fried Chicken"},
		{"Carla", "5", "Water"},
		{"Carla", "5", "Ceviche"},
		{"Rous", "3", "Ceviche"}}

	fmt.Println(displayTable(orders))
}

func displayTable(orders [][]string) [][]string {
	dish := make(map[string]bool)
	table_no := make(map[string]bool)
	mk := make(map[string]map[string]int)
	var food []string
	var table []int

	for i := 0; i < len(orders); i++ {
		if !dish[orders[i][2]] {
			food = append(food, orders[i][2])
		}

		if !table_no[orders[i][1]] {
			a, _ := strconv.Atoi(orders[i][1])
			table = append(table, a)
		}
		if mk[orders[i][1]] == nil {
			mk[orders[i][1]] = make(map[string]int)
		}
		dish[orders[i][2]] = true
		table_no[orders[i][1]] = true
		mk[orders[i][1]][orders[i][2]]++
	}

	sort.Strings(food)
	sort.Ints(table)

	var first_Arr = make([]string, 1)
	first_Arr[0] = "Table"
	first_Arr = append(first_Arr, food...)
	var res [][]string
	res = append(res, first_Arr)
	for _, t := range table {
		out := make([]string, 1)
		a := strconv.Itoa(t)
		out[0] = a
		for j := 1; j < len(first_Arr); j++ {
			out = append(out, strconv.Itoa(mk[a][first_Arr[j]]))
		}
		res = append(res, out)
		

	}
	return res
}
