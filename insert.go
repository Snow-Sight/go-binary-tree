package bt

// Insert adds the key to the node or its children
func (n *Node) Insert(key int, value interface{}) {
	if n.Key < key {
		if n.Right == nil { // we found an empty spot, done!
			n.Right = &Node{Key: key, Value: value}
		} else { // look right
			n.Right.Insert(key, value)
		}
	} else if n.Key > key {
		if n.Left == nil { // we found an empty spot, done!
			n.Left = &Node{Key: key, Value: value}
		} else { // look left
			n.Left.Insert(key, value)
		}
	}
	// n.Key == key, don't need to do anything
}
