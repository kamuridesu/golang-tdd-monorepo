package main

func Sum(a []int) int {
	var result int
	for _, value := range a {
		result += value
	}
	return result
}

func SumAll(slices ...[]int) []int {
	var sums []int
	for _, s := range slices {
		sums = append(sums, Sum(s))
	}
	return sums
}

func SumAllTails(slices ...[]int) []int {
	var sums []int
	for _, s := range slices {
		if len(s) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(s[1:]))
		}
	}
	return sums
}
