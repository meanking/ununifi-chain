package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftmarket/types"
)

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func validateListNftMsg(k Keeper, ctx sdk.Context, msg *types.MsgListNft) error {
	return checkListNft(k, ctx, msg.Sender.AccAddress(), msg.NftId, msg.BidToken)
}

func checkListNft(k Keeper, ctx sdk.Context, sender sdk.AccAddress, nftId types.NftIdentifier, bidToken string) error {
	// check listing already exists
	_, err := k.GetNftListingByIdBytes(ctx, nftId.IdBytes())
	if err == nil {
		return types.ErrNftListingAlreadyExists
	}

	// Check nft exists
	_, found := k.nftKeeper.GetNFT(ctx, nftId.ClassId, nftId.NftId)
	if !found {
		return types.ErrNftDoesNotExists
	}

	// check ownership of nft
	owner := k.nftKeeper.GetOwner(ctx, nftId.ClassId, nftId.NftId)
	if owner.String() != sender.String() {
		return types.ErrNotNftOwner
	}
	params := k.GetParamSet(ctx)
	for !Contains(params.BidTokens, bidToken) {
		return types.ErrNotSupportedBidToken
	}
	return nil
}
