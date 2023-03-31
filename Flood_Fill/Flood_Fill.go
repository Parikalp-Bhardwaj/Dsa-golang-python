package main

import "fmt"

func main() {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	sr := 1
	sc := 1
	color := 2
	fmt.Println(floodFill(image, sr, sc, color))
}

func floodFill(image [][]int, sr, sc, color int) [][]int {
	c := image[sr][sc]
	floodFillUitls(&image, sr, sc, c, color)
	return image
}

func floodFillUitls(image *[][]int, sr, sc, c, color int) {
	if sr < 0 || sr >= len(*image) {
		return
	}
	if sc < 0 || sc >= len((*image)[0]) {
		return
	}
	if (*image)[sr][sc] == color {
		return
	}
	if (*image)[sr][sc] != c {
		return
	}
	(*image)[sr][sc] = color

	floodFillUitls(image, sr+1, sc, c, color)
	floodFillUitls(image, sr-1, sc, c, color)
	floodFillUitls(image, sr, sc+1, c, color)
	floodFillUitls(image, sr, sc-1, c, color)

}
