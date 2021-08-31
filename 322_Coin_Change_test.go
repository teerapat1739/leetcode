package leetcode

import "testing"

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				coins:  []int{1, 2},
				amount: 5,
			},
			want: 3,
		},
		// {
		// 	name: "case 2",
		// 	args: args{
		// 		coins:  []int{1, 2, 5},
		// 		amount: 11,
		// 	},
		// 	want: 3,
		// },
		// {
		// 	name: "case 3",
		// 	args: args{
		// 		coins:  []int{2},
		// 		amount: 3,
		// 	},
		// 	want: -1,
		// },
		// {
		// 	name: "case 4",
		// 	args: args{
		// 		coins:  []int{1},
		// 		amount: 0,
		// 	},
		// 	want: 0,
		// },
		// {
		// 	name: "case 5",
		// 	args: args{
		// 		coins:  []int{1, 2147483647},
		// 		amount: 2,
		// 	},
		// 	want: 2,
		// },
		// {
		// 	name: "case 6",
		// 	args: args{
		// 		coins:  []int{1, 2, 5},
		// 		amount: 10,
		// 	},
		// 	want: 2,
		// },
		// {
		// 	name: "case 7",
		// 	args: args{
		// 		coins:  []int{186, 419, 83, 408},
		// 		amount: 6249,
		// 	},
		// 	want: 20,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange2(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
