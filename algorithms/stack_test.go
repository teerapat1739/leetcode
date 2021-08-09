package algorithms

import "testing"

func TestDemo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Example",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Demo()
		})
	}
}
