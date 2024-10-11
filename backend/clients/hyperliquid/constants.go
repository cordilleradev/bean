package hyperliquid

const hyperliquidApiUrl = "https://api.hyperliquid.xyz/info"
const hyperLiquidUiApi = "https://api-ui.hyperliquid.xyz/info"
const hyperliquidWsUrl = "wss://api.hyperliquid.xyz/ws"
const leaderboardUrl = "https://stats-data.hyperliquid.xyz/Mainnet/leaderboard"

var timePeriodMap map[string]string = map[string]string{
	"1d":    "day",
	"7d":    "week",
	"30d":   "month",
	"total": "allTime",
}
