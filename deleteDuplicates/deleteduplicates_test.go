package deleteDuplicates

import "testing"

//Test for function deleteDuplicates

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "Example 1",
			input: []int{1, 1, 2},
			want:  []int{1, 2},
		},
		{
			name:  "Example 2",
			input: []int{1, 1, 2, 3, 3},
			want:  []int{1, 2, 3},
		},
		{
			name:  "Example 3",
			input: []int{0, 0, 0, 0, 0},
			want:  []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deleteDuplicates(createListNode(tt.input))
			gotSlice := createSlice(got)
			for i := range gotSlice {
				if gotSlice[i] != tt.want[i] {
					t.Errorf("%v deleteDuplicates() = %v, want %v", tt.name, gotSlice, tt.want)
				}
			}
		})
	}

}
