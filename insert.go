package bt

// Insert adds the key to the node or its children
func (n *Node) Insert(key int) {
	if n.Key < key {
		if n.Right == nil { // we found an empty spot, done!
			n.Right = &Node{Key: key}
		} else { // look right
			n.Right.Insert(key)
		}
	} else if n.Key > key {
		if n.Left == nil { // we found an empty spot, done!
			n.Left = &Node{Key: key}
		} else { // look left
			n.Left.Insert(key)
		}
	}
	// n.Key == key, don't need to do anything
}
