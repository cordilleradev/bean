package gmx

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/cordilleradev/bean/common/utils"
	"github.com/machinebox/graphql"
)

// market represents an individual market entity.
type market struct {
	ID         string `json:"id"`
	IndexToken string `json:"indexToken"`
	LongToken  string `json:"longToken"`
	ShortToken string `json:"shortToken"`
}

// token represents a single cryptocurrency token along with its price.
type token struct {
	Symbol   string  `json:"symbol"`
	Address  string  `json:"address"`
	Decimals int     `json:"decimals"`
	MidPrice float64 // MidPrice holds the calculated middle price for the token
}

// marketManager manages market and token states along with their prices.
type marketManager struct {
	client    *graphql.Client   // GraphQL client for fetching market data
	markets   map[string]market // Map to track markets by ID
	tokens    map[string]token  // Map to track tokens by address
	mu        sync.RWMutex      // Mutex to ensure concurrent map access safety
	pricesUrl string
	tokensUrl string
}

// newMarketManager initializes a new market manager, fetches the initial set of markets and tokens,
// and returns an instance of marketManager.
func newMarketManager(client *graphql.Client, pricesUrl, tokensUrl string) (*marketManager, error) {
	mm := &marketManager{
		client:    client,
		markets:   make(map[string]market),
		tokens:    make(map[string]token),
		pricesUrl: pricesUrl,
		tokensUrl: tokensUrl,
	}

	// Load initial markets
	if err := mm.loadMarkets(); err != nil {
		return nil, err
	}

	// Load tokens and prices
	if err := mm.loadTokensAndPrices(); err != nil {
		return nil, err
	}

	go func() {
		for {
			mm.loadMarkets()
			time.Sleep(60 * time.Second)
		}
	}()

	go func() {
		for {
			mm.loadTokensAndPrices()
			time.Sleep(5 * time.Second)
		}
	}()

	return mm, nil
}

// loadMarkets fetches market data from the GraphQL API and populates the internal map.
func (mm *marketManager) loadMarkets() error {
	req := graphql.NewRequest(`
		query {
			markets {
				id
				indexToken
				longToken
				shortToken
			}
		}
	`)
	var marketsResponse struct {
		Markets []market `json:"markets"`
	}

	if err := mm.client.Run(context.Background(), req, &marketsResponse); err != nil {
		return fmt.Errorf("failed to fetch markets: %w", err)
	}

	mm.mu.Lock()
	defer mm.mu.Unlock()
	for _, mkt := range marketsResponse.Markets {
		mm.markets[mkt.ID] = mkt
	}
	return nil
}

// loadTokensAndPrices fetches token information and their prices from APIs and populates the token map.
func (mm *marketManager) loadTokensAndPrices() error {
	// Fetch tokens
	tokenResp, err := http.Get(mm.tokensUrl)
	if err != nil {
		return fmt.Errorf("failed to fetch tokens: %w", err)
	}
	defer tokenResp.Body.Close()

	var tokensResponse struct {
		Tokens []token `json:"tokens"`
	}

	if err := json.NewDecoder(tokenResp.Body).Decode(&tokensResponse); err != nil {
		return fmt.Errorf("failed to decode tokens response: %w", err)
	}

	// Fetch prices
	priceResp, err := http.Get(mm.pricesUrl)
	if err != nil {
		return fmt.Errorf("failed to fetch prices: %w", err)
	}
	defer priceResp.Body.Close()

	var pricesResponse []struct {
		TokenAddress string `json:"tokenAddress"`
		MinPrice     string `json:"minPrice"`
		MaxPrice     string `json:"maxPrice"`
	}

	if err := json.NewDecoder(priceResp.Body).Decode(&pricesResponse); err != nil {
		return fmt.Errorf("failed to decode prices response: %w", err)
	}

	// Map to hold address to Price mapping
	priceMap := make(map[string]float64)
	for _, priceInfo := range pricesResponse {
		minPrice, _ := strconv.ParseFloat(priceInfo.MinPrice, 64)
		maxPrice, _ := strconv.ParseFloat(priceInfo.MaxPrice, 64)
		priceMap[priceInfo.TokenAddress] = (minPrice + maxPrice) / 2
	}

	mm.mu.Lock()
	defer mm.mu.Unlock()
	for _, token := range tokensResponse.Tokens {
		midPrice, ok := priceMap[token.Address]
		if ok {
			token.MidPrice = utils.RoundToNDecimalsOrSigFigs(midPrice/math.Pow(10, float64(30-token.Decimals)), 5)
		}

		mm.tokens[token.Address] = token
	}

	return nil
}

// getMarket retrieves a market by its ID, if it exists.
func (mm *marketManager) getMarket(id string) (market, bool) {
	mm.mu.RLock()
	defer mm.mu.RUnlock()

	mkt, exists := mm.markets[id]
	return mkt, exists
}

// getToken returns the token info by its address.
func (mm *marketManager) getToken(address string) (token, bool) {
	mm.mu.RLock()
	defer mm.mu.RUnlock()

	token, exists := mm.tokens[address]
	return token, exists
}
