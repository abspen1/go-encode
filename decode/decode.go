package decode

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/abspen1/go-encode/output"
	"github.com/abspen1/go-encode/sort"
)

// Run function that will decode our message
func Run(arg string) {
	scanner := bufio.NewScanner(os.Stdin)
	var index int
	var str string
	lines := 0
	/**
	 * The reason for this flag is to control when the entire line consists of just spaces
	 * For some reason which I haven't figured out, in the Anne of Avonlea large txt file
	 * one of the lines is just 100 spaces. This was causing my decoder to output binary
	 * for that line. This flag is avoiding that error by checking if the entire line
	 * only consists of spaces.
	 */
	// While there is input in the txt file, set str equal to the current line
	for scanner.Scan() {
		str = scanner.Text()
		// If the line isn't empty, decode the line
		if len(str) > 0 {
			// Increment number of lines, this will help us keep track of non-empty lines
			lines++

			if lines%2 != 0 {
				// Convert given index to an int type
				if tmpIndex, err := strconv.Atoi(str); err == nil {
					index = tmpIndex
				} else {
					fmt.Println(str, "is not an integer.")
					return
				}
			} else {
				//Ignore new line for the first line
				if lines != 2 {
					fmt.Println()
				}

				// Initialize variables we will be using
				n := len(str)
				flag := false
				last := make([]string, n)
				charCount := 0
				j := 0
				var bigNum string
				var num int
				sum := 0

				for i := 0; i < n; i++ {
					// If j is an even number then we are looking for an Int
					if j%2 == 0 {
						// This will check to see if the number of letters is more than 9
						if str[i+1] != ' ' {
							bigNum = string(str[i])
							s := 1
							// While loop that ends when there is a space, allows for multiple digit numbers
							for str[i+s] != ' ' {
								bigNum = bigNum + string(str[i+s])
								s++
							}
							// increment i for the amount of times the while loop went - 1
							i += (s - 1)
							num, _ = strconv.Atoi(bigNum) // typecast bigNum from String to Int
							sum += num
						} else {
							num, _ = strconv.Atoi(string(str[i]))  // reset each time so we know how many letters
							add, _ := strconv.Atoi(string(str[i])) // keeps track of total length of str array
							sum += add
						}

					} else { // j is odd so we are looking for a char
						if str[i] != ' ' {
							// If the char isn't a space then set the flag to true
							flag = true
						}
						for k := 0; k < num; k++ {
							// decompressing the string
							last[charCount] = string(str[i])
							charCount++
						}
					}
					i++ // second incrementer to skip over spaces
					j++ // increment for just numbers and letters in string
				}

				first := make([]string, sum) // Initialize first string array
				next := make([]int, sum)     // initialize next int array

				if flag == false {
					// If our entire line was spaces, just print out the line
					for k := 0; k < num; k++ {
						fmt.Print(last[k])
					}
				} else {
					for i := 0; i < sum; i++ {
						first[i] = last[i]
						// next[i] = -1 // reason for this, numNotInNext was giving false positives
					}
					// Check keyword for which sorting algorithm to use
					if arg == "quick" {
						fmt.Println("Working on quick sort")
						return
					} else if arg == "insertion" {
						sort.InsertionSort(&first, sum) // Pass the address of our string array
					} else {
						fmt.Println("invalid command line argument")
						return // return since an error occurred
					}

					var letter string
					// Getting the indices
					for i := 0; i < sum; i++ {
						letter = first[i]
						for k := 0; k < sum; k++ {
							if letter == last[k] && numNotInNext(k, next, sum) {
								next[i] = k
								break
							}
						}
					}
					// Print our decoded line
					output.PrintDecodedLine(first, next, sum, index)
				}
			}
		} else {
			// If the line is empty, print an empty line
			fmt.Println()
		}
	}
}

func numNotInNext(num int, next []int, n int) bool {
	for i := 0; i < n; i++ {
		if num == next[i] {
			return false
		}
	}
	return true
}
