package leetcode

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Example 1",
			args: args{
				strs: []string{"flower", "flow", "flight"},
			},
			want: "fl",
		},
		{
			name: "Case: when empty array.",
			args: args{
				strs: []string{},
			},
			want: "",
		},
		{
			name: "Case: when len array equal one.",
			args: args{
				strs: []string{},
			},
			want: "",
		},
		{
			name: "Case: when have not common prefix.",
			args: args{
				strs: []string{"aaxx", "bbb", "wwwwww"},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
