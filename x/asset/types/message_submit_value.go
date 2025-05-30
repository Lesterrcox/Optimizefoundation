package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSubmitValue{}

func NewMsgSubmitValue(creator string, assetId uint64, valueUsd string) *MsgSubmitValue {
	return &MsgSubmitValue{
		Creator:  creator,
		AssetId:  assetId,
		ValueUsd: valueUsd,
	}
}

func (msg *MsgSubmitValue) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
