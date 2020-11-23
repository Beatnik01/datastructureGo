package datastruct

import "fmt"

// Node ...
type Node struct {
	Next *Node
	Prev *Node
	Val int
}

// LinkedList ...
type LinkedList struct {
	Root *Node
	Tail *Node
}

// 첫글자가 대문자면 외부로 공개, 소문자면 외부 비공개

// AddNode ...
func (l *LinkedList) AddNode(Val int) {
	if l.Root == nil { // Root가 nil이면 node가 없다는 뜻
		l.Root = &Node{Val:Val} // Root 생성
		l.Tail = l.Root // node가 1개뿐이니 Tail도 Root와 똑같음
		return
	}
	// Root가 이미 존재한다면
	l.Tail.Next = &Node{Val:Val} // Tail.Next 다음에 Node 생성
	Prev := l.Tail // node를 생성했으니 Prev에 Tail 주소를 넘김
	l.Tail = l.Tail.Next // l.Tail.Next 즉 새로 생성된 node를 Tail로 설정
	l.Tail.Prev = Prev // 새로 생성된 node의 Prev를 전 Tail 주소로 설정
	// 1   ->  2           1   ->   2
	// Tail Tail.Next  Tail.Prev  Tail 
}

// Back ...
func (l *LinkedList) Back() int {
	if l.Tail != nil {
		return l.Tail.Val
	}
	return 0
}

// Front ...
func (l *LinkedList) Front() int {
	if l.Root != nil {
		return l.Root.Val
	}
	return 0
}

// Empty ...
func (l *LinkedList) Empty() bool {
	return l.Root == nil
}

// PopBack ...
func (l *LinkedList) PopBack() {
	if l.Tail == nil {
		return
	}
	l.RemoveNode(l.Tail)
} 

// PopFront ...
func (l *LinkedList) PopFront() {
	if l.Root == nil {
		return
	}
	l.RemoveNode(l.Root)
}

// RemoveNode ...
func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.Root { // 입력받은 node가 Root면 
		l.Root = l.Root.Next // Root 다음 node의 주소를 Root에 넣어 Root를 변경한다.
		if l.Root != nil {
			l.Root.Prev = nil // Root는 첫 시작값이니 Prev가 없어야함
		}
		node.Next = nil // 입력받은 node의 Next 주소를 끊어, 모든 노드와의 연결을 삭제 = 연결된게 없으므로 Garbage Collection이 자동으로 삭제함 
		return
	}
	Prev := node.Prev // 입력받은 node의 Prev 주소를 받아옴

	if node == l.Tail { // 입력받은 node가 Tail이면
		Prev.Next = nil // Prev(위에서 입력 받은 node의 Prev 주소를 받아옴)의 Next를 nil로, 왜? Tail = 끝이니까 Next가 없음
		l.Tail.Prev = nil // Tail.Prev의 주소를 nil로 바꿔 연결을 끊어줌 Root 삭제 과정과 비슷한 이유
		l.Tail = Prev // 입력받은 node의 Prev가 새로운 Tail로 바뀜 ex) 9를 삭제해서 8이 새로운 Tail이 된다.
	} else { // Root도 아니고 Tail도 아니라면 중간에 있는 node를 삭제한다는 뜻
		node.Prev = nil // 입력받은 node의 Prev 주소를 nil로 바꿔 Prev 연결 삭제
		Prev.Next = Prev.Next.Next // Prev의 Next를 Next.Next 한 노드 건너뛴 주소로 바꿈
		Prev.Next.Prev = Prev // 한 노드 건너뛴 주소의 Prev 주소를 입력받은 node의 Prev 주소로 변경해서 Prev 연결을 끊음
		/* 헷갈리면 안되는것, 위에서 node.Prev를 받은 Prev와 node.Prev를 같다고 생각하면 안된다.
			Printf를 해보면 알겠지만 가지고 있는 값이 다름 Prev는 주소값을 가지고 있고 node.Prev는 nil을 가지고 있다.
		*/
		// RemoveNode(list.Root.Next)를 할 경우 0 -> 1 -> 2 -> 3 // 0 -> 2 -> 3으로 변경됨
		// Prev.Next = Prev.Next.Next를 통해 0의(0은 node의 Prev 이전 값이다.) Next 주소를 2로 바꾸고
		// Prev.Next.Prev 2의 Prev 주소를 1에서 0(Prev := node.Prev)으로 바꾼다. 
	}
	node.Next = nil // 마지막으로 node의 Next를 nil로 바꿔 입력받은 node 연결을 완벽하게 끊는다.
}

// PrintNodes ...
func (l *LinkedList) PrintNodes() {
	node := l.Root // node는 Root 
	for node.Next != nil { // Root부터 nil이 나올때까지 반복문을 돌린다.
		fmt.Printf("%d -> ", node.Val) // nil이 나올때까지 node.Val 출력
		node = node.Next // 출력하면 다음 주소를 넘겨줌, 여기서 nil이 나오면 반복문 중단
	}
	fmt.Printf("%d\n", node.Val)
}

// PrintReverse ...
func (l *LinkedList) PrintReverse() {
	node := l.Tail // node는 Tail
	for node.Prev != nil { // Tail부터 nil이 나올때까지 반복문을 돌린다.
		fmt.Printf("%d -> ", node.Val) // 나머지는 위에 Printnodes와 같음
		node = node.Prev
	}
	fmt.Printf("%d\n", node.Val)
}


