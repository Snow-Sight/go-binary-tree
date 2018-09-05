package bt

// Retrieve returns a bool as to whether the key exists, and the corresponding value of the key.
// The value will be nil if the key is not present.
func (n *Node) Retrieve(key int) (bool, interface{}) {
	// This is our base case. If n == nil, `key`
	// doesn't exist in our binary search tree.
	if n == nil {
		return false, nil
	}

	if n.Key < key { // move right
		return n.Right.Retrieve(key)
	} else if n.Key > key { // move left
		return n.Left.Retrieve(key)
	}

	// n.Key == key, we found it!
	return true, n.Value
}
