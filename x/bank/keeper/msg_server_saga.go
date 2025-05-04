package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
)

func (k msgServer) SetMetadata(goCtx context.Context, msg *types.MsgSetMetadata) (*types.MsgSetMetadataResponse, error) {
	if k.GetAuthority() != msg.Authority {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("invalid authority; expected %s, got %s", k.GetAuthority(), msg.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	err := msg.Metadata.Validate()
	if err != nil {
		return nil, err
	}

	k.SetDenomMetaData(ctx, msg.Metadata)

	return &types.MsgSetMetadataResponse{}, nil
}
