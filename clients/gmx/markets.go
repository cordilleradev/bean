package gmx

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/machinebox/graphql"
)

// market represents the structure of each market in the GraphQL response
type market struct {
	ID         string `json:"id"`
	IndexToken string `json:"indexToken"`
	LongToken  string `json:"longToken"`
	ShortToken string `json:"shortToken"`
	Name       string
}

// marketsResponse represents the structure of the GraphQL response
type marketsResponse struct {
	Markets []market `json:"markets"`
}

// token represents the structure of each token in the API response
type token struct {
	Symbol   string `json:"symbol"`
	Address  string `json:"address"`
	Decimals int    `json:"decimals"`
}

// tokenInfo represents the structure of the token information in the map
type tokenInfo struct {
	Symbol   string
	Decimals int
}

// getMarkets retrieves the markets and tokens, returning any errors encountered
func getMarkets(client *graphql.Client, tokensUrl string) (tokenMap map[string]tokenInfo, marketMap map[string]market, err error) {

	tokenMap = make(map[string]tokenInfo)
	marketMap = make(map[string]market)

	// Prepare the GraphQL query
	req := graphql.NewRequest(`
		query MyQuery {
			markets {
				id
				indexToken
				longToken
				shortToken
			}
		}
	`)

	// Execute the GraphQL query
	var marketsResponse marketsResponse
	if err := client.Run(context.Background(), req, &marketsResponse); err != nil {
		return nil, nil, fmt.Errorf("failed to fetch markets: %w", err)
	}

	// Fetch the tokens from the API
	resp, err := http.Get(tokensUrl)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch tokens: %w", err)
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Parse the API response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read tokens response: %w", err)
	}

	var tokens struct {
		Tokens []token `json:"tokens"`
	}
	if err := json.Unmarshal(body, &tokens); err != nil {
		return nil, nil, fmt.Errorf("failed to parse tokens JSON: %w", err)
	}

	for _, t := range tokens.Tokens {
		tokenMap[t.Address] = tokenInfo{Symbol: t.Symbol, Decimals: t.Decimals}
	}

	for _, m := range marketsResponse.Markets {
		token, ok := tokenMap[m.IndexToken]
		if ok {
			m.Name = token.Symbol + "-" + "USD"
			marketMap[m.ID] = m
		}
	}

	return tokenMap, marketMap, nil
}
