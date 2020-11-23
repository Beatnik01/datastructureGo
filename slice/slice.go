package main

import "fmt"

func main() {
	a := make([]int, 0, 8)
	// int / length(실제 사용하는 공간) / capacity(확보한 공간)

	fmt.Printf("len(a) = %d\n", len(a))
	fmt.Printf("cap(a) = %d\n", cap(a))
	fmt.Printf("-------------\n")

	a = append(a, 1)
	// append(a라는 slice에, 1을 추가하고(length)) -> a slice에 넣어라
	// -> 다른 슬라이스에 넣을 수 있다는 뜻

	fmt.Printf("len(a) = %d\n", len(a))
	fmt.Printf("cap(a) = %d\n", cap(a))
	fmt.Printf("-------------\n")

	b := make([]int, 3, 3)

	fmt.Printf("len(b) = %d\n", len(b))
	fmt.Printf("cap(b) = %d\n", cap(b))
	fmt.Printf("-------------\n")

	a = append(b, 3)

	fmt.Printf("len(a) = %d\n", len(a))
	fmt.Printf("cap(a) = %d\n", cap(a))
	fmt.Printf("a = %d\n", a)
	fmt.Printf("b = %d\n", b)
}