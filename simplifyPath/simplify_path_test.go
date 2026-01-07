package simplifyPath

import "testing"

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "Example 1",
			path: "/home/",
			want: "/home",
		},
		{
			name: "Example 2",
			path: "/../",
			want: "/",
		},
		{
			name: "Example 3",
			path: "/home//foo/",
			want: "/home/foo",
		},
		{
			name: "Example 4",
			path: "/a/./b/../../c/",
			want: "/c",
		},
		{
			name: "Example 5",
			path: "/a/../../b/../c//.//",
			want: "/c",
		},
		{
			name: "Example 6",
			path: "/a//b////c/d//././/..",
			want: "/a/b/c",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
