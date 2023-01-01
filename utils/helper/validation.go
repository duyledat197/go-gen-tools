package helper

import (
	"crypto/sha256"
	"regexp"

	validation "github.com/asaskevich/govalidator"
)

// IsEmail check valid Email
func IsEmail(email string) bool {
	valid := validation.IsEmail(email)
	return valid
}

// IsPassword check valid password  8 - 32
func IsPassword(password string) bool {
	valid, _ := regexp.MatchString("^.{6,20}$", password)
	return valid
}

// IsURL to check valid url
func IsURL(url string) bool {
	valid := validation.IsURL(url)
	return valid
}

// IsHexDigit ...
func IsHexDigit(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	if r >= 'A' && r <= 'F' {
		return true
	}
	if r >= 'a' && r <= 'f' {
		return true
	}
	return false
}

const sha256HexSize = sha256.Size * 2

// IsSHA256 to check sha256
func IsSHA256(s string) bool {
	if len(s) != sha256HexSize {
		return false
	}
	for _, c := range s {
		if !IsHexDigit(c) {
			return false
		}
	}
	return true
}
