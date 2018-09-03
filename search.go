package bt

// Search returns a boolean as to whether the key is contained in the node or it's children
func (n *Node) Search(key int) bool {
	// This is our base case. If n == nil, `key`
	// doesn't exist in our binary search tree.
	if n == nil {
		return false
	}

	if n.Key < key { // move right
		return n.Right.Search(key)
	} else if n.Key > key { // move left
		return n.Left.Search(key)
	}

	// n.Key == key, we found it!
	return true
}
