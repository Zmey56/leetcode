package reverseOnlyLetters

import "testing"

func Test_reverseOnlyLetters(t *testing.T) {
	tests := []struct {
		name string
		S    string
		want string
	}{
		{
			name: "Test case 1",
			S:    "ab-cd",
			want: "dc-ba",
		},
		{
			name: "Test case 2",
			S:    "a-bC-dEf-ghIj",
			want: "j-Ih-gfE-dCba",
		},
		{
			name: "Test case 3",
			S:    "Test1ng-Leet=code-Q!",
			want: "Qedo1ct-eeLg=ntse-T!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseOnlyLetters(tt.S); got != tt.want {
				t.Errorf("reverseOnlyLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
