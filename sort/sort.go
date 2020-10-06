package sort

// InsertionSort function that sorts a pointer to string slice
func InsertionSort(items *[]string, n int) {
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

func QuickSort() {
	//Quick Sort function will be here
}
