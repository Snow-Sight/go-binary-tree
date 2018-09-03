package bt

import (
	"testing"
)

func TestMin(t *testing.T) {
	tt := []struct {
		n        *Node
		scenario string
		out      int
	}{
		{
			n: &Node{
				Key:   5,
				Left:  nil,
				Right: nil,
			},
			scenario: "Only root node",
			out:      5,
		}, {
			n: &Node{
				Key:   5,
				Left:  &Node{3, nil, nil},
				Right: nil,
			},
			scenario: "Check left node present, right node absent",
			out:      3,
		},
		{
			n: &Node{
				Key:   5,
				Left:  nil,
				Right: &Node{7, nil, nil},
			},
			scenario: "Left node absent, right node present",
			out:      5,
		},
	}

	for _, test := range tt {
		if out := test.n.Min(); out != test.out {
			t.Errorf("In case: \"%s\", got %d, expected %d", test.scenario, out, test.out)
		}
	}
}

func TestMax(t *testing.T) {
	tt := []struct {
		n        *Node
		scenario string
		out      int
	}{
		{
			n: &Node{
				Key:   5,
				Left:  nil,
				Right: nil,
			},
			scenario: "Only root node",
			out:      5,
		}, {
			n: &Node{
				Key:   5,
				Left:  &Node{3, nil, nil},
				Right: nil,
			},
			scenario: "Left node present, right node absent",
			out:      5,
		},
		{
			n: &Node{
				Key:   5,
				Left:  nil,
				Right: &Node{7, nil, nil},
			},
			scenario: "Left node absent, right node present",
			out:      7,
		},
	}

	for _, test := range tt {
		if out := test.n.Max(); out != test.out {
			t.Errorf("In case: \"%s\", got %d, expected %d", test.scenario, out, test.out)
		}
	}
}
