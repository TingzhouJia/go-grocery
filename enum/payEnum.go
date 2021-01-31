package enum
type PayStatus int
const (
	UnPay PayStatus =0
	Paid PayStatus =1
)

func (p PayStatus) String() string  {
	switch p {
	case UnPay:
		return "Unpaid"
	case Paid:
		return "Paid"
	default:
		return "Unknown"
	}
}
