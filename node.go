package bt

// Node represents a point in the binary tree.
// The Node contains a key, its 'own value', and can point to a left or right node.
// The value can be anything, including a nil pointer, this allows for sorting of objects as opposed to just ints.
// The left node contains a key that is of a lower value than the node, whilst the right node contains
// a key of a higher value.
type Node struct {
	Key   int
	Value interface{}
	Left  *Node
	Right *Node
}

// New creates a new node with the given initialKey, and initialValue.
// The value can be empty.
func New(initialKey int, initialValue interface{}) *Node {
	return &Node{initialKey, initialValue, nil, nil}
}
