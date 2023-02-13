package main

type item struct {
	s int
	n int
	m int
}

var cache = make(map[item]int)

func Cache(s, n, m int) int {
	if m < n || s*n < m {
		return 0
	}
	if n == 1 {
		return 1
	}

	i, ok := cache[item{s, n, m}]
	if ok {
		return i
	}

	result := 0
	for i := 1; i <= s; i++ {
		result += Cache(s, n-1, m-i)
	}

	cache[item{s, n, m}] = result

	return result
}
