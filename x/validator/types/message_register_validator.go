package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRegisterValidator{}

func NewMsgRegisterValidator(authority string, ogcAddress string, name string, category ValidatorType, kycInfo string, options string) *MsgRegisterValidator {
	return &MsgRegisterValidator{
		Authority:  authority,
		OgcAddress: ogcAddress,
		Name:     name,
		Category: category,
		KycInfo:  kycInfo,
		Options:  options,
	}
}

func (msg *MsgRegisterValidator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address (%s)", err)
	}

	if msg.OgcAddress == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "ogc address is required")
	}

	if msg.Name == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "name is required")
	}

	if msg.Category == ValidatorType_UNRECOGNIZED {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid category")
	}

	if msg.KycInfo == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "kyc info is required")
	}

	return nil
}
