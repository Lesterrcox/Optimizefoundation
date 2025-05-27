package keeper

import (
	"optimizeglobalcoin/x/asset/types"

	oconsensusTypes "optimizeglobalcoin/x/oconsensus/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) UpdateAssetsValueByVotes(ctx sdk.Context) {
	// Get all available transaction data
	assets := k.GetAssetsForValidation(ctx)

	// Update asset value by votes
	for _, asset := range assets {
		votes := k.GetValueVotesByAssetId(ctx, asset.Id)
		if len(votes) < 1 {
			continue
		}	

		// Update Transaction Request Data using votes
		k.UpdateAssetValue(ctx, votes, asset)
	}
}

func (k Keeper) UpdateAssetValue(ctx sdk.Context, votes []types.ValueVotes, asset types.Asset) {
	payloads := oconsensusTypes.ConsensusPayload{}
	// Loop to generate consensus payload
	for _, data := range votes {
		payloads.Payloads = append(payloads.Payloads, oconsensusTypes.Payload{
			Voter:     data.Voter,
			Result:    true,
			Value:     data.ValueUsd,
			ErrReason: "",
		})
	}

	// Run consensus algorithm in kconsensus module
	valueEvaluated, consensusResult, _ := k.oconsensusKeeper.RunConsensus(ctx, payloads)

	// TODO : Handler long pending ones as it couldn't have recived the enough votes.
	// Doesn't do anything with insufficient votes for now.
	// Later on, we should ask the active nodes to try posting the result again. (Need a discussion on this)
	if consensusResult == oconsensusTypes.CONSENSUS_PASS_SUCCESS {
		// Set tss hash to transaction data
		asset.Value.ValueUsd = valueEvaluated
		asset.Status = types.AssetStatus_VALIDATED
		for _, vote := range votes {
			asset.Value.Validators = append(asset.Value.Validators, vote.Voter)
			asset.Value.Timestamp = uint64(ctx.BlockHeader().Time.Unix())
		}
	} else if consensusResult == oconsensusTypes.CONSENSUS_PASS_FAILED {
		asset.Status = types.AssetStatus_INVALID
	}

	if consensusResult != oconsensusTypes.CONSENSUS_FAIL_INSUFFICIENT {
		k.SetAsset(ctx, asset)
	}
}