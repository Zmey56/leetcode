package toGoatLatin

import "testing"

func Test_toGoatLatin(t *testing.T) {
	tests := []struct {
		name     string
		sentence string
		want     string
	}{
		{
			name:     "Example 1",
			sentence: "I speak Goat Latin",
			want:     "Imaa peaksmaaa oatGmaaaa atinLmaaaaa",
		},
		{
			name:     "Example 2",
			sentence: "The quick brown fox jumped over the lazy dog",
			want:     "heTmaa uickqmaaa ownbrmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa",
		},
		{
			name:     "Example 3",
			sentence: "Each word consists of lowercase and uppercase letters only",
			want:     "Eachmaa ordwmaaa onsistscmaaaa ofmaaaaa owercaselmaaaaaa andmaaaaaaa ppercaseumaaaaaaaa etterslmaaaaaaaaaa onlymaaaaaaaaaaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toGoatLatin(tt.sentence); got != tt.want {
				t.Errorf("toGoatLatin() = %v, want %v", got, tt.want)
			}
		})
	}
}
