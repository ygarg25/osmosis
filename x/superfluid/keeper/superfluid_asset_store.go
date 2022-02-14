package keeper

// This file handles

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/osmosis-labs/osmosis/x/superfluid/types"
)

func (k Keeper) SetSuperfluidAsset(ctx sdk.Context, asset types.SuperfluidAsset) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.KeyPrefixSuperfluidAsset)
	bz, err := proto.Marshal(&asset)
	if err != nil {
		panic(err)
	}
	prefixStore.Set([]byte(asset.Denom), bz)
}

func (k Keeper) DeleteSuperfluidAsset(ctx sdk.Context, denom string) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.KeyPrefixSuperfluidAsset)
	prefixStore.Delete([]byte(denom))
}

func (k Keeper) GetSuperfluidAsset(ctx sdk.Context, denom string) types.SuperfluidAsset {
	asset := types.SuperfluidAsset{}
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.KeyPrefixSuperfluidAsset)
	bz := prefixStore.Get([]byte(denom))
	if bz == nil {
		return asset
	}
	err := proto.Unmarshal(bz, &asset)
	if err != nil {
		panic(err)
	}
	return asset
}

func (k Keeper) GetAllSuperfluidAssets(ctx sdk.Context) []types.SuperfluidAsset {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.KeyPrefixSuperfluidAsset)
	iterator := prefixStore.Iterator(nil, nil)
	defer iterator.Close()

	assets := []types.SuperfluidAsset{}
	for ; iterator.Valid(); iterator.Next() {
		asset := types.SuperfluidAsset{}

		err := proto.Unmarshal(iterator.Value(), &asset)
		if err != nil {
			panic(err)
		}

		assets = append(assets, asset)
	}
	return assets
}

// TODO: Move this function to another file
func (k Keeper) GetRiskAdjustedOsmoValue(ctx sdk.Context, asset types.SuperfluidAsset, amount sdk.Int) sdk.Int {
	// TODO: we need to figure out how to do this later.
	minRiskFactor := k.GetParams(ctx).MinimumRiskFactor
	return amount.Sub(amount.ToDec().Mul(minRiskFactor).RoundInt())
}