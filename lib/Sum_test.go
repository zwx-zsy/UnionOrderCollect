package lib

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	var S = 40.00
	var Sum float64
	i := 1
	Sum = S
	for i < 1000 {
		S = S * 0.5
		Sum = Sum + S
		i++
		fmt.Println(Sum)
	}

	fmt.Println("Sum:", Sum)
}
