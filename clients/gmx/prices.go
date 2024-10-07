package gmx

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/cordilleradev/bean/common/utils"
)

type priceCache struct {
	prices            map[string]float64
	waitPeriodSeconds float64
	pricesUrl         string
	mu                *sync.Mutex
}

func newPriceCache(waitPeriodSeconds float64, url string) (*priceCache, error) {
	pc := &priceCache{
		prices:            make(map[string]float64),
		waitPeriodSeconds: waitPeriodSeconds,
		mu:                &sync.Mutex{},
		pricesUrl:         url,
	}

	if err := pc.updatePrices(); err != nil {
		return nil, err
	}

	return pc, nil
}

func (pc *priceCache) updatePrices() error {
	resp, err := http.Get(pc.pricesUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var data map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
	}

	pc.mu.Lock()
	defer pc.mu.Unlock()

	for token, priceStr := range data {
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return err
		}
		pc.prices[token] = utils.RoundToNDecimalsOrSigFigs(price/1e30, 5) // Converting the price to float
	}

	return nil
}

func (pc *priceCache) streamPrices(waitSeconds float64) {
	for {
		time.Sleep(time.Second * time.Duration(waitSeconds))
		pc.updatePrices()
	}
}

func (pc *priceCache) getPrice(token string) float64 {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	price, _ := pc.prices[token]
	return price
}
