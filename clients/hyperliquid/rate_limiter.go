package hyperliquid

import (
	"net/http"
	"sync"
	"time"
)

type submittedRequest struct {
	madeAt time.Time
	weight requestWeight
}

type hyperLiquidManager struct {
	mu                 sync.Mutex
	allowancePerMinute int
	submittedRequests  []*submittedRequest
	ApiUrl             string
	HttpClient         *http.Client
}

func newHyperLiquidManager(
	allowancePerMinute int,
	apiUrl string,
	initWithPriceCacheCall bool,
) *hyperLiquidManager {
	manager := &hyperLiquidManager{
		allowancePerMinute: allowancePerMinute,
		submittedRequests:  []*submittedRequest{},
		ApiUrl:             apiUrl,
		HttpClient:         &http.Client{},
	}

	if initWithPriceCacheCall {
		manager.submittedRequests = append(manager.submittedRequests,
			&submittedRequest{
				madeAt: time.Now(),
				weight: priceFetchWeight,
			})
	}

	return manager
}

func (hm *hyperLiquidManager) resetRequests() int {
	hm.mu.Lock()
	defer hm.mu.Unlock()

	totalWeightUsed := 0
	now := time.Now()

	newSubmittedRequests := hm.submittedRequests[:0]

	for _, request := range hm.submittedRequests {
		if now.Sub(request.madeAt) < 1*time.Minute {
			newSubmittedRequests = append(newSubmittedRequests, request)
			totalWeightUsed += int(request.weight)
		}
	}

	hm.submittedRequests = newSubmittedRequests
	return totalWeightUsed
}

func (hm *hyperLiquidManager) waitForToken(requestWeight requestWeight) {
	for {
		if int(requestWeight) <= hm.allowancePerMinute-hm.resetRequests() {

			hm.mu.Lock()
			hm.submittedRequests = append(hm.submittedRequests, &submittedRequest{
				madeAt: time.Now(),
				weight: requestWeight,
			})
			hm.mu.Unlock()

			return
		}
		time.Sleep(500 * time.Millisecond)
	}
}
