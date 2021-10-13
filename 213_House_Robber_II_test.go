package leetcode

import "testing"

func Test_robII(t *testing.T) {
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
					1, 2, 3, 1,
				},
			},
			want: 4,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{
					2, 3, 2,
				},
			},
			want: 3,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{
					1, 2, 3,
				},
			},
			want: 3,
		},
		{
			name: "case 4",
			args: args{
				nums: []int{
					1, 2,
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := robII(tt.args.nums); got != tt.want {
				t.Errorf("robII() = %v, want %v", got, tt.want)
			}
		})
	}
}
