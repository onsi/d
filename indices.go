package d

import (
	"math/rand"
	"time"
)

func Range(n int) []int {
	out := []int{}

	for i := 0; i < n; i += 1 {
		out = append(out, i)
	}

	return out
}

func Permutation(n int, seed ...int64) []int {
	var r *rand.Rand
	if len(seed) > 0 {
		r = rand.New(rand.NewSource(seed[0]))
	} else {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	return r.Perm(n)
}
