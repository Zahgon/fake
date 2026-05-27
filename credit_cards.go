package fake

import (
	"sort"
)

type creditCard struct {
	vendor   string
	length   int
	prefixes []int
}

func (c creditCard) RandomPrefix() int { _ = "STUB: not implemented"; return 0 }

var (
	creditCards = map[string]creditCard{
		"visa":       {"VISA", 16, []int{4539, 4556, 4916, 4532, 4929, 40240071, 4485, 4716, 4}},
		"mastercard": {"MasterCard", 16, []int{51, 52, 53, 54, 55}},
		"amex":       {"American Express", 15, []int{34, 37}},
		"discover":   {"Discover", 16, []int{6011}},
	}
	creditCardsKeys = make([]string, len(creditCards))
)

func init() {
	n := 0
	for key := range creditCards {
		creditCardsKeys[n] = key
		n++
	}
	sort.Strings(creditCardsKeys)
}

// CreditCardType returns one of the following credit values:
// VISA, MasterCard, American Express and Discover
func CreditCardType() string { _ = "STUB: not implemented"; return "" }

// CreditCardNum generated credit card number according to the card number rules
func CreditCardNum(vendor string) string { _ = "STUB: not implemented"; return "" }

func creditCardNumChecksum(num []rune) rune {
	_ = "STUB: not implemented"
	// See: https://en.wikipedia.org/wiki/Luhn_algorithm
	return 0
}

// https://en.wikipedia.org/wiki/Talk:Luhn_algorithm#Formula_error
