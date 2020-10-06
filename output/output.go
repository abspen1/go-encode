package output

import "fmt"

// PrintEncodedLine prints out our encoded message
func PrintEncodedLine(last []string, n int) {
	num := 1

	for i := 0; i < n; i++ {
		// The first conditional is to make sure we dont go out of index range for last[i+1]
		if i+1 != n && last[i] == last[i+1] {
			num++
			continue
		}

		fmt.Print(num, " ", last[i])

		if i != (n - 1) {
			fmt.Print(" ")
		}

		num = 1
	}
}