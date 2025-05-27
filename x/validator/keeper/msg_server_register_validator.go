package keeper

import (
	"context"

	"optimizeglobalcoin/x/validator/types"

	errorsmod "cosmossdk.io/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"optimizeglobalcoin/common"
)

func (k msgServer) RegisterValidator(goCtx context.Context, msg *types.MsgRegisterValidator) (*types.MsgRegisterValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if k.authority != msg.Authority {
		return nil, errorsmod.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, msg.Authority)
	}

	// Check if the ogc address is already a validator
	if k.CheckIfValidAccess(ctx, msg.OgcAddress) {
		return &types.MsgRegisterValidatorResponse{Code: common.RET_FAIL_INTERNAL, Message: "already a validator"}, nil
	}

	k.AppendValidator(ctx, types.Validator{
		Name:       msg.Name,
		Category:   msg.Category,
		Timestamp:  uint64(ctx.BlockTime().Unix()),
		OgcAddress: msg.Authority,
		KycInfo:    msg.KycInfo,
		Status:     types.ValidatorStatus_ACTIVE,
	})

	return &types.MsgRegisterValidatorResponse{Code: common.RET_OK, Message: "success"}, nil
}
