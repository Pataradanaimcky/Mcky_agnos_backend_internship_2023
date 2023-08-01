package tests

import (
	"Mcky_agnos_backend_internship_2023/utils"
	"testing"
)

func TestCountActionsToMakePasswordStrong(t *testing.T) {
	tests := []struct {
		input      string
		wantSteps  int
		wantStrong bool
	}{
		{"aA1", 3, false},
		{"1445D1cd", 0, true},
		{"Abcd123", 0, true},
		{"abcd123", 1, false},
		{"abcdABCD1234", 0, true},
	}

	for _, test := range tests {
		gotSteps, gotStrong := utils.CountActionsToMakePasswordStrong(test.input)
		if gotSteps != test.wantSteps || gotStrong != test.wantStrong {
			t.Errorf("CountActionsToMakePasswordStrong(%q) = (%d, %t), want (%d, %t)", test.input, gotSteps, gotStrong, test.wantSteps, test.wantStrong)
		}
	}
}
