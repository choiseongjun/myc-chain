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

func (k msgServer) CreateMerchant(ctx context.Context, msg *types.MsgCreateMerchant) (*types.MsgCreateMerchantResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// Check if the value already exists
	ok, err := k.Merchant.Has(ctx, msg.Index)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	} else if ok {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var merchant = types.Merchant{
		Creator:      msg.Creator,
		Index:        msg.Index,
		Name:         msg.Name,
		Status:       msg.Status,
		RegisteredAt: msg.RegisteredAt,
	}

	if err := k.Merchant.Set(ctx, merchant.Index, merchant); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	return &types.MsgCreateMerchantResponse{}, nil
}

func (k msgServer) UpdateMerchant(ctx context.Context, msg *types.MsgUpdateMerchant) (*types.MsgUpdateMerchantResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.Merchant.Get(ctx, msg.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var merchant = types.Merchant{
		Creator:      msg.Creator,
		Index:        msg.Index,
		Name:         msg.Name,
		Status:       msg.Status,
		RegisteredAt: msg.RegisteredAt,
	}

	if err := k.Merchant.Set(ctx, merchant.Index, merchant); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update merchant")
	}

	return &types.MsgUpdateMerchantResponse{}, nil
}

func (k msgServer) DeleteMerchant(ctx context.Context, msg *types.MsgDeleteMerchant) (*types.MsgDeleteMerchantResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.Merchant.Get(ctx, msg.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if err := k.Merchant.Remove(ctx, msg.Index); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to remove merchant")
	}

	return &types.MsgDeleteMerchantResponse{}, nil
}
