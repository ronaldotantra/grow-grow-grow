package main

// https://www.codewars.com/kata/62c93765cef6f10030dfa92b

func Cats(start, finish int) int {
	result := (finish - start) / 3
	result += (finish - start) % 3
	return result
}
