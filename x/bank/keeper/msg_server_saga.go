package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const SagaAddress = "saga1h8r6gm4jehflfn2nn7mtw53l37skrke5kyax8l"

func (k msgServer) SetMetadata(goCtx context.Context, msg *types.MsgSetMetadata) (*types.MsgSetMetadataResponse, error) {
	if k.GetAuthority() != msg.Authority {
		return nil, sdkerrors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.GetAuthority(), msg.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.Metadata.Validate(); err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid metadata: %v", err)
	}

	k.SetDenomMetaData(ctx, msg.Metadata)

	ctx.Logger().Info("Denomination metadata set", "authority", msg.Authority, "denom", msg.Metadata.Base)

	return &types.MsgSetMetadataResponse{}, nil
}
