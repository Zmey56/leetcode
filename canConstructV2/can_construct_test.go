package canConstructV2

import "testing"

func TestCanConstruct(t *testing.T) {
	test := []struct {
		ransomNote string
		magazine   string
		expected   bool
	}{
		{
			ransomNote: "a",
			magazine:   "b",
			expected:   false,
		},
		{
			ransomNote: "aa",
			magazine:   "ab",
			expected:   false,
		},
		{
			ransomNote: "aa",
			magazine:   "aab",
			expected:   true,
		},
	}

	for _, test := range test {
		result := canConstruct(test.ransomNote, test.magazine)
		if result != test.expected {
			t.Errorf("canConstruct(%s, %s) = %t, expected %t", test.ransomNote, test.magazine, result, test.expected)
		}
	}
}

func BenchmarkCanConstruct(b *testing.B) {
	benchmarks := []struct {
		name       string
		ransomNote string
		magazine   string
	}{
		{
			name:       "Short/Success",
			ransomNote: "aa",
			magazine:   "aab",
		},
		{
			name:       "Short/Fail",
			ransomNote: "aa",
			magazine:   "ab",
		},
		{
			name:       "Medium/Success",
			ransomNote: "thequickbrownfox",
			magazine:   "thequickbrownfoxjumpsoverthelazydog",
		},
		{
			name:       "Medium/Fail",
			ransomNote: "thequickbrownfoxz",
			magazine:   "thequickbrownfoxjumpsoverthelazydog",
		},
		{
			name:       "Long/Success",
			ransomNote: "abcdefghijklmnopqrstuvwxyz",
			magazine:   "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
		},
		{
			name:       "Long/EarlyFail",
			ransomNote: "zzzzzzzzzzzzzzzzzzzzzzzzzzz",
			magazine:   "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name+"/Map", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				canConstruct(bm.ransomNote, bm.magazine)
			}
		})

		b.Run(bm.name+"/Array", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				canConstructV2(bm.ransomNote, bm.magazine)
			}
		})

		b.Run(bm.name+"/Hash", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				canConstructV3(bm.ransomNote, bm.magazine)
			}
		})

		b.Run(bm.name+"/Bit", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				canConstructV4(bm.ransomNote, bm.magazine)
			}
		})
	}
}
