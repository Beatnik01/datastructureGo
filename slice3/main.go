package main

import "fmt"

func main() {
	stack := []int{}
	for i := 1; i <= 5; i++ {
		stack = append(stack, i)
	}
	fmt.Println(stack)

	var last int
	for len(stack) > 0 {
		last, stack = stack[len(stack)-1], stack[:len(stack)-1]
		fmt.Println(last)
	}

	queue := []int{}
	for i := 1; i <= 5; i++ {
		queue = append(queue, i)
	}
	fmt.Println(queue)

	var front int
	for len(queue) > 0 {
		front, queue = queue[0], queue[1:]
		fmt.Println(front)
	}
}