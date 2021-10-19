package algoexpert

import "testing"

func TestIsValidSubsequence(t *testing.T) {
	type args struct {
		array    []int
		sequence []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				array:    []int{5, 1, 22, 25, 6, -1, 8, 10},
				sequence: []int{1, 6, -1, 10},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidSubsequence(tt.args.array, tt.args.sequence); got != tt.want {
				t.Errorf("IsValidSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
