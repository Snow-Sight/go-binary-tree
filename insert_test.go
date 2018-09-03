package bt

// You only need to import one library!
import (
	"math/rand"
	"testing"
)

func TestInsert(t *testing.T) {
	tt := []struct {
		node     *Node
		insert   int
		scenario string
	}{
		{
			node: &Node{
				Key:   5,
				Left:  nil,
				Right: nil,
			},
			insert:   5,
			scenario: "Insert root value",
		},
		{
			node: &Node{
				Key:   5,
				Left:  nil,
				Right: nil,
			},
			insert:   1,
			scenario: "Insert less than root key",
		},
		{
			node: &Node{
				Key:   5,
				Left:  &Node{3, nil, nil},
				Right: nil,
			},
			insert:   7,
			scenario: "Insert greater than root key",
		},
		{
			node: &Node{
				Key:   5,
				Left:  &Node{3, nil, nil},
				Right: nil,
			},
			insert:   2,
			scenario: "Insert as child of left",
		},
		{
			node: &Node{
				Key:   5,
				Right: &Node{8, nil, nil},
				Left:  nil,
			},
			insert:   9,
			scenario: "Insert as child of right",
		},
	}

	for _, test := range tt {
		if test.node.Insert(test.insert); !test.node.Search(test.insert) {
			t.Errorf("%s, got %v, expected %v", test.scenario, false, true)
		}
	}

}

func BenchmarkInsert(b *testing.B) {
	b.Run("1 Node", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tree := New(5)

			tree.Insert(3)
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
		}
	})
}
