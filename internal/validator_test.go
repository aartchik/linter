package internal

import (
	"testing"
)


func TestIsEnglish(t *testing.T) {
	tests := []struct {
        name     string
        input string
        wantOutput bool
    }{
		{
			name: "empty string",
			input: "",
			wantOutput: true,
		},
		{
			name: "correct string",
			input: "correct string only english",
			wantOutput: true,
		},
		{
			name: "uncorrect string",
			input: "некорректная строка",
			wantOutput: false,
		},
		{
			name: "two languages",
			input: "test тест",
			wantOutput: false,
		},
		{
			name: "strange correct string",
			input: "tEst TESDT1412312 kljsdfn___234234 ;lksdajf 2",
			wantOutput: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name , func(t *testing.T) {
			res := isEnglish(tt.input)
			if res != tt.wantOutput {
				t.Errorf("got: %v; want: %v", res, tt.wantOutput)
			}
		})
	}
}


func TestHasSpecialSymbols(t *testing.T) {
	tests := []struct {
        name     string
        input string
        wantOutput bool
		suggestedFixes string
    }{
		{
			name: "empty string",
			input: "",
			wantOutput: true,
			suggestedFixes: "",
		},
		{
			name: "correct string",
			input: "correct string",
			wantOutput: true,
			suggestedFixes: "correct string",
		},
		{
			name: "uncorrect string",
			input: "некорректная строка!!!",
			wantOutput: false,
			suggestedFixes: "некорректная строка",
		},
		{
			name: "special",
			input: "test тест!!!",
			wantOutput: false,
			suggestedFixes: "test тест",
		},
		{
			name: "strange correct string",
			input: "ljasdfhnasl длфывоалдфывоаь                     ",
			wantOutput: true,
			suggestedFixes: "ljasdfhnasl длфывоалдфывоаь",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name , func(t *testing.T) {
			res := notHasSpecialSymbols(tt.input)
			if res != tt.wantOutput {
				t.Errorf("got: %v; want: %v", res, tt.wantOutput)
			}
			fixed := toStandardSymbols(tt.input)
			if fixed != tt.suggestedFixes {
				t.Errorf("got: %v; want: %v", fixed, tt.suggestedFixes)
			}
		})
	}
}

func TestLowerCase(t *testing.T) {
	tests := []struct {
        name     string
        input string
        wantOutput bool
		suggestedFixes string
    }{
		{
			name: "empty string",
			input: "",
			wantOutput: true,
			suggestedFixes: "",
		},
		{
			name: "correct string",
			input: "correct string",
			wantOutput: true,
			suggestedFixes: "correct string",
		},
		{
			name: "uncorrect string",
			input: "Uncorrect string!!!",
			wantOutput: false,
			suggestedFixes: "uncorrect string!!!",
		},
		{
			name: "special",
			input: "Test тест!!!",
			wantOutput: false,
			suggestedFixes: "test тест!!!",
		},
		{
			name: "strange correct string",
			input: "ljasdfhnasl длфывоалдфывоаь                     ",
			wantOutput: true,
			suggestedFixes: "ljasdfhnasl длфывоалдфывоаь                     ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name , func(t *testing.T) {
			res := checkLowerCase(tt.input)
			if res != tt.wantOutput {
				t.Errorf("got: %v; want: %v", res, tt.wantOutput)
			}
			fixed := toLowerCase(tt.input)
			if fixed != tt.suggestedFixes {
				t.Errorf("got: %v; want: %v", fixed, tt.suggestedFixes)
			}
		})
	}
}

func TestContainsSesnitiveWord(t *testing.T) {
	tests := []struct {
        name     string
        input string
        wantOutput bool
    }{
		{
			name: "empty string",
			input: "",
			wantOutput: true,
		},
		{
			name: "correct string",
			input: "correct string",
			wantOutput: true,
		},
		{
			name: "uncorrect string",
			input: "password=",
			wantOutput: false,
		},
		{
			name: "uncorrect string",
			input: "token:",
			wantOutput: false,
		},
		{
			name: "uncorrect string",
			input: "apikey=",
			wantOutput: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name , func(t *testing.T) {
			res := notContainsSensitiveWordInMsg(tt.input)
			if res != tt.wantOutput {
				t.Errorf("got: %v; want: %v", res, tt.wantOutput)
			}
		})
	}
}