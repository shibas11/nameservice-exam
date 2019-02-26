package nameservice

import (
	"github.com/cosmos/cosmos-sdk/codec"     // Amino 인코딩 제공 툴
	sdk "github.com/cosmos/cosmos-sdk/types" // SDK 통해 사용되는 types
	"github.com/cosmos/cosmos-sdk/x/bank"    // accounts and coin transfers 제공 모듈
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	coinKeeper bank.Keeper // bank 모듈을 사용하기 위해서 포함(object capabilities approach)

	nameStoreKey   sdk.StoreKey // name 을 서비스
	ownerStoreKey  sdk.StoreKey // name 의 소유자
	pricesStoreKey sdk.StoreKey // name 의 가격

	cdc *codec.Codec
}

func (k Keeper) SetName(ctx sdk.Context, name string, value string) {
	store := ctx.KVStore(k.nameStoreKey)

	store.Set([]byte(name), []byte(value))
}

func (k Keeper) ResolveName(ctx sdk.Context, name string) string {
	store := ctx.KVStore(k.nameStoreKey)

	bz := store.Get([]byte(name))

	return string(bz)
}

func (k Keeper) HasOwner(ctx sdk.Context, name string) bool {
	store := ctx.KVStore(k.ownerStoreKey)

	bz := store.Get([]byte(name))

	return bz != nil
}

func (k Keeper) GetOwner(ctx sdk.Context, name string) sdk.AccAddress {
	store := ctx.KVStore(k.ownerStoreKey)

	bz := store.Get([]byte(name))

	return bz
}

func (k Keeper) SetOwner(ctx sdk.Context, name string, owner sdk.AccAddress) {
	store := ctx.KVStore(k.ownerStoreKey)

	store.Set([]byte(name), owner)
}

func (k Keeper) GetPrice(ctx sdk.Context, name string) sdk.Coins {
	if !k.HasOwner(ctx, name) {
		return sdk.Coins{sdk.NewInt64Coin("mycoin", 1)} // 소유자가 없으면, 1 mycoin 리턴
	}

	store := ctx.KVStore(k.pricesStoreKey)
	bz := store.Get([]byte(name))

	var price sdk.Coins
	k.cdc.MustUnmarshalBinaryBare(bz, &price)

	return price
}

func (k Keeper) SetPrice(ctx sdk.Context, name string, price sdk.Coins) {
	store := ctx.KVStore(k.pricesStoreKey)

	store.Set([]byte(name), k.cdc.MustMarshalBinaryBare(price))
}

func NewKeeper(coinKeeper bank.Keeper, namesStoreKey sdk.StoreKey, ownersStoreKey sdk.StoreKey, priceStoreKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper:     coinKeeper,
		nameStoreKey:   namesStoreKey,
		ownerStoreKey:  ownersStoreKey,
		pricesStoreKey: priceStoreKey,
		cdc:            cdc,
	}
}
