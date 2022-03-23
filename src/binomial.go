package main

func Binomial(s, n, m int) int {
	result := 0
	i := 0
	for s*i <= m-n {
		j := m - n - i*s
		result += sign(i+j) * coef(n, i) * sign(j) * coef(n+j-1, j)
		i++
	}
	return result
}

func coef(n, k int) int {
	c := 1
	for i := 0; i < min(k, n-k); i++ {
		c *= (n - i) / (i + 1)
	}
	return c
}

func sign(n int) int {
	if n%2 == 0 {
		return 1
	}
	return -1
}

func min(n, k int) int {
	if n < k {
		return n
	}
	return k
}
