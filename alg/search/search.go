package search

func BinarySearch(sortedArr []int64, element int64) int {
	s := 0
	e := len(sortedArr) - 1

	for s < e {
		m := s + (e-s)/2

		if sortedArr[m] == element {
			return m
		}

		if sortedArr[m] < element {
			s = m + 1
		}else {
			e = m
		}
	}

	return -1
}
