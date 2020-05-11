package quicksort

func qsort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	var less []int
	var more []int
	for _, v := range arr[1:] {
		if v <= pivot {
			less = append(less, v)
		} else {
			more = append(more, v)
		}
	}
	return append(append(qsort(less), pivot), qsort(more)...)
}
