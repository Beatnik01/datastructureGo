package main

import (
	"datastruct"
	"fmt"
)
func main() {
	tree := datastruct.Tree{}

	val := 1
	tree.AddNode(val)
	val++

	for i := 0; i < 3; i++ {
		tree.Root.AddNode(val)
		val++
	}

	for i := 0; i < len(tree.Root.Childs); i++ {
		for j := 0; j < 2; j++ {
			tree.Root.Childs[i].AddNode(val)
			val++
		}
	}
	
	tree.DFS1();
	fmt.Println("DFS2 Recursive");
	
	tree.DFS2();
	fmt.Println("DFS2 Stack LinkedList");
	
	tree.BFS();
	fmt.Println("BFS Stack")
	/*
	stack := []int{}

	for i := 1; i <= 5; i++ {
		stack = append(stack,i)
	}
	for len(stack) > 0 {
		var last int
		last, stack = stack[len(stack)-1], stack[:len(stack)-1]
		fmt.Println(last)
	}

	queue := []int{}
	for i := 1; i <= 5; i++ {
		queue = append(queue, i)
	}

	fmt.Println(queue)

	for len(queue) > 0 {
		var front int
		front, queue = queue[0], queue[1:]
		fmt.Println(front)
	}

	stack2 := datastruct.NewStack()

	for i := 1; i <= 5; i++ {
		stack2.Push(i)
	}

	fmt.Println("NewStack")

	for !stack2.Empty() {
		val := stack2.Pop()
		fmt.Printf("%d -> ", val)
	}

	queue2 := datastruct.NewQueue()
	for i := 1; i <= 5; i++ {
		queue2.Push(i)
	}

	fmt.Println("NewQueue")

	for !queue2.Empty() {
		val := queue2.Pop()
		fmt.Printf("%d -> ", val)
	}
	*/
}