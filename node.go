package bt

// Node represents a point in the binary tree
// The Node contains a key, its 'own value', and can point to a left or right node.
// The left node contains a key that is of a lower value than the node, whilst the right node contains
// a key of a higher value
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// New creates a new node with the given initialKey
func New(initialKey int) *Node {
	return &Node{initialKey, nil, nil}
}
