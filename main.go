package main

import (
	"fmt"
)

func main() {
	ci := &ChannelIterator[int]{
		data: []int{-1, 0, 1},
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

	cim := &ChannelIteratorMap[int, string]{
		data: map[int]string{
			-1: "negative one",
			0:  "zero",
			1:  "one",
		},
	}

	fmt.Println("\nExample with iterators using two values:")
	for k, v := range cim.GetValues() {
		fmt.Printf("%d: %s\n", k, v)
		if k == 0 {
			break
		}
	}
}
