package keeper

import (
	"context"
	"encoding/binary"

	"optimizeglobalcoin/x/asset/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetValueVotesCount get the total number of valueVotes
func (k Keeper) GetValueVotesCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.ValueVotesCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetValueVotesCount set the total number of valueVotes
func (k Keeper) SetValueVotesCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.ValueVotesCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendValueVotes appends a valueVotes in the store with a new id and update the count
func (k Keeper) AppendValueVotes(
	ctx context.Context,
	valueVotes types.ValueVotes,
) uint64 {
	// Create the valueVotes
	count := k.GetValueVotesCount(ctx)

	// Set the ID of the appended value
	valueVotes.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ValueVotesKey))
	appendedValue := k.cdc.MustMarshal(&valueVotes)
	store.Set(GetValueVotesIDBytes(valueVotes.Id), appendedValue)

	// Update valueVotes count
	k.SetValueVotesCount(ctx, count+1)

	return count
}

// SetValueVotes set a specific valueVotes in the store
func (k Keeper) SetValueVotes(ctx context.Context, valueVotes types.ValueVotes) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ValueVotesKey))
	b := k.cdc.MustMarshal(&valueVotes)
	store.Set(GetValueVotesIDBytes(valueVotes.Id), b)
}

// GetValueVotes returns a valueVotes from its id
func (k Keeper) GetValueVotes(ctx context.Context, id uint64) (val types.ValueVotes, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ValueVotesKey))
	b := store.Get(GetValueVotesIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveValueVotes removes a valueVotes from the store
func (k Keeper) RemoveValueVotes(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ValueVotesKey))
	store.Delete(GetValueVotesIDBytes(id))
}

// GetAllValueVotes returns all valueVotes
func (k Keeper) GetAllValueVotes(ctx context.Context) (list []types.ValueVotes) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ValueVotesKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ValueVotes
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetValueVotesIDBytes returns the byte representation of the ID
func GetValueVotesIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.ValueVotesKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}

// GetValueVotesByAssetIdAndVoter returns available valueVotes by the specified assetId and voter.
func (k Keeper) GetValueVotesByAssetIdAndVoter(ctx context.Context, assetId uint64, voter string) (list []types.ValueVotes) {
	return k.GetValueVotesListByFilter(ctx, func(valueVotes types.ValueVotes) bool {
		if valueVotes.AssetId != assetId || valueVotes.Voter != voter {
			return true
		}
		return false
	})
}
// GetValueVotesByAssetId returns available valueVotes by the specified assetId.
func (k Keeper) GetValueVotesByAssetId(ctx context.Context, assetId uint64) (list []types.ValueVotes) {
	return k.GetValueVotesListByFilter(ctx, func(valueVotes types.ValueVotes) bool {
		if valueVotes.AssetId != assetId {
			return true
		}
		return false
	})
}

// GetValueVotesListByFilter returns the valueVotes using filter handler.
func (k Keeper) GetValueVotesListByFilter(ctx context.Context, filterFn func(valueVotes types.ValueVotes) (skip bool)) (list []types.ValueVotes) {

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ValueVotesKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var valueVotes types.ValueVotes
		k.cdc.MustUnmarshal(iterator.Value(), &valueVotes)

		if filterFn(valueVotes) {
			continue
		}
		list = append(list, valueVotes)
	}

	return list
}
