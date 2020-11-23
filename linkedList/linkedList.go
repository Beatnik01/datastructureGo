package main

import "fmt"

// Node ...
type Node struct {
	next *Node
	val  int
}

func main() {
	var root *Node
	var tail *Node

	root = &Node{val: 0} // root Node 생성
	tail = root // tail은 root Node를 가르킴

	for i := 1; i < 10; i++ {
		tail = AddNode(tail, i)
	}

	// 왜 PrintNodes(root)를 했을때 1~8까지 노드가 나오고, tail을 했을때는 9가 나올까?
	/* 
		-> PrintNodes 함수를 보면 변수 node에 root를 초기값으로 넣고 node.next를 nil이 나올때까지
		for문을 돌리기 때문 때문에 tail 제외, nil 전까지 있는 모든 노드들이 출력되는 것 
	*/

	PrintNodes(root)
	
	root, tail = RemoveNode(root.next, root, tail)

	PrintNodes(root)

	root, tail = RemoveNode(root, root, tail)

	PrintNodes(root)

	root, tail = RemoveNode(tail, root, tail)
	PrintNodes(root)
	fmt.Printf("tail: %d\n", tail.val)
}

// AddNode ...
func AddNode(tail *Node, val1 int) *Node {
	node := &Node{val:val1} // Node.val에 함수로 받아온 val 값을 삽입해 노드 생성
	tail.next = node // 새로운 노드가 생겼으니 tail을 옮김
	return node
}

// RemoveNode ...
func RemoveNode(node *Node, root *Node, tail *Node) (*Node, *Node) {
	if node == root { // 지우고자 하는 node가 root인 경우
		root = root.next // 새로운 root는 기존 root의 next로 변경
		if root == nil { // root가 없을 경우(nil)
			tail = nil // tail도 nil로 변경 => 값이 없다는 거니까
		}
		return root, tail
	}
	
	prev := root
	for prev.next != node { // prev.next가 node값이면 for문 빠져나옴
		prev = prev.next // for문을 빠져나오면서 이 부분은 실행 안되고 딱 node의 전 node의 next값을 가지게됨
	}

	if node == tail { // node와 tail이 같다면 끝부터 지운다는 뜻
		prev.next = nil // 끝부터 지워지니 prev.next에 nil을 넣고
		tail = prev // tail에 prev 값을 넣어서 tail을 변경함
	}else{
		prev.next = prev.next.next // 아니라면 중간값을 삭제한다는 뜻이므로 건너띄기로 삭제
	}
	
	return root, tail
}

// PrintNodes ...
func PrintNodes(root *Node) {
	node := root // node에 root값 삽입
	for node.next != nil { // root부터 nil 나오기 전까지 반복문을 통해 출력함
		fmt.Printf("%d -> ", node.val)
		node = node.next // 다음 node값을 불러옴
	}
	fmt.Printf("%d\n", node.val)
}