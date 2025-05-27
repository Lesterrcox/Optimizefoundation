package keeper

import (
	"optimizeglobalcoin/x/oconsensus/types"
)

var _ types.QueryServer = Keeper{}
