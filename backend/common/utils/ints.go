package utils

import (
	"math"
	"math/big"
)

func MaxBigint() *big.Int {
	maxInt := big.NewInt(1)
	maxInt.Lsh(maxInt, 63).Sub(maxInt, big.NewInt(1)) // 2^63 - 1
	return maxInt
}

func RoundToNDecimalsOrSigFigs(value float64, decimalsOrSigFigs float64) float64 {
	const decimalPlacesThreshold = 1e5 // Threshold to switch between decimal places and significant figures
	orderOfMagnitude := math.Floor(math.Log10(math.Abs(value)))

	if value == 0 {
		return 0 // Special case for zero
	}

	if math.Abs(value) >= 1 {
		// Round to specified number of decimal places for larger numbers
		factor := math.Pow(10, decimalsOrSigFigs)
		return math.Round(value*factor) / factor
	} else {
		// Round to n significant figures for smaller numbers
		scale := math.Pow(10, -orderOfMagnitude+(decimalsOrSigFigs-1))
		return math.Round(value*scale) / scale
	}
}

func BigIntToRelevantFloat(value *big.Int, decimals float64, roundTo float64) float64 {
	floatVal, _ := value.Float64()
	relevantFloat := floatVal / math.Pow(10, decimals)
	return RoundToNDecimalsOrSigFigs(relevantFloat, roundTo)
}
