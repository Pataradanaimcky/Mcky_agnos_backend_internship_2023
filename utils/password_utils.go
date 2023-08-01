package utils

import "regexp"

const (
	minPasswordLength = 6
	maxPasswordLength = 20
)

func CountActionsToMakePasswordStrong(password string) (int, bool) {
	actions := 0

	if len(password) < minPasswordLength {
		actions += minPasswordLength - len(password)
	} else if len(password) > maxPasswordLength {
		actions += len(password) - maxPasswordLength
	}

	if !hasLowerCase(password) {
		actions++
	}

	if !hasUpperCase(password) {
		actions++
	}

	if !hasDigit(password) {
		actions++
	}

	if !hasDotOrExclamation(password) {
		actions++
	}

	actions += countRepeatingCharactersActions(password) - 1

	isStrong := actions == 0

	return actions, isStrong
}

func hasLowerCase(password string) bool {
	return regexp.MustCompile("[a-z]").MatchString(password)
}

func hasUpperCase(password string) bool {
	return regexp.MustCompile("[A-Z]").MatchString(password)
}

func hasDigit(password string) bool {
	return regexp.MustCompile("[0-9]").MatchString(password)
}

func hasDotOrExclamation(password string) bool {
	return regexp.MustCompile(`[.!]`).MatchString(password)
}

func countRepeatingCharactersActions(password string) int {
	actions := 0
	repeatingPattern := regexp.MustCompile(`(.)\\1\\1`)
	for repeatingPattern.MatchString(password) {
		actions++
		password = repeatingPattern.ReplaceAllString(password, "$1$1")
	}
	return actions
}
