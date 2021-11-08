package algoexpert

import (
	"reflect"
	"testing"
)

func TestArrayOfProducts(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				[]int{
					5, 1, 4, 2,
				},
			},
			want: []int{
				8, 40, 10, 20,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayOfProducts(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayOfProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}
