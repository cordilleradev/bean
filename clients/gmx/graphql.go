package gmx

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/machinebox/graphql"
)

type periodAccountStats struct {
	ClosedCount           int     `json:"closedCount"`
	CumsumCollateral      float64 `json:"cumsumCollateral"`
	CumsumSize            float64 `json:"cumsumSize"`
	Losses                int     `json:"losses"`
	MaxCapital            float64 `json:"maxCapital"`
	NetCapital            float64 `json:"netCapital"`
	RealizedFees          float64 `json:"realizedFees"`
	RealizedPnl           float64 `json:"realizedPnl"`
	RealizedPriceImpact   float64 `json:"realizedPriceImpact"`
	StartUnrealizedFees   float64 `json:"startUnrealizedFees"`
	StartUnrealizedPnl    float64 `json:"startUnrealizedPnl"`
	StartUnrealizedImpact float64 `json:"startUnrealizedPriceImpact"`
	SumMaxSize            float64 `json:"sumMaxSize"`
	Volume                float64 `json:"volume"`
	Wins                  int     `json:"wins"`
	ID                    string  `json:"id"`
}

func (p *periodAccountStats) UnmarshalJSON(data []byte) error {
	type Alias periodAccountStats
	aux := &struct {
		CumsumCollateral      string `json:"cumsumCollateral"`
		CumsumSize            string `json:"cumsumSize"`
		MaxCapital            string `json:"maxCapital"`
		NetCapital            string `json:"netCapital"`
		RealizedFees          string `json:"realizedFees"`
		RealizedPnl           string `json:"realizedPnl"`
		RealizedPriceImpact   string `json:"realizedPriceImpact"`
		StartUnrealizedFees   string `json:"startUnrealizedFees"`
		StartUnrealizedPnl    string `json:"startUnrealizedPnl"`
		StartUnrealizedImpact string `json:"startUnrealizedPriceImpact"`
		SumMaxSize            string `json:"sumMaxSize"`
		Volume                string `json:"volume"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	p.CumsumCollateral = stringToFloat(aux.CumsumCollateral)
	p.CumsumSize = stringToFloat(aux.CumsumSize)
	p.MaxCapital = stringToFloat(aux.MaxCapital)
	p.NetCapital = stringToFloat(aux.NetCapital)
	p.RealizedFees = stringToFloat(aux.RealizedFees)
	p.RealizedPnl = stringToFloat(aux.RealizedPnl)
	p.RealizedPriceImpact = stringToFloat(aux.RealizedPriceImpact)
	p.StartUnrealizedFees = stringToFloat(aux.StartUnrealizedFees)
	p.StartUnrealizedPnl = stringToFloat(aux.StartUnrealizedPnl)
	p.StartUnrealizedImpact = stringToFloat(aux.StartUnrealizedImpact)
	p.SumMaxSize = stringToFloat(aux.SumMaxSize)
	p.Volume = stringToFloat(aux.Volume)

	return nil
}

func stringToFloat(s string) float64 {
	if s == "" {
		return 0
	}
	bigInt, _ := new(big.Int).SetString(s, 10)
	f, _ := new(big.Float).Quo(new(big.Float).SetInt(bigInt), big.NewFloat(1e30)).Float64()
	return f
}

func fetchPeriodAccountStats(client *graphql.Client, from int64) ([]periodAccountStats, error) {

	query := "" +
		"query PeriodAccountStats($requiredMaxCapital: String, $from: Int, $to: Int) {" +
		"	all: periodAccountStats(" +
		"		limit: 100000" +
		"		where: {maxCapital_gte: $requiredMaxCapital, from: $from, to: $to}" +
		"	) {" +
		"		id" +
		"		closedCount" +
		"		cumsumCollateral" +
		"		cumsumSize" +
		"		losses" +
		"		maxCapital" +
		"		realizedPriceImpact" +
		"		sumMaxSize" +
		"		netCapital" +
		"		realizedFees" +
		"		realizedPnl" +
		"		volume" +
		"		wins" +
		"		startUnrealizedPnl" +
		"		startUnrealizedFees" +
		"		startUnrealizedPriceImpact" +
		"		__typename" +
		"	}" +
		"}"

	to := 0
	if from > 0 {
		to = int(time.Now().Truncate(24 * time.Hour).Unix())
	}

	variables := map[string]interface{}{
		"requiredMaxCapital": "500000000000000000000000000000000",
		"from":               from,
		"to":                 to,
	}

	req := graphql.NewRequest(query)

	for key, value := range variables {
		req.Var(key, value)
	}

	req.Header.Set("Cache-Control", "no-cache")

	var response struct {
		All []periodAccountStats `json:"all"`
	}

	if err := client.Run(context.Background(), req, &response); err != nil {
		return nil, fmt.Errorf("decoding response - %v", err)
	}

	return response.All, nil
}
