package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
)

type UpdateResponse struct {
	Parsed []PriceUpdate `json:"parsed"`
}

type PriceUpdate struct {
	ID    string `json:"id"`
	Price struct {
		Price       string `json:"price"`
		Conf        string `json:"conf"`
		Expo        int    `json:"expo"`
		PublishTime int64  `json:"publish_time"`
	} `json:"price"`
	EmaPrice struct {
		Price       string `json:"price"`
		Conf        string `json:"conf"`
		Expo        int    `json:"expo"`
		PublishTime int64  `json:"publish_time"`
	} `json:"ema_price"`
	Metadata struct {
		Slot               int64 `json:"slot"`
		ProofAvailableTime int64 `json:"proof_available_time"`
		PrevPublishTime    int64 `json:"prev_publish_time"`
	} `json:"metadata"`
}

type PriceCache struct {
	URL   *url.URL
	cache map[string]float64
	mu    sync.RWMutex
}

func NewPriceCache() (*PriceCache, error) {
	feedIds, err := GetFeedIds()
	if err != nil {
		return nil, err
	}

	var ids []string
	for id := range feedIds {
		ids = append(ids, id)
	}

	rawQuery := "ids[]=" + strings.Join(ids, "&ids[]=")
	u := url.URL{Scheme: "https", Host: "hermes.pyth.network", Path: "/v2/updates/price/stream", RawQuery: rawQuery}

	priceCache := &PriceCache{
		URL:   &u,
		cache: make(map[string]float64),
	}

	return priceCache, nil
}

func GetFeedIds() (map[string]string, error) {
	resp, err := http.Get("https://raw.githubusercontent.com/Avantis-Labs/avantis_trader_sdk/main/avantis_trader_sdk/feed/feedIds.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var feedIds map[string]struct {
		ID    string `json:"id"`
		Group string `json:"group"`
	}

	err = json.Unmarshal(body, &feedIds)
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)
	for pair, data := range feedIds {
		pair = strings.ReplaceAll(pair, "/", "-")
		result[data.ID] = pair
	}

	return result, nil
}

func (p *PriceCache) FetchPriceUpdates() error {
	resp, err := http.Get(p.URL.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-OK HTTP status: %s", resp.Status)
	}

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		if scanner.Text() != "" {
			message := scanner.Text()[5:]
			var updateResponse UpdateResponse
			err := json.Unmarshal([]byte(message), &updateResponse)
			if err != nil {
				continue
			}

			for _, priceUpdate := range updateResponse.Parsed {
				unformattedPrice, err := strconv.ParseFloat(priceUpdate.Price.Price, 64)
				if err != nil {
					continue
				}
				realPrice := unformattedPrice / math.Pow(10, math.Abs(float64(priceUpdate.Price.Expo)))
				fmt.Println(realPrice, priceUpdate.ID)
				p.mu.Lock()
				p.cache[priceUpdate.ID] = realPrice
				p.mu.Unlock()
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil

}

func (p *PriceCache) ExportCache() map[string]float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	// Create a copy of the cache to export
	cacheCopy := make(map[string]float64)
	for k, v := range p.cache {
		cacheCopy[k] = v
	}

	return cacheCopy
}

func main() {
	priceCache, err := NewPriceCache()
	if err != nil {
		log.Fatalf("Failed to create PriceCache: %v", err)
	}

	// Use priceCache as needed
	err = priceCache.FetchPriceUpdates()
	if err != nil {
		log.Fatalf("Failed to fetch price updates: %v", err)
	}

}
