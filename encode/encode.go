package encode

import (
	"fmt"
	"log"

	"github.com/abspen1/go-encode/output"
	"github.com/abspen1/go-encode/sort"
)

// Run function will encode our string
func Run(str, arg string) {
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
}
