package services

import (
	"unicode"
)

type PasswordService struct{}

func NewPasswordService() *PasswordService {
	return &PasswordService{}
}

func (s *PasswordService) GetMinimumStepsToStrongPassword(password string) int {
	// Implement the logic to determine the minimum number of actions to make the password strong

	// Criteria of a strong password
	lengthCriteriaMet := len(password) >= 6 && len(password) < 20
	lowercaseCriteriaMet := false
	uppercaseCriteriaMet := false
	digitCriteriaMet := false
	repeatingCharsCriteriaMet := true

	lastChar := ' '
	repeatingCharCount := 1

	for _, char := range password {
		if unicode.IsLower(char) {
			lowercaseCriteriaMet = true
		} else if unicode.IsUpper(char) {
			uppercaseCriteriaMet = true
		} else if unicode.IsDigit(char) {
			digitCriteriaMet = true
		}

		if char == lastChar {
			repeatingCharCount++
			if repeatingCharCount >= 3 {
				repeatingCharsCriteriaMet = false
				break
			}
		} else {
			repeatingCharCount = 1
		}

		lastChar = char
	}

	// Calculate the minimum number of actions required to make the password strong
	steps := 0

	if !lengthCriteriaMet {
		steps++
	}

	if !lowercaseCriteriaMet {
		steps++
	}

	if !uppercaseCriteriaMet {
		steps++
	}

	if !digitCriteriaMet {
		steps++
	}

	if !repeatingCharsCriteriaMet {
		steps++
	}

	return steps
}
