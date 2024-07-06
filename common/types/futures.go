package types

type FuturesPosition struct {
	Trader              string     `json:"trader"`
	Platform            string     `json:"platform"`
	Market              string     `json:"market"`
	CollateralToken     string     `json:"collateral_token"`
	EntryPrice          float64    `json:"entry_price"`
	CurrentPrice        float64    `json:"current_price"`
	IsolatedMarginShare float64    `json:"isolated_margin_share"`
	CrossMarginShare    float64    `json:"cross_margin_share"`
	Status              Status     `json:"status"`
	Direction           Direction  `json:"direction"`
	MarginType          MarginType `json:"margin_type"`
}
type Direction string

const (
	Long  Direction = "long"
	Short Direction = "short"
)

func (d Direction) String() string {
	return string(d)
}

type MarginType string

const (
	Isolated MarginType = "isolated"
	Cross    MarginType = "cross"
)

func (m MarginType) String() string {
	return string(m)
}

type Status string

const (
	Open       Status = "open"
	Closed     Status = "closed"
	Liquidated Status = "liquidated"
)

func (s Status) String() string {
	return string(s)
}
