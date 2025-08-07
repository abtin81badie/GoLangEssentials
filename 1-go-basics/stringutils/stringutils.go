package stringutils

import "strings"

// ToUpperCase converts a string to uppercase.
func ToUpperCase(s string) string {
	return strings.ToUpper(s)
}

// ToLowerCase converts a string to lowercase.
func ToLowerCase(s string) string {
	return strings.ToLower(s)
}

// Contains checks if a substring exists within a string.
func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// TrimSpaces removes leading and trailing spaces from a string.
func TrimSpaces(s string) string {
	return strings.TrimSpace(s)
}

// ReplaceSubstring replaces all occurrences of old with new in a string.
func ReplaceSubstring(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// SplitString splits a string into a slice of substrings based on a delimiter .
func SplitString(s, delimiter string) []string {
	return strings.Split(s, delimiter)
}
