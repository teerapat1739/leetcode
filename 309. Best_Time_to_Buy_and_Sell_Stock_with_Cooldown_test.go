package leetcode

import "testing"

func Test_maxProfitWithCoolDown(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				prices: []int{
					1, 2, 3, 0, 2,
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitWithCoolDown(tt.args.prices); got != tt.want {
				t.Errorf("maxProfitWithCoolDown() = %v, want %v", got, tt.want)
			}
		})
	}
}
