// mathUtils servs nice utils for adding up integers
package mathUtils

func TheSumOf(inputValues ...int) int {

	sum := 0

	for _, v := range inputValues {
		sum += v
	}

	return sum
}

// addTwoIntegers adds up two integers and returns the sum
func addTwoIntegers(a, b int) int {
	return a + b
}

// SumUpIntegers sums up integers and makes use of addTwoIntegers
func SumUpIntegers(inputValues ...int) int {
	var sum int
	sum = inputValues[0]
	for _, v := range inputValues[1:] {
		sum = addTwoIntegers(sum, v)
	}
	return sum
}
