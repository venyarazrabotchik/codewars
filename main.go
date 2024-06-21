package main

import (
	"fmt"

	"github.com/venyarazrabotchik/codewars/kata8"
)

func main() {
	fmt.Println(kata8.MergeArrays([]int{1, 2, 3, 4, 5}, []int{5, 6, 7, 8, 9}))
	fmt.Println(kata8.MergeArrays([]int{1, 2, 3, 4, 5}, []int{4, -9, 5, 8}))

}
