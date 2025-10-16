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

func (k msgServer) CreatePayment(ctx context.Context, msg *types.MsgCreatePayment) (*types.MsgCreatePaymentResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	nextId, err := k.PaymentSeq.Next(ctx)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "failed to get next id")
	}

	var payment = types.Payment{
		Id:         nextId,
		Creator:    msg.Creator,
		MerchantId: msg.MerchantId,
		CustomerId: msg.CustomerId,
		Amount:     msg.Amount,
		Status:     msg.Status,
		CreatedAt:  msg.CreatedAt,
	}

	if err = k.Payment.Set(
		ctx,
		nextId,
		payment,
	); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to set payment")
	}

	return &types.MsgCreatePaymentResponse{
		Id: nextId,
	}, nil
}

func (k msgServer) UpdatePayment(ctx context.Context, msg *types.MsgUpdatePayment) (*types.MsgUpdatePaymentResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	var payment = types.Payment{
		Creator:    msg.Creator,
		Id:         msg.Id,
		MerchantId: msg.MerchantId,
		CustomerId: msg.CustomerId,
		Amount:     msg.Amount,
		Status:     msg.Status,
		CreatedAt:  msg.CreatedAt,
	}

	// Checks that the element exists
	val, err := k.Payment.Get(ctx, msg.Id)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to get payment")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if err := k.Payment.Set(ctx, msg.Id, payment); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update payment")
	}

	return &types.MsgUpdatePaymentResponse{}, nil
}

func (k msgServer) DeletePayment(ctx context.Context, msg *types.MsgDeletePayment) (*types.MsgDeletePaymentResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// Checks that the element exists
	val, err := k.Payment.Get(ctx, msg.Id)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to get payment")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if err := k.Payment.Remove(ctx, msg.Id); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to delete payment")
	}

	return &types.MsgDeletePaymentResponse{}, nil
}
