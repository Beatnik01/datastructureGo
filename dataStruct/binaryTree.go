package datastruct

import "fmt"

// BinaryTreeNode Val : 값 / Left : 현재 노드보다 작은 값 / Right : 현재 노드보다 큰 값
type BinaryTreeNode struct {
	Val int

	left  *BinaryTreeNode
	right *BinaryTreeNode
}

// BinaryTree Node의 Root 저장
type BinaryTree struct {
	Root *BinaryTreeNode
}

// NewBinaryTree Root Node 생성
func NewBinaryTree(val int) *BinaryTree {
	tree := &BinaryTree{}
	tree.Root = &BinaryTreeNode{Val: val}
	return tree
}

// AddNode 현재 Node의 Val이 입력받은 val보다 작다면 left, 크면 right에 AddNode
func (n *BinaryTreeNode) AddNode(val int) *BinaryTreeNode {
	if n.Val > val { // 입력받은 값이 현재 값보다 작다면
		if n.left == nil { // left가 없다면 생성하고 반환
			n.left = &BinaryTreeNode{Val: val}
			return n.left
		} else {
			return n.left.AddNode(val)
		}
	} else { // 입력받은 값이 현재 값보다 크다면
		if n.right == nil { // right가 없다면 생성하고 반환
			n.right = &BinaryTreeNode{Val: val} 
			return n.right
		} else {
			return n.right.AddNode(val)
		}
	}
}

// depthNode Print를 위한 깊이(depth) 구조체
type depthNode struct {
	depth int
	node  *BinaryTreeNode
}

// Print depth로 묶어서 한줄씩 출력
func (t *BinaryTree) Print() {
	q := []depthNode{}
	q = append(q, depthNode{depth: 0, node: t.Root})
	currentDepth := 0

	for len(q) > 0 {
		var first depthNode
		first, q = q[0], q[1:] // Queue

		if first.depth != currentDepth {
			fmt.Println()
			currentDepth = first.depth
		}
		fmt.Print(first.node.Val," ")

		if first.node.left != nil {
			q = append(q, depthNode{depth: currentDepth +1, node: first.node.left })
		}
		if first.node.right != nil {
			q = append(q, depthNode{depth: currentDepth + 1, node: first.node.right})
		}
	}
}
   
// Search ... 
func (t *BinaryTree) Search(val int)(bool, int) {
	return t.Root.Search(val, 1)
}

// Search 이진트리 안에 특정 값이 있는지 검색
func (n *BinaryTreeNode) Search(val int, cnt int) (bool, int) {
	if n.Val == val {
		return true, cnt
	} else if n.Val > val {
		if n.left != nil {
			return n.left.Search(val, cnt+1)
		}
		return false,cnt
	}else {
		if n.right != nil {
			return n.right.Search(val, cnt+1)
		}
		return false,cnt
	}
}