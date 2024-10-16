package utils

import (
	"fmt"
	"math"

	"github.com/cordilleradev/bean/common/types"
)

func IsLongAsType(isLong bool) types.Direction {
	if isLong {
		return types.Long
	}
	return types.Short
}

func CalculatePositionDeltas(oldPositions, newPositions []types.FuturesPosition) []types.FuturesDelta {
	deltas := []types.FuturesDelta{}

	// Create maps for quick access
	oldPositionsMap := createPositionsMap(oldPositions)
	newPositionsMap := createPositionsMap(newPositions)

	// Detect changes
	for key, newPos := range newPositionsMap {
		if oldPos, exists := oldPositionsMap[key]; exists {
			delta := calculateDelta(oldPos, newPos)
			if delta != nil {
				deltas = append(deltas, *delta)
			}
		} else {
			delta := createNewPositionDelta(newPos)
			deltas = append(deltas, delta)
		}
	}

	// Detect removed positions
	for key, oldPos := range oldPositionsMap {
		if _, exists := newPositionsMap[key]; !exists {
			delta := createRemovedPositionDelta(oldPos)
			deltas = append(deltas, delta)
		}
	}

	if len(deltas) == 0 {
		return make([]types.FuturesDelta, 0)
	}

	return deltas

}

func createPositionsMap(positions []types.FuturesPosition) map[string]types.FuturesPosition {
	positionsMap := make(map[string]types.FuturesPosition)
	for _, pos := range positions {
		key := fmt.Sprintf("%s:%s:%s", pos.CollateralToken, pos.Market, pos.MarginType)
		positionsMap[key] = pos
	}
	return positionsMap
}

func calculateDelta(oldPos, newPos types.FuturesPosition) *types.FuturesDelta {
	entryPriceDelta := RoundToNDecimalsOrSigFigs(newPos.EntryPrice-oldPos.EntryPrice, 5)
	sizeDelta := RoundToNDecimalsOrSigFigs(newPos.SizeToken-oldPos.SizeToken, 5)
	sizeUsdDelta := RoundToNDecimalsOrSigFigs(newPos.SizeUsd-oldPos.SizeUsd, 5)
	leverageDelta := RoundToNDecimalsOrSigFigs(newPos.LeverageAmount-oldPos.LeverageAmount, 5)

	if !shouldAddDelta(oldPos, newPos, entryPriceDelta, sizeDelta, leverageDelta) {
		return nil
	}

	delta := types.FuturesDelta{
		Market:        newPos.Market,
		Status:        types.Open,
		Direction:     newPos.Direction,
		SizeDelta:     sizeDelta,
		SizeUsdDelta:  sizeUsdDelta,
		LeverageDelta: leverageDelta,
	}

	return &delta
}

func shouldAddDelta(oldPos, newPos types.FuturesPosition, entryPriceDelta, sizeDelta, leverageDelta float64) bool {
	return math.Abs(entryPriceDelta/oldPos.EntryPrice) > 0.001 ||
		math.Abs(sizeDelta/oldPos.SizeToken) > 0.001 ||
		math.Abs(leverageDelta/oldPos.LeverageAmount) > 1 ||
		oldPos.MarginType != newPos.MarginType
}

func createNewPositionDelta(newPos types.FuturesPosition) types.FuturesDelta {
	delta := types.FuturesDelta{
		Market:        newPos.Market,
		Status:        types.Open,
		Direction:     newPos.Direction,
		SizeDelta:     RoundToNDecimalsOrSigFigs(newPos.SizeToken, 5),
		SizeUsdDelta:  RoundToNDecimalsOrSigFigs(newPos.SizeUsd, 5),
		LeverageDelta: RoundToNDecimalsOrSigFigs(newPos.LeverageAmount, 5),
	}

	return delta
}

func createRemovedPositionDelta(oldPos types.FuturesPosition) types.FuturesDelta {
	delta := types.FuturesDelta{
		Market:        oldPos.Market,
		Status:        types.Closed,
		Direction:     oldPos.Direction,
		SizeDelta:     RoundToNDecimalsOrSigFigs(-oldPos.SizeToken, 5),
		SizeUsdDelta:  RoundToNDecimalsOrSigFigs(-oldPos.SizeUsd, 5),
		LeverageDelta: RoundToNDecimalsOrSigFigs(-oldPos.LeverageAmount, 5),
	}

	return delta
}
