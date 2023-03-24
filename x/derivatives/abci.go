package derivatives

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/keeper"
	"github.com/UnUniFi/chain/x/derivatives/types"
)

func setPoolMarketCapSnapshot(ctx sdk.Context, k keeper.Keeper) {
	// move on only if price is ready
	if !k.IsPriceReady(ctx) {
		return
	}

	k.SetPoolMarketCapSnapshot(ctx, ctx.BlockHeight(), k.GetPoolMarketCap(ctx))
}

func saveBlockTime(ctx sdk.Context, k keeper.Keeper) {
	k.SaveBlockTimestamp(ctx, ctx.BlockHeight(), ctx.BlockTime())
}

// BeginBlocker
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	CheckPosition(ctx, k)
}

// todo: fixme this function is temporary treatment.
func CheckPosition(ctx sdk.Context, k keeper.Keeper) {
	positions := k.GetAllPositions(ctx)
	params := k.GetParams(ctx)
	for _, position := range positions {
		currentBaseUsdRate, currentQuoteUsdRate, err := k.GetPairUsdPriceFromMarket(ctx, position.Market)
		if err != nil {
			// todo: user logger
			fmt.Println("failed to get pair usd price from market")
			fmt.Println(err)
			continue
		}
		if position.NeedLiquidation(params.PerpetualFutures.MarginMaintenanceRate, currentBaseUsdRate, currentQuoteUsdRate) {
			msg := types.MsgReportLiquidation{
				Sender:          position.Address,
				PositionId:      position.Id,
				RewardRecipient: position.Address,
			}
			k.ReportLiquidation(ctx, &msg)
		}
		msg := types.MsgReportLevyPeriod{
			Sender:          position.Address,
			PositionId:      position.Id,
			RewardRecipient: position.Address,
		}
		k.ReportLevyPeriod(ctx, &msg)
	}
}

// EndBlocker
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	setPoolMarketCapSnapshot(ctx, k)
	saveBlockTime(ctx, k)
}