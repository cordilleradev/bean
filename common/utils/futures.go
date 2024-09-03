package utils

import "github.com/cordilleradev/bean/common/types"

func IsLongAsType(isLong bool) types.Direction {
	if isLong {
		return types.Long
	}

	return types.Short
}
