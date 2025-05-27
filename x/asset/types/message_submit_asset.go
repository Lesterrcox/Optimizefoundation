package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSubmitAsset{}

func NewMsgSubmitAsset(creator string, name string, assetType AssetType, jurisdiction string, options string) *MsgSubmitAsset {
	return &MsgSubmitAsset{
		Creator:      creator,
		Name:         name,
		AssetType:    assetType,
		Jurisdiction: jurisdiction,
		Options:      options,
	}
}

func (msg *MsgSubmitAsset) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
