package leetcode

import "testing"

func Test_rob(t *testing.T) {
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
					2, 7, 9, 3, 1,
				},
			},
			want: 12,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{
					9, 1, 9, 100, 1, 4,
				},
			},
			want: 113,
		},
		{
			name: "case 4",
			args: args{
				nums: []int{
					2, 2, 2, 9, 1, 1, 1, 1, 9, 1, 7, 1,
				},
			},
			want: 28,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
