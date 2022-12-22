package math

import "testing"

func TestAverage(t *testing.T) {

}

func TestAddTwo(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Sample test 1",
			args: args{num: 2},
			want: 4,
		},
		{
			name: "Sample test 2",
			args: args{num: 0},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTwo(tt.args.num); got != tt.want {
				t.Errorf("AddTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
