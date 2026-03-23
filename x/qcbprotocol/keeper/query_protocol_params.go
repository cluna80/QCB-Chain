package keeper

import (
	"context"

	"qcb/x/qcbprotocol/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ProtocolParams(goCtx context.Context, req *types.QueryProtocolParamsRequest) (*types.QueryProtocolParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryProtocolParamsResponse{}, nil
}
