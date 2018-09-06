package bt

type KV struct {
	Key   int
	Value interface{}
}

func (n *Node) RangeLToR() []KV {
	s := make([]KV, 0)

	if n.Left != nil {
		s = append(s, n.RangeLToR()...)
	}

	s = append(s, KV{
		Key:   n.Key,
		Value: n.Value,
	})

	if n.Right != nil {
		s = append(s, n.RangeLToR()...)
	}

	return s
}
