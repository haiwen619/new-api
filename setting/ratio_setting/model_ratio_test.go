package ratio_setting

import "testing"

func TestFormatMatchingModelNameNormalizeGeminiImageFlow(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "gemini-2.5-flash-image-flow",
			in:   "gemini-2.5-flash-image-Flow",
			want: "gemini-2.5-flash-image",
		},
		{
			name: "gemini-3.0-pro-image-flow",
			in:   "gemini-3.0-pro-image-Flow",
			want: "gemini-3.0-pro-image",
		},
		{
			name: "gemini-3.1-flash-image-flow",
			in:   "gemini-3.1-flash-image-Flow",
			want: "gemini-3.1-flash-image",
		},
		{
			name: "non-gemini-flow-kept",
			in:   "gpt-4o-Flow",
			want: "gpt-4o-Flow",
		},
		{
			name: "gemini-non-image-flow-kept",
			in:   "gemini-3.1-flash-Flow",
			want: "gemini-3.1-flash-Flow",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := FormatMatchingModelName(tt.in)
			if got != tt.want {
				t.Fatalf("FormatMatchingModelName(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

