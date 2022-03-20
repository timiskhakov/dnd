package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSmart(t *testing.T) {
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("s = %d, n = %d, m = %d", c.s, c.n, c.m), func(t *testing.T) {
			assert.Equal(t, c.expected, smart(c.s, c.n, c.m))
		})
	}
}

func BenchmarkSmart(b *testing.B) {
	for n := 0; n < b.N; n++ {
		smart(6, 10, 31)
	}
}
