package leetcode

import "testing"

func Test_syracuse(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example",
			args: args{
				n: 5,
			},
			want: "5 16 8 4 2 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := syracuse(tt.args.n); got != tt.want {
				t.Errorf("syracuse() = %v, want %v", got, tt.want)
			}
		})
	}
}
