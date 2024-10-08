package hyperliquid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// hyperLiquidPriceCache maintains a live cache of prices.
type hyperLiquidPriceCache struct {
	url       string
	apiUrl    string
	mu        *sync.RWMutex
	pricesMap map[string]float64
	conn      *websocket.Conn
	done      chan struct{}
}

// newHyperLiquidPriceCache initializes a new hyperLiquidPriceCache instance and starts the WebSocket client.
func newHyperLiquidPriceCache(apiUrl, wsUrl string) (*hyperLiquidPriceCache, error) {
	if apiUrl == "" || wsUrl == "" {
		return nil, fmt.Errorf("URLs cannot be empty")
	}
	_, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("Invalid API URL: %v", err)
	}
	_, err = url.ParseRequestURI(wsUrl)
	if err != nil {
		return nil, fmt.Errorf("Invalid WebSocket URL: %v", err)
	}

	cache := &hyperLiquidPriceCache{
		url:       wsUrl,
		apiUrl:    apiUrl,
		mu:        new(sync.RWMutex),
		pricesMap: make(map[string]float64),
		done:      make(chan struct{}),
	}

	if err := cache.syncInitialPrices(); err != nil {
		return nil, fmt.Errorf("Error synchronizing initial prices: %v", err)
	}

	go cache.connect()

	return cache, nil
}

// syncInitialPrices gets the initial price data via a POST request and updates the pricesMap.
func (cache *hyperLiquidPriceCache) syncInitialPrices() error {
	requestBody, err := json.Marshal(map[string]string{"type": "allMids"})
	if err != nil {
		return fmt.Errorf("Error marshaling request body: %v", err)
	}

	resp, err := http.Post(cache.apiUrl, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("Error making POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error reading response body: %v", err)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return fmt.Errorf("Error unmarshaling response: %v", err)
	}

	return cache.updatePrices(response)
}

// connect establishes the WebSocket connection and initiates the data streaming.
func (cache *hyperLiquidPriceCache) connect() error {
	dialer := websocket.Dialer{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, _, err := dialer.Dial(cache.url, nil)
	if err != nil {
		return fmt.Errorf("Error connecting to WebSocket server: %v", err)
	}
	cache.conn = conn

	// Send initial query
	initialQuery := map[string]interface{}{
		"method": "subscribe",
		"subscription": map[string]string{
			"type": "allMids",
		},
	}
	if err := cache.conn.WriteJSON(initialQuery); err != nil {
		return fmt.Errorf("Error sending initial query: %v", err)
	}

	go cache.listen()
	go cache.sendHeartbeats()

	return nil
}

// listen continuously listens to incoming WebSocket messages and updates the cache.
func (cache *hyperLiquidPriceCache) listen() {
	defer cache.conn.Close()

	for {
		select {
		case <-cache.done:
			return
		default:
			_, message, err := cache.conn.ReadMessage()
			if err != nil {
				return
			}

			var response map[string]interface{}
			if err := json.Unmarshal(message, &response); err != nil {
				continue
			}

			cache.updatePrices(response)
		}
	}
}

// updatePrices updates the pricesMap with the given response data.
func (cache *hyperLiquidPriceCache) updatePrices(response map[string]interface{}) error {
	// Check if the response is from the initial HTTP request or WebSocket
	if data, dataOk := response["data"].(map[string]interface{}); dataOk {
		channel, channelOk := response["channel"].(string)
		if channelOk && channel == "allMids" {
			mids, midsOk := data["mids"].(map[string]interface{})
			if midsOk {
				cache.mu.Lock()
				for key, value := range mids {
					if strValue, ok := value.(string); ok {
						if floatValue, err := strconv.ParseFloat(strValue, 64); err == nil {
							cache.pricesMap[key] = floatValue
						}
					}
				}
				cache.mu.Unlock()
			}
		}
	} else {
		cache.mu.Lock()
		for key, value := range response {
			if strValue, ok := value.(string); ok {
				if floatValue, err := strconv.ParseFloat(strValue, 64); err == nil {
					cache.pricesMap[key] = floatValue
				}
			}
		}
		cache.mu.Unlock()
	}
	return nil
}

// sendHeartbeats sends periodic heartbeat messages to keep the connection alive.
func (cache *hyperLiquidPriceCache) sendHeartbeats() {
	ticker := time.NewTicker(50 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-cache.done:
			return
		case <-ticker.C:
			heartbeatMessage := map[string]interface{}{
				"method": "ping",
			}
			cache.mu.Lock()
			err := cache.conn.WriteJSON(heartbeatMessage)
			cache.mu.Unlock()
			if err != nil {
				return
			}
		}
	}
}

// getValue retrieves the cached value for a given key concurrently.
func (cache *hyperLiquidPriceCache) getValue(key string) float64 {
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	value, _ := cache.pricesMap[key]
	return value
}

// close signals the WebSocket listener to shut down.
func (cache *hyperLiquidPriceCache) close() {
	close(cache.done)
	cache.conn.Close()
}
