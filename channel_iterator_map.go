package main

import (
	"cmp"
	"iter"
)

type ChannelIteratorMap[K cmp.Ordered, V cmp.Ordered] struct {
	data map[K]V
}

func (cimp *ChannelIteratorMap[K, V]) GetValues() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range cimp.data {
			keepGoing := yield(k, v)

			if !keepGoing {
				return
			}
		}
	}
}
