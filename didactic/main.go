package main

import (
	"fmt"
	"iter"
)

func Tasks() iter.Seq[string] {
	fmt.Println("2. Tasks creates and returns the sequence.")

	// yield is a callback named for yielding each sequence value to range.
	return func(yield func(string) bool) {
		fmt.Println("3. range runs the sequence and passes the yield callback.")

		fmt.Println("4. The sequence calls yield with task A.")
		keepGoing := yield("A")
		if !keepGoing {
			fmt.Println("The sequence stops because the outer loop stopped.")

			return
		}
		fmt.Println("6. yield returned true, so the sequence continues.")

		fmt.Println("7. The sequence calls yield with task B.")
		keepGoing = yield("B")
		if !keepGoing {
			fmt.Println("The sequence stops because the outer loop stopped.")

			return
		}
		fmt.Println("9. yield returned true, so the sequence continues.")

		fmt.Println("10. The sequence sent all tasks.")
	}
}

func main() {
	fmt.Println("1. The outer loop calls Tasks to get the sequence.")

	receivedTasks := 0
	for task := range Tasks() {
		receivedTasks++
		fmt.Printf("The outer loop receives task %d: %q.\n", receivedTasks, task)
	}

	fmt.Println("11. The outer loop ends because the sequence returned.")
}
