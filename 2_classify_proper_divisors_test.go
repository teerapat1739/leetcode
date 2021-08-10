package leetcode

import "testing"

func Test_classifyProperdisvors(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				number: 15,
			},
			want: "deficient",
		},
		{
			name: "test2",
			args: args{
				number: 18,
			},
			want: "abundant",
		},
		{
			name: "test3",
			args: args{
				number: 6,
			},
			want: "perfect",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := classifyProperdisvors(tt.args.number); got != tt.want {
				t.Errorf("classifyProperdisvors() = %v, want %v", got, tt.want)
			}
		})
	}
}
