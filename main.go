package main

import (
	"fmt"
	"mathUtils"
)

func main() {

	inputValueSlice := []int{1, 2, 3, 4, 5, 6}

	fmt.Println("Sum", mathUtils.TheSumOf(inputValueSlice...))
	fmt.Println("Sum", mathUtils.SumUpIntegers(inputValueSlice...))
}
