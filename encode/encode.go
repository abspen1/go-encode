package encode

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/abspen1/go-encode/output"
	"github.com/abspen1/go-encode/sort"
)

// Run function will encode our string
func Run(arg string) {
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

			if arg == "insertion" {
				sort.InsertionSort(&str2, n) // Pass the address of our string array
			} else if arg == "quick" {
				log.Fatal("Working on quicksort")
			} else {
				log.Fatal("invalid command", arg)
			}

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
			output.PrintEncodedLine(last, n)
		} else {
			fmt.Println()
		}
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
