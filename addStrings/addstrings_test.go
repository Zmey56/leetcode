package addStrings

import "testing"

//Tet for function addStrings

func Test_addStrings(t *testing.T) {
	tests := []struct {
		num1 string
		num2 string
		want string
	}{
		{
			num1: "11",
			num2: "123",
			want: "134",
		},
		{
			num1: "456",
			num2: "77",
			want: "533",
		},
		{
			num1: "0",
			num2: "0",
			want: "0",
		},
		{
			num1: "999",
			num2: "999",
			want: "1998",
		},
		{
			num1: "584",
			num2: "18",
			want: "602",
		},
	}
	for _, tt := range tests {
		if got := addStrings(tt.num1, tt.num2); got != tt.want {
			t.Errorf("addStrings() = %v, want %v", got, tt.want)
		}
	}
}
