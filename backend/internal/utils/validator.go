package utils

import (
	"regexp"
	"strings"
	"unicode"
)

var (
	regexUsername = regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	regexPhone    = regexp.MustCompile(`^0\d{9}$`)
	regexEmail    = regexp.MustCompile(`^[\w\.-]+@[\w\.-]+\.\w+$`)
)

func ValidateUsername(s string) bool {
	if !regexUsername.MatchString(s) {
		return false
	}
	hasLetter, hasDigit := false, false
	for _, r := range s {
		if unicode.IsLetter(r) {
			hasLetter = true
		}
		if unicode.IsDigit(r) {
			hasDigit = true
		}
	}
	return hasLetter && hasDigit
}

func ValidatePassword(s string) bool {
	if len(s) < 8 || len(s) > 20 {
		return false
	}
	var hasUpper, hasLower, hasDigit, hasSpecial bool
	for _, r := range s {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsDigit(r):
			hasDigit = true
		case !unicode.IsLetter(r) && !unicode.IsDigit(r):
			hasSpecial = true
		}
	}
	return hasUpper && hasLower && hasDigit && hasSpecial
}

func ValidatePhone(s string) bool {
	return regexPhone.MatchString(s)
}

func ValidateEmail(s string) bool {
	return regexEmail.MatchString(s)
}

func MakeSlug(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	s = strings.ReplaceAll(s, " ", "-")
	repl := map[string]string{
		"à": "a", "á": "a", "ạ": "a", "ả": "a", "ã": "a",
		"â": "a", "ầ": "a", "ấ": "a", "ậ": "a", "ẩ": "a", "ẫ": "a",
		"ă": "a", "ằ": "a", "ắ": "a", "ặ": "a", "ẳ": "a", "ẵ": "a",
		"è": "e", "é": "e", "ẹ": "e", "ẻ": "e", "ẽ": "e",
		"ê": "e", "ề": "e", "ế": "e", "ệ": "e", "ể": "e", "ễ": "e",
		"ì": "i", "í": "i", "ị": "i", "ỉ": "i", "ĩ": "i",
		"ò": "o", "ó": "o", "ọ": "o", "ỏ": "o", "õ": "o",
		"ô": "o", "ồ": "o", "ố": "o", "ộ": "o", "ổ": "o", "ỗ": "o",
		"ơ": "o", "ờ": "o", "ớ": "o", "ợ": "o", "ở": "o", "ỡ": "o",
		"ù": "u", "ú": "u", "ụ": "u", "ủ": "u", "ũ": "u",
		"ư": "u", "ừ": "u", "ứ": "u", "ự": "u", "ử": "u", "ữ": "u",
		"ỳ": "y", "ý": "y", "ỵ": "y", "ỷ": "y", "ỹ": "y",
		"đ": "d",
	}
	for k, v := range repl {
		s = strings.ReplaceAll(s, k, v)
	}
	var b strings.Builder
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			b.WriteRune(r)
		}
	}
	return b.String()
}
