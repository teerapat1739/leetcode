package algorithms

import "testing"

func TestDemoQueue(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "TestQueue",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DemoQueue()
		})
	}
}
