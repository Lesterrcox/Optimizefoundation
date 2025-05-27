package keeper

import (
	"optimizeglobalcoin/x/validator/types"
)

var _ types.QueryServer = Keeper{}
