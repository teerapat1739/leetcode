package slice

import "testing"

func Test_slideFirst5Digits(t *testing.T) {
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
				s: "014020993338888",
			},
			want: "0993338888",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slideFirst5Digits(tt.args.s); got != tt.want {
				t.Errorf("slideFirst5Digits() = %v, want %v", got, tt.want)
			}
		})
	}
}
