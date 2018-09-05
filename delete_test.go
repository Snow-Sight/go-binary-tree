package bt

// You only need to import one library!
import (
	"math/rand"
	"testing"
)

func TestDelete(t *testing.T) {
	tt := []struct {
		node     *Node
		delete   int
		scenario string
	}{
		{
			node: &Node{
				Key:   5,
				Value: nil,
				Left:  &Node{1, nil, nil, nil},
				Right: nil,
			},
			delete:   1,
			scenario: "Delete left child",
		},
		{
			node: &Node{
				Key:   5,
				Value: nil,
				Left:  nil,
				Right: &Node{7, nil, nil, nil},
			},
			delete:   7,
			scenario: "Delete right child",
		},
		{
			node: &Node{
				Key:   5,
				Value: nil,
				Left:  nil,
				Right: &Node{7, nil, nil, &Node{9, nil, nil, nil}},
			},
			delete:   7,
			scenario: "Delete node with right child",
		},
		{
			node: &Node{
				Key:   5,
				Value: nil,
				Left:  nil,
				Right: &Node{7, nil, &Node{6, nil, nil, nil}, nil},
			},
			delete:   7,
			scenario: "Delete node with left child",
		},
		{
			node: &Node{
				Key:   5,
				Value: nil,
				Left:  nil,
				Right: &Node{7, nil, &Node{6, nil, nil, nil}, &Node{8, nil, nil, nil}},
			},
			delete:   7,
			scenario: "Delete node with left and right child",
		},
	}

	for _, test := range tt {
		if test.node.Delete(test.delete); test.node.Search(test.delete) {
			t.Errorf("%s, got %v, expected %v", test.scenario, true, false)
		}
	}

}

func BenchmarkDelete(b *testing.B) {
	b.Run("1 Node", func(b *testing.B) {
		tree := New(5, nil)

		tree.Insert(3, nil)

		for i := 0; i < b.N; i++ {
			t := *tree

			t.Delete(3)
		}
	})

	b.Run("10 Nodes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// generate root
			tree := New(rand.Intn(10), nil)

			// Add 8 random nodes to the tree
			for i := 0; i <= 8; i++ {
				tree.Insert(rand.Intn(10), nil)
			}

			val := rand.Intn(10)

			tree.Insert(val, nil)

			tree.Delete(val)
		}
	})

	b.Run("50 Nodes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// generate root
			tree := New(rand.Intn(50), nil)

			// Add 48 random nodes to the tree
			for i := 0; i <= 48; i++ {
				tree.Insert(rand.Intn(50), nil)
			}

			val := rand.Intn(50)

			tree.Insert(val, nil)

			tree.Delete(val)
		}
	})

	b.Run("500 Nodes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// generate root
			tree := New(rand.Intn(500), nil)

			// Add 498 random nodes to the tree
			for i := 0; i <= 498; i++ {
				tree.Insert(rand.Intn(500), nil)
			}

			val := rand.Intn(500)

			tree.Insert(val, nil)

			tree.Delete(val)
		}
	})
}
