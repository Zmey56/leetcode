package numUniqueEmails

import "testing"

func Test_numUniqueEmails(t *testing.T) {
	tests := []struct {
		name   string
		emails []string
		want   int
	}{
		{
			name: "Test case 1",
			emails: []string{
				"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"},
			want: 2,
		},
		{
			name: "Test case 2",
			emails: []string{
				"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numUniqueEmails(tt.emails); got != tt.want {
				t.Errorf("numUniqueEmails() = %v, want %v", got, tt.want)
			}
		})
	}

}
