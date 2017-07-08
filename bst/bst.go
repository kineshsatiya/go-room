package bst

type Node struct {
	value  int
	parent *Node
	left   *Node
	right  *Node
}

type BST struct {
	root          *Node
	totalElements int
}

// create new binary search tree
func New() *BST {
	return &BST{}
}

// insert an integer in the given binary search tree
func (bst *BST) Insert(item int) {
	newNode := &Node{value: item}
	bst.totalElements++
	var parent *Node
	current := bst.root
	for current != nil {
		// allow duplicates?
		if item <= current.value {
			parent, current = current, current.left
		} else {
			parent, current = current, current.right
		}
	}

	if parent != nil {
		newNode.parent = parent
		if newNode.value <= parent.value {
			parent.left = newNode
		} else {
			parent.right = newNode
		}
	} else {
		bst.root = newNode
		bst.root.parent = nil
	}
}

// find the given integer in binary search tree
func (b *BST) Contains(item int) bool {
	current := b.root
	for current != nil {
		// if duplicates allowed then first match returns
		if current.value == item {
			return true
		} else {
			if item < current.value {
				current = current.left
			} else {
				current = current.right
			}
		}
	}
	return false
}

// iterative solution
func (b *BST) InOrder() []int {
	if b.root == nil {
		return nil
	}
	inOrder := make([]int, 0)
	stack := make([]*Node, 0)
	current := b.root

	for len(stack) != 0 || current != nil {
		if current != nil {
			stack = append(stack, current)
			current = current.left
		} else {
			l := len(stack) - 1
			current, stack = stack[l], stack[:l] // pop & shrink stack
			inOrder = append(inOrder, current.value)
			current = current.right
		}
	}
	return inOrder
}
