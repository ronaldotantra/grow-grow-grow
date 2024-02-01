package main

func FindOdd(seq []int) int {
	mapSeq := make(map[int]int)
	for _, num := range seq {
		mapSeq[num]++
	}
	result := 0
	for key, value := range mapSeq {
		if value%2 != 0 {
			result = key
		}
	}
	return result
}
