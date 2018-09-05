package bt

import (
	"testing"
)

func TestNew(t *testing.T) {
	tt := []struct {
		initialKey int
		scenario   string
		key        int
	}{
		{
			initialKey: 5,
			scenario:   "Initial key is positive integer",
			key:        5,
		},
		{
			initialKey: 0,
			scenario:   "Initial key is 0",
			key:        0,
		},
		{
			initialKey: -1,
			scenario:   "Initial key is negative integer",
			key:        -1,
		},
	}

	for _, test := range tt {
		if out := New(test.initialKey, nil); out.Key != test.key {
			t.Errorf("In case \"%s\", expected %d got %d", test.scenario, test.key, out.Key)
		}
	}
}
