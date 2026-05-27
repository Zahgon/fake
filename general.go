package fake

var lowerLetters = []rune("abcdefghijklmnopqrstuvwxyz")
var upperLetters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numeric = []rune("0123456789")
var specialChars = []rune(`!'@#$%^&*()_+-=[]{};:",./?`)
var hexDigits = []rune("0123456789abcdef")

func text(atLeast, atMost int, allowLower, allowUpper, allowNumeric, allowSpecial bool) string {
	_ = "STUB: not implemented"
	return ""
}

// Password generates password with the length from atLeast to atMOst charachers,
// allow* parameters specify whether corresponding symbols can be used
func Password(atLeast, atMost int, allowUpper, allowNumeric, allowSpecial bool) string {
	_ = "STUB: not implemented"
	return ""
}

// SimplePassword is a wrapper around Password,
// it generates password with the length from 6 to 12 symbols, with upper characters and numeric symbols allowed
func SimplePassword() string { _ = "STUB: not implemented"; return "" }

// Color generates color name
func Color() string { _ = "STUB: not implemented"; return "" }

// DigitsN returns n digits as a string
func DigitsN(n int) string { _ = "STUB: not implemented"; return "" }

// Digits returns from 1 to 5 digits as a string
func Digits() string { _ = "STUB: not implemented"; return "" }

func hexDigitsStr(n int) string { _ = "STUB: not implemented"; return "" }

// HexColor generates hex color name
func HexColor() string { _ = "STUB: not implemented"; return "" }

// HexColorShort generates short hex color name
func HexColorShort() string { _ = "STUB: not implemented"; return "" }
