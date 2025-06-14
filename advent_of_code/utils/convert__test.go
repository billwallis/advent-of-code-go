package utils_test

import (
	"testing"

	"github.com/billwallis/advent-of-code-go/advent_of_code/utils"
)

func Test_ToInt(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  int
	}{
		{name: "convert str", input: "42", want: 42},
		{name: "convert int", input: 1234, want: 1234},
		{name: "convert ''", input: "", want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := utils.ToInt(test.input)
			if got != test.want {
				t.Errorf("ToInt(%v) = %v, want %v", test.input, got, test.want)
			}
		})
	}
}
