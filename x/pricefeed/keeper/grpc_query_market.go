package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/pricefeed/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) MarketAll(c context.Context, req *types.QueryAllMarketRequest) (*types.QueryAllMarketResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var markets []types.Market
	ctx := sdk.UnwrapSDKContext(c)

	markets = k.GetMarkets(ctx)

	return &types.QueryAllMarketResponse{Markets: markets}, nil
}
