package main

import "fmt"

// Node ...
type Node struct {
	next *Node
	prev *Node
	val int
}

// LinkedList ...
type LinkedList struct {
	root *Node
	tail *Node
}

// AddNode ...
func (l *LinkedList) AddNode(val int) {
	if l.root == nil { // root가 nil이면 node가 없다는 뜻
		l.root = &Node{val:val} // root 생성
		l.tail = l.root // node가 1개뿐이니 tail도 root와 똑같음
		return
	}
	// root가 이미 존재한다면
	l.tail.next = &Node{val:val} // tail.next 다음에 Node 생성
	prev := l.tail // node를 생성했으니 prev에 tail 주소를 넘김
	l.tail = l.tail.next // l.tail.next 즉 새로 생성된 node를 tail로 설정
	l.tail.prev = prev // 새로 생성된 node의 prev를 전 tail 주소로 설정
	// 1   ->  2           1   ->   2
	// tail tail.next  tail.prev  tail 
}

// RemoveNode ...
func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.root { // 입력받은 node가 root면 
		l.root = l.root.next // root 다음 node의 주소를 root에 넣어 root를 변경한다.
		l.root.prev = nil // root는 첫 시작값이니 prev가 없어야함
		node.next = nil // 입력받은 node의 next 주소를 끊어, 모든 노드와의 연결을 삭제 = 연결된게 없으므로 Garbage Collection이 자동으로 삭제함 
		return
	}
	prev := node.prev // 입력받은 node의 prev 주소를 받아옴

	if node == l.tail { // 입력받은 node가 tail이면
		prev.next = nil // prev(위에서 입력 받은 node의 prev 주소를 받아옴)의 next를 nil로, 왜? tail = 끝이니까 next가 없음
		l.tail.prev = nil // tail.prev의 주소를 nil로 바꿔 연결을 끊어줌 root 삭제 과정과 비슷한 이유
		l.tail = prev // 입력받은 node의 prev가 새로운 tail로 바뀜 ex) 9를 삭제해서 8이 새로운 tail이 된다.
	} else { // root도 아니고 tail도 아니라면 중간에 있는 node를 삭제한다는 뜻
		node.prev = nil // 입력받은 node의 prev 주소를 nil로 바꿔 prev 연결 삭제
		prev.next = prev.next.next // prev의 next를 next.next 한 노드 건너뛴 주소로 바꿈
		prev.next.prev = prev // 한 노드 건너뛴 주소의 prev 주소를 입력받은 node의 prev 주소로 변경해서 prev 연결을 끊음
		/* 헷갈리면 안되는것, 위에서 node.prev를 받은 prev와 node.prev를 같다고 생각하면 안된다.
			Printf를 해보면 알겠지만 가지고 있는 값이 다름 prev는 주소값을 가지고 있고 node.prev는 nil을 가지고 있다.
		*/
		// RemoveNode(list.root.next)를 할 경우 0 -> 1 -> 2 -> 3 // 0 -> 2 -> 3으로 변경됨
		// prev.next = prev.next.next를 통해 0의(0은 node의 prev 이전 값이다.) next 주소를 2로 바꾸고
		// prev.next.prev 2의 prev 주소를 1에서 0(prev := node.prev)으로 바꾼다. 
	}
	node.next = nil // 마지막으로 node의 next를 nil로 바꿔 입력받은 node 연결을 완벽하게 끊는다.
}

// PrintNodes ...
func (l *LinkedList) PrintNodes() {
	node := l.root // node는 root 
	for node.next != nil { // root부터 nil이 나올때까지 반복문을 돌린다.
		fmt.Printf("%d -> ", node.val) // nil이 나올때까지 node.val 출력
		node = node.next // 출력하면 다음 주소를 넘겨줌, 여기서 nil이 나오면 반복문 중단
	}
	fmt.Printf("%d\n", node.val)
}

// PrintReverse ...
func (l *LinkedList) PrintReverse() {
	node := l.tail // node는 tail
	for node.prev != nil { // tail부터 nil이 나올때까지 반복문을 돌린다.
		fmt.Printf("%d -> ", node.val) // 나머지는 위에 Printnodes와 같음
		node = node.prev
	}
	fmt.Printf("%d\n", node.val)
}

func main() {
	list := &LinkedList{}
	list.AddNode(0) // root 생성

	for i := 1; i < 10; i++ { // 반복문을 돌려 1부터 9까지 노드 생성
		list.AddNode(i)
	}

	list.PrintNodes()

	list.RemoveNode(list.root.next)

	list.PrintNodes()

	list.RemoveNode(list.root)

	list.PrintNodes()

	list.RemoveNode(list.tail)

	list.PrintNodes()
	list.PrintReverse()
	fmt.Printf("tail : %d", list.tail.val)

}



