package bt

// Delete removes the key from the tree
func (n *Node) Delete(key int) *Node {
	// search for `key`
	if n.Key < key {
		n.Right = n.Right.Delete(key)
	} else if n.Key > key {
		n.Left = n.Left.Delete(key)
		// n.Key == `key`
	} else {
		if n.Left == nil { // just point to opposite node
			return n.Right
		} else if n.Right == nil { // just point to opposite node
			return n.Left
		}

		// if `n` has two children, you need to
		// find the next highest number that
		// should go in `n`'s position so that
		// the BST stays correct
		min := n.Right.Min()

		// we only update `n`'s key with min
		// instead of replacing n with the min
		// node so n's immediate children aren't orphaned
		n.Key = min
		n.Right = n.Right.Delete(min)
	}
	return n
}
