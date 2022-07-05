package keeper

import (
	"context"
	"strconv"

	"github.com/UnUniFi/chain/x/nftmint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	keeper Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{
		keeper: keeper,
	}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) CreateClass(c context.Context, msg *types.MsgCreateClass) (*types.MsgCreateClassResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	seq, err := k.keeper.accountKeeper.GetSequence(ctx, msg.Sender.AccAddress())
	if err != nil {
		return nil, err
	}

	classID := CreateClassId(seq, msg.Sender.AccAddress())
	err = k.keeper.CreateClass(ctx, classID, msg)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitTypedEvent(&types.EventCreateClass{
		Owner:             msg.Sender.AccAddress().String(),
		ClassId:           classID,
		BaseTokenUri:      msg.BaseTokenUri,
		TokenSupplyCap:    strconv.FormatUint(msg.TokenSupplyCap, 10),
		MintingPermission: msg.MintingPermission,
	})

	return &types.MsgCreateClassResponse{}, nil
}

func (k msgServer) SendClass(c context.Context, msg *types.MsgSendClass) (*types.MsgSendClassResponse, error) {
	return &types.MsgSendClassResponse{}, nil
}
func (k msgServer) UpdateBaseTokenUri(c context.Context, msg *types.MsgUpdateBaseTokenUri) (*types.MsgUpdateBaseTokenUriResponse, error) {
	return &types.MsgUpdateBaseTokenUriResponse{}, nil
}
func (k msgServer) UpdateTokenSupplyCap(c context.Context, msg *types.MsgUpdateTokenSupplyCap) (*types.MsgUpdateTokenSupplyCapResponse, error) {
	return &types.MsgUpdateTokenSupplyCapResponse{}, nil
}

func (k msgServer) MintNFT(c context.Context, msg *types.MsgMintNFT) (*types.MsgMintNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if err := k.keeper.MintNFT(ctx, msg); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitTypedEvent(&types.EventMintNFT{
		ClassId: msg.ClassId,
		NftId:   msg.NftId,
		Owner:   msg.Recipient.AccAddress().String(),
		Minter:  msg.Sender.AccAddress().String(),
	})
	return &types.MsgMintNFTResponse{}, nil
}

func (k msgServer) BurnNFT(c context.Context, msg *types.MsgBurnNFT) (*types.MsgBurnNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if err := k.keeper.BurnNFT(ctx, msg); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitTypedEvent(&types.EventBurnNFT{
		Burner:  msg.Sender.AccAddress().String(),
		ClassId: msg.ClassId,
		NftId:   msg.NftId,
	})
	return &types.MsgBurnNFTResponse{}, nil
}
