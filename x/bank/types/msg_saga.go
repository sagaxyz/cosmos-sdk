package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// bank message types
// const (
// 	TypeMsgSetMetadata = "set_metadata"
// )

var (
	_ sdk.Msg = &MsgSetMetadata{}
)

// NewMsgSetMetadata - construct a msg to send coins from one account to another.
//
//nolint:interfacer
func NewMsgSetMetadata(authority sdk.AccAddress, metadata Metadata) *MsgSetMetadata {
	return &MsgSetMetadata{Authority: authority.String(), Metadata: metadata}
}

// // Route Implements Msg.
// func (msg MsgSetMetadata) Route() string { return RouterKey }

// // Type Implements Msg.
// func (msg MsgSetMetadata) Type() string { return TypeMsgSetMetadata }

// // ValidateBasic Implements Msg.
// func (msg MsgSetMetadata) ValidateBasic() error {
// 	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
// 		return sdkerrors.ErrInvalidAddress.Wrapf("invalid sender address: %s", err)
// 	}

// 	if err := msg.Metadata.Validate(); err != nil {
// 		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid metadata: %v", err)
// 	}

// 	return nil
// }

// // GetSignBytes Implements Msg.
// func (msg MsgSetMetadata) GetSignBytes() []byte {
// 	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
// }

// // GetSigners Implements Msg.
// func (msg MsgSetMetadata) GetSigners() []sdk.AccAddress {
// 	signer, _ := sdk.AccAddressFromBech32(msg.Authority)
// 	return []sdk.AccAddress{signer}
// }
