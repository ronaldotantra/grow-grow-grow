package main

// https://www.codewars.com/kata/55a5c82cd8e9baa49000004c

func DivisibleCount(x, y, k uint64) uint64 {
	for i := x; i <= y; i++ {
		if i%k == 0 {
			return (y-i)/k + 1
		}
	}
	return 0
}
