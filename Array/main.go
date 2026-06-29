package main

import (
	"fmt"

	"Array/problems"
)

func main() {
	nums := []int{0, 1, 2, 0, 2, 1, 1, 0, 2}
	fmt.Println(problems.DutchFlagPartition(1, nums))
}
