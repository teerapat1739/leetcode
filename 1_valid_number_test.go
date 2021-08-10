package leetcode

import "testing"

func Test_validSameNumber(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{a: 9876543, b: 3467985},
			want: "Valid",
		},
		{
			name: "test2",
			args: args{a: 9876543, b: 7865439},
			want: "Invalid",
		},
		{
			name: "test1",
			args: args{a: 9876543, b: 3467985},
			want: "Valid",
		},
		{
			name: "test1",
			args: args{a: 9876543, b: 3456789},
			want: "Invalid",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validSameNumber(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("validSameNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
