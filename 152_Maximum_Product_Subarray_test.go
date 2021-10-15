package leetcode

import "testing"

func Test_maxProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{
					2, 3, -2, 4,
				},
			},
			want: 6,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{
					-1, 2, 3, -4,
				},
			},
			want: 24,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{
					-2, 0, -1,
				},
			},
			want: 0,
		},
		{
			name: "case 4",
			args: args{
				nums: []int{
					0, 2,
				},
			},
			want: 2,
		},
		{
			name: "case 5",
			args: args{
				nums: []int{
					3, -1, 4,
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
