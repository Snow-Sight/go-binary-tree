package bt

// Min returns the lowest value in the node or its children
func (n *Node) Min() int {
	if n.Left == nil {
		return n.Key
	}
	return n.Left.Min()
}

// Max returns the highest value in the node or its children
func (n *Node) Max() int {
	if n.Right == nil {
		return n.Key
	}

	return n.Right.Max()
}
