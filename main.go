package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/cordilleradev/bean/clients/hyperliquid"
)

func main() {
	pool, err := hyperliquid.NewHyperLiquidPool(false)
	if err != nil {
		log.Fatalf("Failed to create HyperLiquidPool: %v", err)
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var totalRequests int
	var requestPerSecond float64
	var totalRecords int
	requestsChan := make(chan struct{}, 10)

	address := "0xFDE316cc38F7b9F0711EcD4E9797AACdE4FaB7d0"

	startTime := time.Now()
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for range ticker.C {
			mu.Lock()
			elapsed := time.Since(startTime).Seconds()
			requestPerSecond = float64(totalRequests) / elapsed
			fmt.Printf("RPS: %.2f, Total Requests: %d, Total Records: %d\n", requestPerSecond, totalRequests, totalRecords)
			mu.Unlock()
		}
	}()

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		requestsChan <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-requestsChan }()

			p, err := pool.Positions(address)
			if err != nil {
				log.Printf("Error fetching positions: %v", err)
				return
			}

			mu.Lock()
			totalRequests++
			totalRecords += len(p)
			mu.Unlock()
		}()
	}

	wg.Wait()
	ticker.Stop()
	close(requestsChan)

}
