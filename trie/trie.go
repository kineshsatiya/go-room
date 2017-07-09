package trie

type Node struct {
	children map[byte]*Node
	values   []interface{}
}

type Trie struct {
	root *Node
}

// create new trie
func New() *Trie {
	return &Trie{root: nil}
}

// store the value for the given keys at the leaf of the trie
func (trie *Trie) Insert(keys []byte, value interface{}) {
	if trie.root == nil {
		trie.root = &Node{children: make(map[byte]*Node), values: make([]interface{}, 0)}
	}
	current := trie.root
	for _, k := range keys {
		child, contains := current.children[k]
		if !contains {
			node := &Node{children: make(map[byte]*Node), values: make([]interface{}, 0)}
			current.children[k], current = node, node
		} else {
			current = child
		}
	}
	current.values = append(current.values, value)
}

func (trie *Trie) Find(keys []byte) ([]interface{}, bool) {
	if trie.root == nil {
		return nil, false
	}
	// traverse the trie using keys
	current := trie.root
	for _, k := range keys {
		child, ok := current.children[k]
		if ok {
			current = child
		} else {
			return nil, false
		}
	}
	return current.values, true
}
