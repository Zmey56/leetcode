package summaryRangesV2

import (
	"reflect"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	testCases := []struct {
		nums []int
		want []string
	}{
		{[]int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}},
		{[]int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}},
	}

	for _, tc := range testCases {
		res := summaryRanges(tc.nums)
		if len(res) != len(tc.want) {
			t.Errorf("nums: %v, want: %v, got: %v", tc.nums, tc.want, res)
		}
		if reflect.DeepEqual(res, tc.want) == false {
			t.Errorf("nums: %v, want: %v, got: %v", tc.nums, tc.want, res)
		}
	}
}
