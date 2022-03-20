package main

func smart(s, n, m int) int {
	if m < n || s*n < m {
		return 0
	}
	if n == 1 {
		return 1
	}

	result := 0
	for i := 1; i <= s; i++ {
		result += smart(s, n-1, m-i)
	}
	return result
}
