package leetcode

import "testing"

func Test_deleteAndEarn(t *testing.T) {
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
					2, 3, 3, 4,
				},
			},
			want: 6,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{
					3, 4, 2, 5, 6, 3,
				},
			},
			want: 12,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{
					1,
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteAndEarn(tt.args.nums); got != tt.want {
				t.Errorf("deleteAndEarn() = %v, want %v", got, tt.want)
			}
		})
	}
}
