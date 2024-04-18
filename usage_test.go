package anthropic

import "testing"

func TestUsage_String(t *testing.T) {
	tests := []struct {
		name  string
		usage *Usage
		want  string
	}{
		{
			name:  "empty",
			usage: &Usage{},
			want:  "Input: 0, Output: 0",
		},
		{
			name:  "input",
			usage: &Usage{InputTokens: 1},
			want:  "Input: 1, Output: 0",
		},
		{
			name:  "output",
			usage: &Usage{OutputTokens: 2},
			want:  "Input: 0, Output: 2",
		},
		{
			name:  "both",
			usage: &Usage{InputTokens: 3, OutputTokens: 4},
			want:  "Input: 3, Output: 4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.usage.String(); got != tt.want {
				t.Errorf("Usage.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
