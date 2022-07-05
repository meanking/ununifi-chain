package keeper

import (
	"github.com/UnUniFi/chain/x/nftmint/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
)

func (k Keeper) MintNFT(ctx sdk.Context, msg *types.MsgMintNFT) error {
	exists := k.nftKeeper.HasClass(ctx, msg.ClassId)
	if !exists {
		return sdkerrors.Wrap(nfttypes.ErrClassExists, msg.ClassId)
	}

	classAttributes, exists := k.GetClassAttributes(ctx, msg.ClassId)
	if !exists {
		return sdkerrors.Wrapf(types.ErrClassAttributesNotExists, "class attributes with class id %s doesn't exist", msg.ClassId)
	}
	// TODO: validate minting permission from ClassAttributes
	err := types.ValidateMintingPermission(classAttributes, msg.Sender.AccAddress())
	if err != nil {
		return err
	}

	nftUri := classAttributes.BaseTokenUri + msg.NftId
	// TODO: validate uri
	// err := types.ValidateUri(nftUri)

	// TODO: validate token supply cap

	err = k.nftKeeper.Mint(ctx, types.NewNFT(msg.ClassId, msg.NftId, nftUri), msg.Recipient.AccAddress())
	if err != nil {
		return err
	}

	err = k.SetNFTAttributes(ctx, types.NewNFTAttributes(msg.ClassId, msg.NftId, msg.Sender.AccAddress()))
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) BurnNFT(ctx sdk.Context, msg *types.MsgBurnNFT) error {
	if !k.nftKeeper.HasClass(ctx, msg.ClassId) {
		return sdkerrors.Wrap(nfttypes.ErrClassNotExists, msg.ClassId)
	}

	if !k.nftKeeper.HasNFT(ctx, msg.ClassId, msg.NftId) {
		return sdkerrors.Wrap(nfttypes.ErrNFTNotExists, msg.NftId)
	}

	owner := k.nftKeeper.GetOwner(ctx, msg.ClassId, msg.NftId)
	if !owner.Equals(msg.Sender.AccAddress()) {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not the owner of nft %s", msg.Sender.AccAddress().String(), msg.NftId)
	}

	err := k.nftKeeper.Burn(ctx, msg.ClassId, msg.NftId)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) SetNFTAttributes(ctx sdk.Context, nftAttributes types.NFTAttributes) error {
	bz := k.cdc.MustMarshal(&nftAttributes)
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixNFTAttributes))

	nftAttributesKey := types.NFTAttributesKey(nftAttributes.ClassId, nftAttributes.NftId)
	prefixStore.Set(nftAttributesKey, bz)
	return nil
}

func (k Keeper) GetNFTAttributes(ctx sdk.Context, classID, nftID string) (types.NFTAttributes, bool) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixNFTAttributes))

	var nftAttributes types.NFTAttributes
	bz := prefixStore.Get(types.NFTAttributesKey(classID, nftID))
	if len(bz) == 0 {
		return types.NFTAttributes{}, false
	}

	k.cdc.MustUnmarshal(bz, &nftAttributes)
	return nftAttributes, true
}
