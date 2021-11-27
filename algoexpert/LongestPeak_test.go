package algoexpert

import "testing"

func TestLongestPeak(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				array: []int{
					1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3,
				},
			},
			want: 6,
		},
		{
			name: "case 2",
			args: args{
				array: []int{
					1, 3, 2,
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestPeak(tt.args.array); got != tt.want {
				t.Errorf("LongestPeak() = %v, want %v", got, tt.want)
			}
		})
	}
}
