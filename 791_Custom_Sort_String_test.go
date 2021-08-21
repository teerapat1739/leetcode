package leetcode

import "testing"

func Test_customSortString(t *testing.T) {
	type args struct {
		order string
		s     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Example",
			args: args{
				order: "cba",
				s:     "abcd",
			},
			want: "dcba",
		},
		{
			name: "Example2",
			args: args{
				order: "kqep",
				s:     "pekeq",
			},
			want: "kqeep",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := customSortString(tt.args.order, tt.args.s); got != tt.want {
				t.Errorf("customSortString() = %v, want %v", got, tt.want)
			}
		})
	}
}
