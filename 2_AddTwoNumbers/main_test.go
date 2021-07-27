package main

import (
	"reflect"
	"testing"
)

func Test_helper(t *testing.T) {
	type args struct {
		left  *ListNode
		right *ListNode
		carry int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := helper(tt.args.left, tt.args.right, tt.args.carry); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("helper() = %v, want %v", got, tt.want)
			}
		})
	}
}
