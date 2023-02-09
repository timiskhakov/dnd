package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLemire(t *testing.T) {
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("s = %d, n = %d, m = %d", c.s, c.n, c.m), func(t *testing.T) {
			assert.Equal(t, c.expected, Lemire(c.s, c.n, c.m))
		})
	}
}

func BenchmarkLemire(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Lemire(6, 10, 31)
	}
}
