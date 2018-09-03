package bt

// You only need to import one library!
import (
	"math/rand"
	"testing"
)

func TestSearch(t *testing.T) {
	tt := []struct {
		node     *Node
		find     int
		scenario string
		output   bool
	}{
		{
			node: &Node{
				Key:   5,
				Left:  nil,
				Right: nil,
			},
			find:     5,
			scenario: "Finding only a root key",
			output:   true,
		},
		{
			node: &Node{
				Key:   5,
				Left:  nil,
				Right: nil,
			},
			find:     1,
			scenario: "Finding only a missing root key",
			output:   false,
		},
		{
			node: &Node{
				Key:   5,
				Left:  &Node{3, nil, nil},
				Right: nil,
			},
			find:     3,
			scenario: "Finding a child of the tree",
			output:   true,
		},
		{
			node: &Node{
				Key:   5,
				Left:  &Node{3, nil, nil},
				Right: nil,
			},
			find:     2,
			scenario: "Finding a missing child of the tree",
			output:   false,
		},
	}

	for _, test := range tt {
		if res := test.node.Search(test.find); res != test.output {
			t.Errorf("%s, got %v, expected %v", test.scenario, res, test.output)
		}
	}

}

func BenchmarkSearch(b *testing.B) {
	b.Run("1 Node", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tree := New(5)

			tree.Search(5)
		}
	})

	b.Run("10 Nodes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// generate root
			tree := New(rand.Intn(10))

			// Add 8 random nodes to the tree
			for i := 0; i <= 8; i++ {
				tree.Insert(rand.Intn(10))
			}

			val := rand.Intn(10)

			tree.Insert(val)

			tree.Search(val)
		}
	})

	b.Run("50 Nodes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// generate root
			tree := New(rand.Intn(50))

			// Add 48 random nodes to the tree
			for i := 0; i <= 48; i++ {
				tree.Insert(rand.Intn(50))
			}

			val := rand.Intn(50)

			tree.Insert(val)

			tree.Search(val)
		}
	})

	b.Run("500 Nodes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// generate root
			tree := New(rand.Intn(500))

			// Add 498 random nodes to the tree
			for i := 0; i <= 498; i++ {
				tree.Insert(rand.Intn(500))
			}

			val := rand.Intn(500)

			tree.Insert(val)

			tree.Search(val)
		}
	})
}
