package summaryRanges

import (
	"testing"
)

// Test for summaryRanges
func TestSummaryRanges(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		want    []string
		wantErr bool
	}{
		{
			name:    "Test Case 1",
			nums:    []int{0, 1, 2, 4, 5, 7},
			want:    []string{"0->2", "4->5", "7"},
			wantErr: false,
		},
		{
			name:    "Test Case 2",
			nums:    []int{0, 2, 3, 4, 6, 8, 9},
			want:    []string{"0", "2->4", "6", "8->9"},
			wantErr: false,
		},
		{
			name:    "Test Case 3",
			nums:    []int{},
			want:    []string{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := summaryRanges(tt.nums)
			if len(got) != len(tt.want) {
				t.Errorf("Expected %v but got %v", tt.want, got)
			}
			for i := 0; i < len(got); i++ {
				if got[i] != tt.want[i] {
					t.Errorf("Expected %v but got %v", tt.want, got)
				}
			}
		})
	}
}
