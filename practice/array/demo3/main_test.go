package main

import "testing"

func Test_search_2(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "hello", args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 444}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search_2(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search_2() = %v, want %v", got, tt.want)
			}
		})
	}
}
