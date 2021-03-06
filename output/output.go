package output

import (
	"fmt"
	"strconv"
)

// PrintEncodedLine prints out our encoded message (trucated)
func PrintEncodedLine(last []string, n int) string {
	var str string
	num := 1

	for i := 0; i < n; i++ {
		// The first conditional is to make sure we dont go out of index range for last[i+1]
		if i+1 != n && last[i] == last[i+1] {
			num++
			continue
		}

		fmt.Print(num, " ", last[i])
		str += strconv.Itoa(num) + " " + last[i]
		if i != (n - 1) {
			fmt.Print(" ")
			str += " "
		}

		num = 1
	}
	return str
}

// PrintDecodedLine prints out our decoded message
func PrintDecodedLine(first []string, next []int, sum, index int) {
	var decodedString []string

	decodedString[0] = string(first[index])
	fmt.Print(decodedString[0])

	for i := 1; i < sum; i++ {
		index = next[index]
		decodedString[i] = string(first[index])
		fmt.Print(decodedString[i])
	}
}
