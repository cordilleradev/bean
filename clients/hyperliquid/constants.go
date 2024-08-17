package hyperliquid

const hyperliquidApiUrl = "https://app.hyperliquid.xyz/info"

var timePeriodMap map[string]string = map[string]string{
	"1d":    "day",
	"7d":    "week",
	"30d":   "month",
	"total": "allTime",
}
