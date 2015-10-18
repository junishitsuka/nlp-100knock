package main

import "fmt"
import "strings"
import "strconv"

func main() {
	statement := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."

	words := strings.Split(statement, " ")
	result := make([]int, len(words))
	// resultを空のスライスで宣言してappendすることもできるが
	// performanceが良くないのでサイズで初期化する

	for i, v := range words {
		fmt.Println(v + ":" + strconv.Itoa(len(v)))
		result[i] = len(v)
	}
	fmt.Println(result)
}
