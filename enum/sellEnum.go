package enum

type SellStatus int

const (
	Selling SellStatus =0
	StopSelling SellStatus =1
)

func (receiver SellStatus) String() string {
	switch receiver {
	case Selling:
		return "Selling"
	case StopSelling:
		return "Stopped"
	default:
		return "Unknown"
	}
}
