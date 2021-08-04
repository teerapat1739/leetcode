package leetcode

import "testing"

func Test_minFallingPathSum(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// {
		// 	name: "Example 1",
		// 	args: args{
		// 		matrix: [][]int{
		// 			{2, 1, 3},
		// 			{6, 5, 4},
		// 			{7, 8, 9},
		// 		},
		// 	},
		// 	want: 13,
		// },
		// {
		// 	name: "Example 2",
		// 	args: args{
		// 		matrix: [][]int{
		// 			{-19, 57},
		// 			{-40, -5},
		// 		},
		// 	},
		// 	want: -59,
		// },
		// {
		// 	name: "Example 3",
		// 	args: args{
		// 		matrix: [][]int{
		// 			{-48},
		// 		},
		// 	},
		// 	want: -48,
		// },
		{
			name: "Example 3",
			args: args{
				matrix: [][]int{
					{100, -42, -46, -41},
					{31, 97, 10, -10},
					{-58, -51, 82, 89},
					{51, 81, 69, -51},
				},
			},
			want: -48,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFallingPathSum(tt.args.matrix); got != tt.want {
				t.Errorf("minFallingPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
