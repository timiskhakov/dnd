package main

func Lemire(s, n, m int) int {
	result := 0
	i := 0
	for s*i <= m-n {
		j := m - n - i*s
		result += sign(i+j) * fastcoef(n, i) * sign(j) * fastcoef(n+j-1, j)
		i++
	}

	return result
}

func fastcoef(n, k int) int {
	if k == 0 {
		return 1
	}

	np := n - k
	c := np + 1
	for i := 2; i <= k; i++ {
		c *= np + i
		f := precomputed[i]
		c >>= f.shift
		c = int(uint64(c) * f.inverse)
	}

	return c
}

type fastdiv struct {
	shift   int
	inverse uint64
}

var precomputed = []fastdiv{
	{0, 0},
	{0, 0x1},
	{1, 0x1},
	{0, 0xaaaaaaaaaaaaaaab},
	{2, 0x1},
	{0, 0xcccccccccccccccd},
	{1, 0xaaaaaaaaaaaaaaab},
	{0, 0x6db6db6db6db6db7},
	{3, 0x1},
	{0, 0x8e38e38e38e38e39},
	{1, 0xcccccccccccccccd},
	{0, 0x2e8ba2e8ba2e8ba3},
	{2, 0xaaaaaaaaaaaaaaab},
	{0, 0x4ec4ec4ec4ec4ec5},
	{1, 0x6db6db6db6db6db7},
	{0, 0xeeeeeeeeeeeeeeef},
	{4, 0x1},
	{0, 0xf0f0f0f0f0f0f0f1},
	{1, 0x8e38e38e38e38e39},
	{0, 0x86bca1af286bca1b},
	{2, 0xcccccccccccccccd},
	{0, 0xcf3cf3cf3cf3cf3d},
	{1, 0x2e8ba2e8ba2e8ba3},
	{0, 0xd37a6f4de9bd37a7},
	{3, 0xaaaaaaaaaaaaaaab},
	{0, 0x8f5c28f5c28f5c29},
	{1, 0x4ec4ec4ec4ec4ec5},
	{0, 0x84bda12f684bda13},
	{2, 0x6db6db6db6db6db7},
	{0, 0x34f72c234f72c235},
	{1, 0xeeeeeeeeeeeeeeef},
	{0, 0xef7bdef7bdef7bdf},
	{5, 0x1},
	{0, 0xf83e0f83e0f83e1},
	{1, 0xf0f0f0f0f0f0f0f1},
	{0, 0xaf8af8af8af8af8b},
	{2, 0x8e38e38e38e38e39},
	{0, 0x14c1bacf914c1bad},
	{1, 0x86bca1af286bca1b},
	{0, 0x6f96f96f96f96f97},
	{3, 0xcccccccccccccccd},
	{0, 0x8f9c18f9c18f9c19},
	{1, 0xcf3cf3cf3cf3cf3d},
	{0, 0x82fa0be82fa0be83},
	{2, 0x2e8ba2e8ba2e8ba3},
	{0, 0x4fa4fa4fa4fa4fa5},
	{1, 0xd37a6f4de9bd37a7},
	{0, 0x51b3bea3677d46cf},
	{4, 0xaaaaaaaaaaaaaaab},
	{0, 0x7d6343eb1a1f58d1},
	{1, 0x8f5c28f5c28f5c29},
	{0, 0xfafafafafafafafb},
	{2, 0x4ec4ec4ec4ec4ec5},
	{0, 0x21cfb2b78c13521d},
	{1, 0x84bda12f684bda13},
	{0, 0x6fb586fb586fb587},
	{3, 0x6db6db6db6db6db7},
	{0, 0x823ee08fb823ee09},
	{1, 0x34f72c234f72c235},
	{0, 0xcbeea4e1a08ad8f3},
	{2, 0xeeeeeeeeeeeeeeef},
	{0, 0x4fbcda3ac10c9715},
	{1, 0xef7bdef7bdef7bdf},
	{0, 0xefbefbefbefbefbf},
}
