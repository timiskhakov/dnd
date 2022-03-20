package main

func naive(s, n, m int) int {
	values := make([]int, s)
	for i := 0; i < len(values); i++ {
		values[i] = i + 1
	}

	state := make([]int, n)
	for i := 0; i < len(state); i++ {
		state[i] = i + 1
	}

	return bruteforce(values, state, n, 0, m)
}

func bruteforce(values, state []int, limit, index, m int) int {
	result := 0
	if index == limit {
		if sum(state) == m {
			result++
		}
	} else {
		for i := 0; i < len(values); i++ {
			state[index] = values[i]
			result += bruteforce(values, state, limit, index+1, m)
		}
	}
	return result
}

func sum(state []int) int {
	result := 0
	for i := 0; i < len(state); i++ {
		result += state[i]
	}
	return result
}
