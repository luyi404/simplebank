package util

const (
	USD = "USD"
	CAD = "CAD"
	EUR = "EUR"
)

func IsSupportCurrency(currency string) bool {
	switch currency {
	case USD, CAD, EUR:
		return true
	default:
		return false
	}
}
