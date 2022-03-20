package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinomial(t *testing.T) {
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("s = %d, n = %d, m = %d", c.s, c.n, c.m), func(t *testing.T) {
			assert.Equal(t, c.expected, binomial(c.s, c.n, c.m))
		})
	}
}

func BenchmarkBinomial(b *testing.B) {
	for n := 0; n < b.N; n++ {
		binomial(6, 10, 31)
	}
}
