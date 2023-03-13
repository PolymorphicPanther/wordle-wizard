package ds_utils

import "sort"

func SortSlice(s []byte) {
	sort.Slice(s, func(i int, j int) bool {
		return s[i] < s[j]
	})
}

func SliceToSet[K comparable, T, V any](s []T, keyFn func(T) K, valFn func(T) V) map[K]V {
	m := make(map[K]V, len(s))

	for _, t := range s {
		m[keyFn(t)] = valFn(t)
	}
	return m
}
