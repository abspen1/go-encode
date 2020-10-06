package main

import (
	"fmt"
)

func main() {
	str := "Mississippi"
	original := str
	n := len(str)
	str2 := make([]string, n)

	for i := 0; i < n; i++ {
		a := []rune(str)
		str = string(a[1:]) + string(a[0])
		str2[i] = str
	}

	fmt.Println("\n--- Unsorted --- \n\n", str2)
	insertionSort(&str2, n) // Pass the address of our string array

	var originalLocation int
	last := make([]string, n)

	for i := 0; i < n; i++ {
		current := str2[i]

		last[i] = string(current[(n - 1)])

		if str2[i] == original {
			originalLocation = i
		}
	}
	fmt.Println(originalLocation)
	printEncodedLine(last, n)
	fmt.Println("\n--- Sorted ---\n\n", str2)
}

func printEncodedLine(last []string, n int) {
	num := 1

	for i := 0; i < n; i++ {
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

func insertionSort(items *[]string, n int) {
	var x int

	for i := 1; i < n; i++ {
		key := string((*items)[i])
		x = i - 1
		for x >= 0 && string((*items)[x]) > key {
			(*items)[x+1] = (*items)[x]
			x = x - 1
		}
		(*items)[x+1] = key
	}
}
