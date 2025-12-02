package main

import "fmt"

func main() {
	// 1. Classic for Loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 2. While-like Loop
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// 3. Infinite Loop
	// Used when running tasks continuously until an explicit break.
	// for {
	//     fmt.Println("Running forever...")
	// }

	// 4. Loop Control Keywords
	// break — exit the loop immediately
	// continue — skip to the next iteration
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		if i == 4 {
			break
		}
		fmt.Println(i)
	}

	// 6. Looping Over Collections with range
	// range is used to iterate over arrays, slices, strings, maps, and channels.
	numbers := []int{1, 2, 3}
	for index, value := range numbers {
		fmt.Println(index, value)
	}

	// When you don't need the index or value, you can use _
	for _, v := range numbers {
		fmt.Println(v)
	}

	// 7. Loop Labels
	// Used to break or continue outer loops.
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				break outer
			}
		}
	}
}
