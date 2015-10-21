package main

import "fmt"

func main() {
	x := 12
	y := "気温"
	z := 22.4

	fmt.Println(fmt.Sprintf("%d時の%sは%.1f", x, y, z))
}
