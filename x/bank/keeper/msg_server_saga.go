package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func (k msgServer) SetMetadata(goCtx context.Context, msg *types.MsgSetMetadata) (*types.MsgSetMetadataResponse, error) {
	if k.GetAuthority() != msg.Sender {
		return nil, sdkerrors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.GetAuthority(), msg.Sender)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	err := msg.Metadata.Validate()
	if err != nil {
		return nil, err
	}

	k.SetDenomMetaData(ctx, msg.Metadata)

	return &types.MsgSetMetadataResponse{}, nil
}
