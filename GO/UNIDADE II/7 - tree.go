package main

import "fmt"

type BSTNode struct {
	left  *BSTNode
	val   int
	right *BSTNode
}

func createBSTNode(val int) *BSTNode {
	return &BSTNode{val: val}
}

func (node *BSTNode) Add(val int) {
	if val <= node.val {
		if node.left != nil {
			node.left.Add(val)
		} else {
			node.left = createBSTNode(val)
		}

	} else {
		if node.right != nil {
			node.right.Add(val)
		} else {
			node.right = createBSTNode(val)
		}
	}
}

func (node *BSTNode) Search(val int) bool {
	if val == node.val {
		return true
	} else if val < node.val {
		if node.left != nil {
			return node.left.Search(val)
		} else {
			return false
		}
	} else {
		if node.right != nil {
			return node.right.Search(val)
		} else {
			return false
		}
	}
}

func (node *BSTNode) Min() int {
	if node.left != nil {
		return node.left.Min()
	} else {
		return node.val
	}
}

func (node *BSTNode) Max() int {
	if node.right != nil {
		return node.right.Max()
	} else {
		return node.val
	}
}

func (node *BSTNode) PreOrderNav() {
	print(node.val, ", ")
	if node.left != nil {
		node.left.PreOrderNav()
	}
	if node.right != nil {
		node.right.PreOrderNav()
	}
}

func (node *BSTNode) InOrderNav() {
	if node.left != nil {
		node.left.InOrderNav()
	}
	print(node.val, ", ")
	if node.right != nil {
		node.right.InOrderNav()
	}
}

func (node *BSTNode) PostOrderNav() {
	if node.left != nil {
		node.left.PostOrderNav()
	}
	if node.right != nil {
		node.right.PostOrderNav()
	}
	print(node.val, ", ")
}

func (node *BSTNode) LevelOrderNav() {
	queue := make([]*BSTNode, 0)
	queue = append(queue, node)

	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		print(curNode.val, ", ")
		if curNode.left != nil {
			queue = append(queue, curNode.left)
		}
		if curNode.right != nil {
			queue = append(queue, curNode.right)
		}
	}
}

func (node *BSTNode) Height() int {
	leftHeight := 0
	if node.left != nil {
		leftHeight = node.left.Height()
	}
	rightHeight := 0
	if node.right != nil {
		rightHeight = node.right.Height()
	}
	if node.left == nil && node.right == nil {
		return 0
	} else {
		if leftHeight > rightHeight {
			return 1 + leftHeight
		} else {
			return 1 + rightHeight
		}
	}
}

func (node *BSTNode) Remove(val int) *BSTNode {
	if val < node.val && node.left != nil {
		node.left = node.left.Remove(val)
	} else if val > node.val && node.right != nil {
		node.right = node.right.Remove(val)
	} else {
		if node.left == nil && node.right == nil {
			return nil
		} else if node.left != nil && node.right == nil {
			return node.left
		} else if node.left == nil && node.right != nil {
			return node.right
		} else {
			node.val = node.left.Max()
			node.left = node.left.Remove(node.val)
			return node
		}
	}
	return node
}

func (bstNode *BSTNode) SizeEven() int {
	size := 0
	if bstNode.val%2 == 0 {
		size += 1
	}
	if bstNode.left != nil {
		size += bstNode.left.SizeEven()
	}
	if bstNode.right != nil {
		size += bstNode.right.SizeEven()
	}
	return size
}

func (bstNode *BSTNode) Size() int {
	size := 1
	if bstNode.left != nil {
		size += bstNode.left.Size()
	}
	if bstNode.right != nil {
		size += bstNode.right.Size()
	}
	return size
}

func (bstNode *BSTNode) IsBst() bool {
	if bstNode.left != nil {
		if bstNode.val > bstNode.left.val {
			return bstNode.left.IsBst()
		} else {
			return false
		}
	}
	if bstNode.right != nil {
		if bstNode.val < bstNode.right.val {
			return bstNode.right.IsBst()
		} else {
			return false
		}
	}
	return true
}

func main() {
	bst := BSTNode{nil, 5, nil}

	bst.Add(50)
	bst.Add(51)
	bst.Add(-1)
	bst.Add(48)
	bst.Add(52)
	bst.Add(1000)

	fmt.Println(bst.IsBst())
}
