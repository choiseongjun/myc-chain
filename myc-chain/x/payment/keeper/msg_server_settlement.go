package keeper

import (
	"context"
	"errors"
	"fmt"

	"myc-chain/x/payment/types"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateSettlement(ctx context.Context, msg *types.MsgCreateSettlement) (*types.MsgCreateSettlementResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	nextId, err := k.SettlementSeq.Next(ctx)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "failed to get next id")
	}

	var settlement = types.Settlement{
		Id:             nextId,
		Creator:        msg.Creator,
		MerchantId:     msg.MerchantId,
		TotalAmount:    msg.TotalAmount,
		SettlementDate: msg.SettlementDate,
		Status:         msg.Status,
	}

	if err = k.Settlement.Set(
		ctx,
		nextId,
		settlement,
	); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to set settlement")
	}

	return &types.MsgCreateSettlementResponse{
		Id: nextId,
	}, nil
}

func (k msgServer) UpdateSettlement(ctx context.Context, msg *types.MsgUpdateSettlement) (*types.MsgUpdateSettlementResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	var settlement = types.Settlement{
		Creator:        msg.Creator,
		Id:             msg.Id,
		MerchantId:     msg.MerchantId,
		TotalAmount:    msg.TotalAmount,
		SettlementDate: msg.SettlementDate,
		Status:         msg.Status,
	}

	// Checks that the element exists
	val, err := k.Settlement.Get(ctx, msg.Id)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to get settlement")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if err := k.Settlement.Set(ctx, msg.Id, settlement); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update settlement")
	}

	return &types.MsgUpdateSettlementResponse{}, nil
}

func (k msgServer) DeleteSettlement(ctx context.Context, msg *types.MsgDeleteSettlement) (*types.MsgDeleteSettlementResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// Checks that the element exists
	val, err := k.Settlement.Get(ctx, msg.Id)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to get settlement")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if err := k.Settlement.Remove(ctx, msg.Id); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to delete settlement")
	}

	return &types.MsgDeleteSettlementResponse{}, nil
}
