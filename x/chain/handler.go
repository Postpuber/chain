package chain

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sandblockio/chain/x/chain/keeper"
	"github.com/sandblockio/chain/x/chain/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgOpenBeam:
			return handleMsgOpenBeam(ctx, k, msg)

		case *types.MsgIncreaseBeam:
			return handleMsgIncreaseBeam(ctx, k, msg)

		case *types.MsgCancelBeam:
			return handleMsgCancelBeam(ctx, k, msg)

		case *types.MsgCloseBeam:
			return handleMsgCloseBeam(ctx, k, msg)

		case *types.MsgClaimBeam:
			return handleMsgClaimBeam(ctx, k, msg)

		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}