package algorithms

import "testing"

func TestDemoBinarySerchTree(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "BinarySearchTree",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DemoBinarySerchTree()
		})
	}
}
