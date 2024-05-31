package numJewelsInStones

import "testing"

func Test_numJewelsInStones(t *testing.T) {
	tests := []struct {
		name   string
		jewels string
		stones string
		want   int
	}{
		{"Example 1", "aA", "aAAbbbb", 3},
		{"Example 2", "z", "ZZ", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numJewelsInStones(tt.jewels, tt.stones); got != tt.want {
				t.Errorf("numJewelsInStones() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numJewelsInStonesVer2(t *testing.T) {
	tests := []struct {
		name   string
		jewels string
		stones string
		want   int
	}{
		{"Example 1", "aA", "aAAbbbb", 3},
		{"Example 2", "z", "ZZ", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numJewelsInStonesVer2(tt.jewels, tt.stones); got != tt.want {
				t.Errorf("numJewelsInStones() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Banchmark_numJewelsInStones(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numJewelsInStones("aA", "aAAbbbb")
	}
}

func Banchmark_numJewelsInStonesVer2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numJewelsInStonesVer2("aA", "aAAbbbb")
	}
}
