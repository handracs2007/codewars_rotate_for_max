// https://www.codewars.com/kata/56a4872cbb65f3a610000026/train/go

package main

import (
	"fmt"
	"strconv"
)

func rotateLeft(s string) string {
	runes := []rune(s)
	return string(runes[1:]) + string(runes[0])
}

func MaxRot(n int64) int64 {
	// Assume the current number as the biggest number
	max := n

	// Rotate the first time
	strN := rotateLeft(strconv.Itoa(int(n)))

	for i := 0; i < len(strN)-1; i++ {
		n2, _ := strconv.Atoi(strN)

		// Check the bigger number
		if int64(n2) > max {
			max = int64(n2)
		}

		// Do another rotation
		left := strN[0 : i+1]
		right := strN[i+1:]
		strN = left + rotateLeft(right)
	}

	return max
}

func main() {
	fmt.Println(MaxRot(56789))
	//fmt.Println(rotateLeft("Hello"))
}
