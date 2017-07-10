package tst

// Node struct
type Node struct {
	values []interface{}
	key    byte
	left   *Node
	middle *Node
	right  *Node
	parent *Node
}

// TST struct
type TST struct {
	root *Node
}

// New TST
func New() *TST {
	return &TST{}
}

// Insert new element in TST, stores the value at every prefix
func (tst *TST) Insert(keys []byte, value interface{}) {
	var parent *Node
	current := tst.root
	for _, k := range keys {
		if current != nil {
			for current.key != k {
				// pick the correct child
				var child *Node
				if k < current.key {
					if current.left == nil {
						current.left = &Node{values: make([]interface{}, 0)}
						current.left.key, current.left.parent = k, current
					}
					child = current.left
				} else if k > current.key {
					if current.right == nil {
						current.right = &Node{values: make([]interface{}, 0)}
						current.right.key, current.right.parent = k, current
					}
					child = current.right
				} else {
					if current.middle == nil {
						current.middle = &Node{values: make([]interface{}, 0)}
						current.middle.key, current.middle.parent = k, current
					}
					child = current.middle
				}
				parent, current = current, child
			}
			// append the value to current node values and continue iteration with next key
			current.values = append(current.values, value)
		} else {
			// if the root was not initialized
			if tst.root == nil {
				tst.root = &Node{values: make([]interface{}, 0), parent: nil}
			}
			current = tst.root
			current.parent = parent
			current.key, current.values = k, append(current.values, value)
		}
	}
}

// Query TST with given prefix of keys
func (tst *TST) Query(keys []byte) []interface{} {
	values := make([]interface{}, 0)
	current := tst.root

	for _, k := range keys {
		for current != nil && current.key != k {
			if k < current.key {
				current = current.left
			} else if k > current.key {
				current = current.right
			} else {
				current = current.middle
			}
		}
	}
	if current != nil {
		values = current.values
	}
	return values
}

//todo Remove function
