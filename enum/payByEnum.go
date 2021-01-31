package enum

type PayType int
const (
	Bank PayType = 0
	ApplePay PayType =1
)

func (p PayType) String() string{
	switch p {
	case Bank:
		return "bank"
	case ApplePay:
		return "apple"
	default:
		return "Unknown"
	}
}