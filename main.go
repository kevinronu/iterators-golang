package main

import (
	"cmp"
	"context"
	"fmt"
	"iter"
	"time"
)

type ChannelIterator[T cmp.Ordered] struct {
	Data []T
}

func (ci *ChannelIterator[T]) GetValuesWithChannel(ctx context.Context) <-chan T {
	ch := make(chan T)

	go func() {
		defer close(ch)
		defer fmt.Println("Producer fully stopped")

		for i := 0; i < len(ci.Data); i++ { // Producer will stop after 10 iterations
			select {
			case <-ctx.Done(): // To have control over the completion
				fmt.Println("Producer received cancel signal")
				return
			default:
				ch <- ci.Data[i]
				fmt.Println("Produced:", i)
				time.Sleep(500 * time.Millisecond) // Simulate work
			}

			// Keep working in between producing values
			for j := 0; j < 3; j++ { // Simulate some background work
				select {
				case <-ctx.Done():
					fmt.Println("Background work stopped")
					return
				default:
					fmt.Println("Working in background...")
					time.Sleep(200 * time.Millisecond)
				}
			}
		}
	}()

	return ch
}

func IterateWithChannelWithCancel[T cmp.Ordered](ctx context.Context, ci *ChannelIterator[T]) {
	ch := ci.GetValuesWithChannel(ctx)

	var zero T

	for value := range ch {
		fmt.Println("Consumed:", value)
		if value == zero {
			break // Stop receiving, but producer continues
		}
	}

	fmt.Println("Receiver stopped, but producer continues working...")
	time.Sleep(2 * time.Second) // Observe producer activity
}

func IterateWithChannelWithoutCancel[T cmp.Ordered](ctx context.Context, cancelFunc context.CancelFunc, ci *ChannelIterator[T]) {
	ch := ci.GetValuesWithChannel(ctx)

	var zero T

	for value := range ch {
		fmt.Println("Consumed:", value)
		if value == zero {
			cancelFunc() // Stop both receiver and producer
			break
		}
	}

	fmt.Println("Receiver and producer stopped.")
	time.Sleep(2 * time.Second) // Observe no producer activity
}

func (ci *ChannelIterator[T]) GetValuesWithIterator() iter.Seq[T] {
	return func(yield func(T) bool) { // yield is only a name, this can be named callback
		for i := 0; i < len(ci.Data); i++ {
			fmt.Println("Produced:", ci.Data[i])
			keepGoing := yield(ci.Data[i])
			if !keepGoing {
				return
			}
		}
	}
}

func main() {
	ci := &ChannelIterator[int]{
		Data: []int{-1, 0, 1},
	}

	// fmt.Println("Example without cancel:")
	// IterateWithChannelWithCancel(context.Background(), ci) // No cancellation context

	// fmt.Println("\nExample with cancel:")
	// ctx, cancel := context.WithCancel(context.Background())
	// IterateWithChannelWithoutCancel(ctx, cancel, ci)

	fmt.Println("\nExample with iterators:")
	for value := range ci.GetValuesWithIterator() {
		fmt.Println("Consumed:", value)
		if value == 0 {
			break
		}
	}
}
