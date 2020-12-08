package main

import (
	"datastruct"
	"fmt"
)

func main() {
	// 같은 값은 같은 결과가 나와야한다. 다른 값은 다른 결과가 나와야한다.
	fmt.Println("abcde = ", datastruct.Hash("abcde"))
	fmt.Println("abcdf = ", datastruct.Hash("abcdf"))
	
	m := datastruct.CreateMap()
	m.Add("AAA", "0107777777")
	m.Add("BBB", "0108888888")
	m.Add("CCC", "0109999999")

	fmt.Println("AAA = ",m.Get("AAA"))
	fmt.Println("CCC = ",m.Get("CCC"))
}

// Tree
/*
func main() {
	tree := datastruct.NewBinaryTree(5)
	tree.Root.AddNode(3)
	tree.Root.AddNode(2)
	tree.Root.AddNode(4)
	tree.Root.AddNode(8)
	tree.Root.AddNode(7)
	tree.Root.AddNode(6)
	tree.Root.AddNode(10)
	tree.Root.AddNode(9)

	tree.Print()
	fmt.Println()

	if found, cnt := tree.Search(6);found {
		fmt.Println("found 6 cnt", cnt)
	} else {
		fmt.Println("not found 6 cnt", cnt)
	}

	if found, cnt := tree.Search(11);found {
		fmt.Println("found 11 cnt", cnt)
	} else {
		fmt.Println("not found 11 cnt", cnt)
	}
}
*/
