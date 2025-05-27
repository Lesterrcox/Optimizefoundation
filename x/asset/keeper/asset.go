package keeper

import (
	"context"
	"encoding/binary"

	"optimizeglobalcoin/x/asset/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetAssetCount get the total number of asset
func (k Keeper) GetAssetCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.AssetCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetAssetCount set the total number of asset
func (k Keeper) SetAssetCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.AssetCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendAsset appends a asset in the store with a new id and update the count
func (k Keeper) AppendAsset(
	ctx context.Context,
	asset types.Asset,
) uint64 {
	// Create the asset
	count := k.GetAssetCount(ctx)

	// Set the ID of the appended value
	asset.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AssetKey))
	appendedValue := k.cdc.MustMarshal(&asset)
	store.Set(GetAssetIDBytes(asset.Id), appendedValue)

	// Update asset count
	k.SetAssetCount(ctx, count+1)

	return count
}

// SetAsset set a specific asset in the store
func (k Keeper) SetAsset(ctx context.Context, asset types.Asset) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AssetKey))
	b := k.cdc.MustMarshal(&asset)
	store.Set(GetAssetIDBytes(asset.Id), b)
}

// GetAsset returns a asset from its id
func (k Keeper) GetAsset(ctx context.Context, id uint64) (val types.Asset, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AssetKey))
	b := store.Get(GetAssetIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAsset removes a asset from the store
func (k Keeper) RemoveAsset(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AssetKey))
	store.Delete(GetAssetIDBytes(id))
}

// GetAllAsset returns all asset
func (k Keeper) GetAllAsset(ctx context.Context) (list []types.Asset) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AssetKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Asset
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAssetsByNameAndType returns available asset by the specified name and type.
func (k Keeper) GetAssetsByNameAndType(ctx context.Context, name string, assetType types.AssetType) (list []types.Asset) {
	return k.GetAssetListByFilter(ctx, func(asset types.Asset) bool {
		if asset.Name != name || asset.AssetType != assetType {
			return true
		}

		return false
	})
}

// GetAssetsForValidation returns available asset for validation.
func (k Keeper) GetAssetsForValidation(ctx context.Context ) (list []types.Asset) {
	return k.GetAssetListByFilter(ctx, func(asset types.Asset) bool {
		if asset.Status != types.AssetStatus_PENDING {
			return true
		}

		return false
	})
}
// GetAssetListByFilter returns the assets using filter handler.
func (k Keeper) GetAssetListByFilter(ctx context.Context, filterFn func(asset types.Asset) (skip bool)) (list []types.Asset) {

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AssetKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Asset
		k.cdc.MustUnmarshal(iterator.Value(), &val)

		if filterFn(val) {
			continue
		}
		list = append(list, val)
	}

	return list
}

// GetAssetIDBytes returns the byte representation of the ID
func GetAssetIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.AssetKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
