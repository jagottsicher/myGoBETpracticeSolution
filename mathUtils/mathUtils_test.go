package mathUtils

import (
	"fmt"
	"testing"
)

func BenchmarkTheSumOf(b *testing.B) {

	sliceOfInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		TheSumOf(sliceOfInt...)
	}

}

func BenchmarkSumUpIntegers(b *testing.B) {

        sliceOfInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

        b.ResetTimer()

        for i := 0; i < b.N; i++ {
                SumUpIntegers(sliceOfInt...)
        }

}

func ExampleSumUpIntegers() {
        sliceOfInt := []int{1, 2, 3}
	fmt.Println("Sum:", SumUpIntegers(sliceOfInt...))
	// Output:
	// Sum: 6
}

func ExampleAddTwoIntegers() { 
        sliceOfInt := []int{1, 2}
        fmt.Println("Sum:", addTwoIntegers(sliceOfInt[0], sliceOfInt[1]))
        // Output:
        // Sum: 3
}

func TestSumUpIntegers(t *testing.T) {
        sliceOfInt := []int{1, 2, 3}
        wrongOutput := SumUpIntegers(sliceOfInt...)
	if wrongOutput != 6 {
		t.Error("Expected", 6, "got", wrongOutput)
	}
}

func TestAddTwoIntegers(t *testing.T) {
        sliceOfInt := []int{1, 2}
	wrongOutput := addTwoIntegers(sliceOfInt[0], sliceOfInt[1])
	if wrongOutput != 3 { 
                t.Error("Expected", 3, "got", wrongOutput)
        }
}

func TestSumUpIntegersTT(t *testing.T) {

	type testData struct {
		inputInts	[]int
	  	result		int
	}

	tests := []testData{
		testData{[]int{1, 2, 3, 4, 5, 6}, 21},
		testData{[]int{1, 2, 4, 8, 16}, 31},
		testData{[]int{1, -2, 3, -4, 5}, 3},
		testData{[]int{-5, 1, 1, 1, 1, 1}, 0},
		testData{[]int{6, 5, 4, 3, 2, 1}, 21},
	}

	for _, v := range tests {
        	wrongOutput := SumUpIntegers(v.inputInts...)
        	if wrongOutput != v.result {
                	t.Error("Expected", v.result, "got", wrongOutput)
        	}
	}
}

func TestTheSumOfTT(t *testing.T) {

        type testData struct {
                inputInts       []int
                result          int
        }

        tests := []testData{
                testData{[]int{1, 2, 3, 4, 5, 6}, 21},
                testData{[]int{1, 2, 4, 8, 16}, 31},
                testData{[]int{1, -2, 3, -4, 5}, 3},
                testData{[]int{-5, 1, 1, 1, 1, 1}, 0},
                testData{[]int{6, 5, 4, 3, 2, 1}, 21},
        }

        for _, v := range tests {
                wrongOutput := TheSumOf(v.inputInts...)
                if wrongOutput != v.result {
                        t.Error("Expected", v.result, "got", wrongOutput)
                }
        }
}











