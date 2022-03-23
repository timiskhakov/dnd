package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNaive(t *testing.T) {
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("s = %d, n = %d, m = %d", c.s, c.n, c.m), func(t *testing.T) {
			assert.Equal(t, c.expected, Naive(c.s, c.n, c.m))
		})
	}
}

func BenchmarkNaive(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Naive(6, 10, 31)
	}
}
