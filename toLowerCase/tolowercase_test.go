package toLowerCase

import "testing"

func Test_toLowerCase(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"Example 1", "Hello", "hello"},
		{"Example 2", "here", "here"},
		{"Example 3", "LOVELY", "lovely"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toLowerCase(tt.s); got != tt.want {
				t.Errorf("toLowerCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_toLowerCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		toLowerCase("LOVELY")
	}
}

// 57254350	        20.48 ns/op
//71636089	        15.03 ns/op
//Benchmark_toLowerCase-12    	28878404	        40.14 ns/op
//Benchmark_toLowerCase-12    	12000664	        98.34 ns/op
