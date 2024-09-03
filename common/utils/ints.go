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

func RoundToNDecimals(value float64, decimals float64) float64 {
	if value < 1 {
		return math.Round(value*math.Pow(10, decimals)) / math.Pow(10, decimals)
	}
	factor := math.Pow(10, decimals)
	return math.Round(value*factor) / factor
}

func BigIntToRelevantFloat(value *big.Int, decimals float64, roundTo float64) float64 {
	floatVal, _ := value.Float64()
	relevantFloat := floatVal / math.Pow(10, decimals)
	return RoundToNDecimals(relevantFloat, roundTo)
}
