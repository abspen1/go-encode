package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var str string
	lines := 0
	for scanner.Scan() {
		str = scanner.Text()
		if len(str) > 0 {
			if lines != 0 {
				fmt.Println()
			} else {
				lines++
			}
			original := str
			n := len(str)
			str2 := make([]string, n)

			for i := 0; i < n; i++ {
				a := []rune(str)
				str = string(a[1:]) + string(a[0])
				str2[i] = str
			}

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
		} else {
			fmt.Println()
		}
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

}

func printEncodedLine(last []string, n int) {
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
