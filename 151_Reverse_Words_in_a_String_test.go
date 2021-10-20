package leetcode

import "testing"

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				s: "the sky is blue",
			},
			want: "blue is sky the",
		},
		{
			name: "case 2",
			args: args{
				s: "  hello world  ",
			},
			want: "world hello",
		},
		{
			name: "case 3",
			args: args{
				s: "  Bob    Loves  Alice   ",
			},
			want: "Alice Loves Bob",
		},
		{
			name: "case 4",
			args: args{
				s: "Alice does not even like bob",
			},
			want: "bob like even not does Alice",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
