package types

type FuturesResponse struct {
	Trader             string            `json:"trader"`
	Platform           string            `json:"platform"`
	UpdatedPositions   []FuturesPosition `json:"updated_positions"`
	DetectedChanges    []FuturesDelta    `json:"detected_changes"`
	IsInitialPositions bool              `json:"is_initial_positions"`
}

type FuturesPosition struct {
	Market         string     `json:"market"`
	EntryPrice     float64    `json:"entry_price"`
	CurrentPrice   float64    `json:"current_price"`
	Status         Status     `json:"status"`
	Direction      Direction  `json:"direction"`
	MarginType     MarginType `json:"margin_type"`
	SizeUsd        float64    `json:"size_usd"`
	SizeToken      float64    `json:"size_tokens"`
	UnrealizedPnl  float64    `json:"unrealized_pnl"`
	LeverageAmount float64    `json:"leverage_amount,omitempty"`

	// Isolated margin specific
	CollateralTokenAmount    float64 `json:"collateral_token_amount,omitempty"`
	CollateralTokenAmountUsd float64 `json:"collateral_token_amount_usd,omitempty"`
	CollateralToken          string  `json:"collateral_token,omitempty"`

	// Cross margin specific
	HealthRatio      float64 `json:"health_ratio,omitempty"`
	CrossMarginShare float64 `json:"cross_margin_share,omitempty"`
	FreeCollateral   float64 `json:"free_collateral,omitempty"`
}

type FuturesDelta struct {
	Market        string    `json:"market"`
	Status        Status    `json:"status"`
	Direction     Direction `json:"direction"`
	SizeDelta     float64   `json:"size_tokens_delta"`
	SizeUsdDelta  float64   `json:"size_usd_delta"`
	LeverageDelta float64   `json:"leverage_delta"`

	// Isolated margin specific
	CollateralToken            string  `json:"collateral_token,omitempty"`
	CollateralTokenAmountDelta float64 `json:"collateral_token_amount_delta,omitempty"`
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
