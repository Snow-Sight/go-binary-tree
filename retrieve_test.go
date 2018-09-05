package bt

import (
	"testing"
)

func TestRetrieve(t *testing.T) {
	tt := []struct {
		scenario string
		node     *Node
		retrieve int
		outFound bool
		outValue interface{}
	}{
		{
			scenario: "Retrieve nil value of stored key",
			node: &Node{
				Key:   1,
				Value: nil,
				Left:  &Node{},
				Right: &Node{},
			},
			retrieve: 1,
			outFound: true,
			outValue: nil,
		},
		{
			scenario: "Retrieve unstored key",
			node: &Node{
				Key:   5,
				Value: "something",
				Left:  &Node{},
				Right: &Node{},
			},
			retrieve: 1,
			outFound: false,
			outValue: nil,
		},
		{
			scenario: "Retrieve stored left child",
			node: &Node{
				Key:   5,
				Value: "something",
				Left:  &Node{3, "left child", nil, nil},
				Right: &Node{},
			},
			retrieve: 3,
			outFound: true,
			outValue: "left child",
		},
		{
			scenario: "Retrieve stored right child",
			node: &Node{
				Key:   5,
				Value: "something",
				Right: &Node{7, "right child", nil, nil},
				Left:  &Node{},
			},
			retrieve: 7,
			outFound: true,
			outValue: "right child",
		},
	}

	for _, test := range tt {
		boolOut, valOut := test.node.Retrieve(test.retrieve)

		if boolOut != test.outFound {
			t.Errorf("In case \"%s\", expected found to be %t but recieved %t", test.scenario, test.outFound, boolOut)
		}

		if valOut != test.outValue {
			t.Errorf("In case \"%s\", expected value to be %v but recieved %v", test.scenario, test.outValue, valOut)
		}
	}
}
