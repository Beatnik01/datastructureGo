package datastruct

import "fmt"

// TreeNode Val - 값 / Childs - 자식 노드, TreeNode의 포인터 참조
type TreeNode struct {
	Val    int 
	Childs []*TreeNode 
}

// Tree TreeNode의 Root를 나타낸다.
type Tree struct {
	Root *TreeNode
} 

// AddNode ... 
func  (t *TreeNode) AddNode(val int) {
	t.Childs = append(t.Childs, &TreeNode{Val: val})
}

// AddNode val을 받고 *Tree로 반환
func (t *Tree) AddNode(val int) {
	if t.Root == nil { // nil이면 Root가 없다는 뜻이니 새로 생성한다.
		t.Root = &TreeNode{Val: val}
	} else { // 아닐경우 Root가 있다는 뜻이니 Root의 Childs에 새로운 Node를 생성한다.
		t.Root.Childs = append(t.Root.Childs, &TreeNode{Val: val})
	}
}

// DFS1 is DFS1 method 
func (t *Tree) DFS1() {
	DFS1(t.Root)
}

// DFS1 is Tree DFS Recursive Function 
func DFS1(node *TreeNode) {
	fmt.Printf("%d -> ", node.Val)

	for i := 0; i < len(node.Childs); i++ {
		DFS1(node.Childs[i])
	}
}

// DFS2 is Tree DFS Stack method
func (t *Tree) DFS2() {
	s := []*TreeNode{} 
	s = append(s, t.Root)

	for len(s) > 0 {
		var last *TreeNode
		last, s = s[len(s)-1], s[:len(s)-1]

		fmt.Printf("%d -> ", last.Val)

		for i := len(last.Childs)-1; i >= 0; i-- {
			s = append(s, last.Childs[i])
		}
	}
}

// BFS ...
func (t *Tree) BFS() {
	q := []*TreeNode{}
	q = append(q, t.Root)

	for len(q) > 0 {
		var first *TreeNode
		first, q = q[0], q[1:]
		
		fmt.Printf("%d -> ", first.Val)
		for i := 0; i < len(first.Childs); i++ {
			q = append(q, first.Childs[i])
		}
	}
}
