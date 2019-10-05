package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	total := Sum(4, 5)
	if total != 9 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestLowcase(t *testing.T) {
	type args struct {
		name string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{name: "Badu"}, want: "badu"},
		{name: "2", args: args{name: "Kofi"}, want: "kofi"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lowcase(tt.args.name); got != tt.want {
				t.Errorf("Lowcase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpcase(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{name: "Kofi"}, want: "KOFI"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Upcase(tt.args.name); got != tt.want {
				t.Errorf("Upcase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{nums: []int{1, 1, 3, 4}}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumArray(tt.args.nums); got != tt.want {
				t.Errorf("SumArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
