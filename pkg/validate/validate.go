package validate

import "regexp"

var emailRegex = regexp.MustCompile(`^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`)

func IsEmail(s string) bool {
	return emailRegex.MatchString(s)
}
