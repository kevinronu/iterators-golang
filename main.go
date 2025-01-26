package main

import (
	"context"
	"fmt"
)

type ChannelIterator[T any] struct {
	data []T
}

func (ci *ChannelIterator[T]) GetValues(ctx context.Context) <-chan T {
	ch := make(chan T)

	go func() {
		defer close(ch)

		select {
		case <-ctx.Done(): // To have control over the completion
			return
		default:
			for i := 0; i < len(ci.data); i++ {
				ch <- ci.data[i]
			}
		}
	}()

	return ch
}

func main() {
	ci := &ChannelIterator[int]{
		data: []int{0, 1, 2, 3, 4},
	}

	for value := range ci.GetValues(context.Background()) {
		fmt.Println(value)
	}
}
