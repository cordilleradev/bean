package kwenta

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

type respData struct {
	Top []struct {
		AccountOwner    string `json:"accountOwner"`
		Pnl             string `json:"pnl"`
		PnlWithFeesPaid string `json:"pnlWithFeesPaid"`
		Liquidations    string `json:"liquidations"`
		TotalTrades     string `json:"totalTrades"`
		TotalVolume     string `json:"totalVolume"`
	} `json:"top"`
}

func fetchLeaderboardStats(url, queryName string) (*respData, error) {
	client := graphql.NewClient(url)

	var accountFields string
	switch queryName {
	case "futuresStats":
		accountFields = "accountOwner: account"
	default:
		accountFields = "accountOwner"
	}

	req := graphql.NewRequest(fmt.Sprintf(`
		query leaderboardStats($first: Int!, $skip: Int!) {
			top: %s(
				first: $first
				skip: $skip
				orderBy: pnlWithFeesPaid
				orderDirection: desc
			) {
				%s
				pnl
				pnlWithFeesPaid
				liquidations
				totalTrades
				totalVolume
			}
		}
	`, queryName, accountFields))

	req.Var("first", 10000)
	req.Var("skip", 0)

	req.Header.Set("Origin", "https://kwenta.io")

	ctx := context.Background()
	var rawResp respData
	if err := client.Run(ctx, req, &rawResp); err != nil {
		return nil, fmt.Errorf("failed to run GraphQL query: %w", err)
	}

	return &rawResp, nil
}
