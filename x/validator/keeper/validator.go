package keeper

import (
	"context"
	"encoding/binary"

	"optimizeglobalcoin/x/validator/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetValidatorCount get the total number of validator
func (k Keeper) GetValidatorCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.ValidatorCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetValidatorCount set the total number of validator
func (k Keeper) SetValidatorCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.ValidatorCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendValidator appends a validator in the store with a new id and update the count
func (k Keeper) AppendValidator(
	ctx context.Context,
	validator types.Validator,
) uint64 {
	// Create the validator
	count := k.GetValidatorCount(ctx)

	// Set the ID of the appended value
	validator.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ValidatorKey))
	appendedValue := k.cdc.MustMarshal(&validator)
	store.Set(GetValidatorIDBytes(validator.Id), appendedValue)

	// Update validator count
	k.SetValidatorCount(ctx, count+1)

	return count
}

// SetValidator set a specific validator in the store
func (k Keeper) SetValidator(ctx context.Context, validator types.Validator) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ValidatorKey))
	b := k.cdc.MustMarshal(&validator)
	store.Set(GetValidatorIDBytes(validator.Id), b)
}

// GetValidator returns a validator from its id
func (k Keeper) GetValidator(ctx context.Context, id uint64) (val types.Validator, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ValidatorKey))
	b := store.Get(GetValidatorIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveValidator removes a validator from the store
func (k Keeper) RemoveValidator(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ValidatorKey))
	store.Delete(GetValidatorIDBytes(id))
}

// GetAllValidator returns all validator
func (k Keeper) GetAllValidator(ctx context.Context) (list []types.Validator) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ValidatorKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Validator
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// CheckIfValidAccess checks if the voter is a validator
func (k Keeper) CheckIfValidAccess(ctx context.Context, voter string) bool {
	validators := k.GetAllValidator(ctx)
	for _, validator := range validators {
		if validator.Name == voter {
			return true
		}
	}
	return false
}

// GetValidatorIDBytes returns the byte representation of the ID
func GetValidatorIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.ValidatorKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
