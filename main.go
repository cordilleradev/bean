package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"runtime"
	"sort"
	"time"

	"github.com/cordilleradev/bean/clients/avantis"
	"github.com/cordilleradev/bean/clients/gains"
	"github.com/cordilleradev/bean/clients/gmx"
	"github.com/cordilleradev/bean/clients/hyperliquid"
	"github.com/cordilleradev/bean/clients/kwenta"
	"github.com/cordilleradev/bean/common"
)

func main() {
	breakpoint("Before initializing clients")
	gmxAr, _ := gmx.NewGmxArbitrumClient()
	gmxAx, _ := gmx.NewGmxAvalancheClient()
	hyp, _ := hyperliquid.NewHyperliquidClient()
	kweB, _ := kwenta.NewKwentaBaseClient()
	kweO, _ := kwenta.NewKwentaOptimismClient()
	gain, _ := gains.NewGainsClient()
	avan, _ := avantis.NewAvantisClient()

	breakpoint("After initializing clients")
	clients := []common.FuturesClient{gmxAr, gmxAx, hyp, kweB, kweO, gain, avan}
	file, err := os.Create("traders_leaderboard.csv")
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"user_id", "period_pnl_percent", "period_pnl_absolute", "total_trades", "wins", "volume", "avg_win", "avg_loss", "exchange", "period"}
	if err := writer.Write(header); err != nil {
		log.Fatalf("Failed to write header: %s", err)
	}

	breakpoint("Before iterating over clients")
	for _, client := range clients {
		supportedPeriods := client.GetLeaderboardPeriods().FixedPeriods
		for _, period := range supportedPeriods {
			traders, apiErr := client.GetLeaderboard(period)
			if apiErr != nil {
				log.Printf("Error fetching leaderboard for %s: %s", client.ExchangeName(), apiErr)
				continue
			}
			sort.Slice(traders, func(i, j int) bool {
				return traders[i].PeriodPnlAbsolute > traders[j].PeriodPnlAbsolute
			})
			for i, trader := range traders {
				if i >= 100 {
					break
				}
				record := []string{
					trader.UserId,
					formatFloat(trader.PeriodPnlPercent),
					formatFloat(trader.PeriodPnlAbsolute),
					formatInt(trader.TotalTrades),
					formatInt(trader.Wins),
					formatFloat(trader.Volume),
					formatFloat(trader.AvgWin),
					formatFloat(trader.AvgLoss),
					client.ExchangeName(),
					period,
				}
				if err := writer.Write(record); err != nil {
					log.Fatalf("Failed to write record: %s", err)
				}
			}
		}
		breakpoint(fmt.Sprintf("After processing client: %s", client.ExchangeName()))
	}
}

func formatFloat(value float64) string {
	return fmt.Sprintf("%.2f", value)
}

func formatInt(value int) string {
	return fmt.Sprintf("%d", value)
}

func breakpoint(message string) {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("BREAKPOINT: %s - %s:%d", message, file, line)
	time.Sleep(2 * time.Second) // Simulate a pause for debugging
}
