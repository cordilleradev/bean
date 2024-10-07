package hyperliquid

import (
	"net/http"
	"sync"
	"time"
)

type SubmittedRequest struct {
	madeAt time.Time
	weight RequestWeight
}

type HyperLiquidManager struct {
	mu                 sync.Mutex
	allowancePerMinute int
	submittedRequests  []*SubmittedRequest
	ApiUrl             string
	HttpClient         *http.Client
}

func NewHyperLiquidManager(
	allowancePerMinute int,
	apiUrl string,
	initWithPriceCacheCall bool,
) *HyperLiquidManager {
	manager := &HyperLiquidManager{
		allowancePerMinute: allowancePerMinute,
		submittedRequests:  []*SubmittedRequest{},
		ApiUrl:             apiUrl,
		HttpClient:         &http.Client{},
	}

	if initWithPriceCacheCall {
		manager.submittedRequests = append(manager.submittedRequests,
			&SubmittedRequest{
				madeAt: time.Now(),
				weight: PriceFetchWeight,
			})
	}

	return manager
}

func (hm *HyperLiquidManager) resetRequests() int {
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

func (hm *HyperLiquidManager) waitForToken(requestWeight RequestWeight) {
	for {
		if int(requestWeight) <= hm.allowancePerMinute-hm.resetRequests() {

			hm.mu.Lock()
			hm.submittedRequests = append(hm.submittedRequests, &SubmittedRequest{
				madeAt: time.Now(),
				weight: requestWeight,
			})
			hm.mu.Unlock()

			return
		}
		time.Sleep(500 * time.Millisecond)
	}
}
