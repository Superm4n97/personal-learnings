package main

import "testing"

//func TestAdd2(t *testing.T) {
//	if Add2(2) != 4 {
//		t.Error("Incorrect result given")
//	}
//}
//
//func TestAdd21(t *testing.T) {
//	type args struct {
//		num int
//	}
//	tests := []struct {
//		name string
//		args args
//		want int
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := Add2(tt.args.num); got != tt.want {
//				t.Errorf("Add2() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestAdd2(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add2(tt.args.num); got != tt.want {
				t.Errorf("Add2() = %v, want %v", got, tt.want)
			}
		})
	}
}
