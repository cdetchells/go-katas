package textSearch

import "testing"

func TestTextSearch(t *testing.T) {
	textToSearch := "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!"

	tests := []struct {
		name           string
		textToSearch   string
		subtext        string
		expectedOutput string
	}{
		{
			name:           "if textToSearch and subtext are blank expect an empty output",
			textToSearch:   "",
			subtext:        "",
			expectedOutput: "",
		},
		{
			name:           "if textToSearch is blank expect an empty output",
			textToSearch:   "",
			subtext:        "Peter",
			expectedOutput: "",
		},
		{
			name:           "if subtext is blank expect an empty output",
			textToSearch:   textToSearch,
			subtext:        "",
			expectedOutput: "",
		},
		{
			name:           "if subtext is Peter expect output to match 1,20,75",
			textToSearch:   textToSearch,
			subtext:        "Peter",
			expectedOutput: "1,20,75",
		},
		{
			name:           "if subtext is peter expect output to match 1,20,75",
			textToSearch:   textToSearch,
			subtext:        "peter",
			expectedOutput: "1,20,75",
		},
		{
			name:           "if subtext is pick expect output to match 30, 58",
			textToSearch:   textToSearch,
			subtext:        "pick",
			expectedOutput: "30,58",
		},
		{
			name:           "if subtext is pick expect output to match 30, 58",
			textToSearch:   textToSearch,
			subtext:        "pi",
			expectedOutput: "30,37,43,51,58",
		},
		{
			name:           "if subtext is z found expect empty output",
			textToSearch:   textToSearch,
			subtext:        "z",
			expectedOutput: "",
		},
		{
			name:           "if subtext is Peterz found expect empty output",
			textToSearch:   textToSearch,
			subtext:        "Peterz",
			expectedOutput: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := textSearch(tt.textToSearch, tt.subtext)
			if output != tt.expectedOutput {
				t.Errorf("textSearch()\ngot :%#v\nwant: %#v", output, tt.expectedOutput)
			}
		})
	}
}
