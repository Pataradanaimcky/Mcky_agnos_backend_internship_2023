package services_test

import (
	"testing"

	"strong_password_backend/services"
)

func TestPasswordService_GetMinimumStepsToStrongPassword(t *testing.T) {
	tests := []struct {
		name       string
		password   string
		numOfSteps int
	}{
		{
			name:       "Short password",
			password:   "abc",
			numOfSteps: 3,
		},
		{
			name:       "Long password",
			password:   "abcdefghijabcdefghijabcdefghijabcdefghijabcdefghij",
			numOfSteps: 1,
		},
		{
			name:       "No lowercase letter",
			password:   "ABC123",
			numOfSteps: 1,
		},
		{
			name:       "No uppercase letter",
			password:   "abc123",
			numOfSteps: 1,
		},
		{
			name:       "No digit",
			password:   "abcABC",
			numOfSteps: 1,
		},
		{
			name:       "Repeating characters",
			password:   "aaa123",
			numOfSteps: 1,
		},
		{
			name:       "Valid strong password",
			password:   "Abc123!",
			numOfSteps: 0,
		},
	}

	passwordService := services.NewPasswordService()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			steps := passwordService.GetMinimumStepsToStrongPassword(test.password)
			if steps != test.numOfSteps {
				t.Errorf("Expected %d steps, got %d", test.numOfSteps, steps)
			}
		})
	}
}
