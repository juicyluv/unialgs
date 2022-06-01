package mysort

func InsertionSort(items []int) []int {
	for i := 1; i < len(items); i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}

	return items
}
