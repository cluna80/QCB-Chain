package keeper

import (
	"context"
	"fmt"
	"oan/x/oanprotocol/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AddressStatus(goCtx context.Context, req *types.QueryAddressStatusRequest) (*types.QueryAddressStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)
	tier := k.GetAddressTier(ctx, req.Address)
	dailyReceived := k.GetDailyReceived(ctx, req.Address)

	dailyLimit := params.MaxDailyReceive
	maxBalance := params.MaxWalletBalance
	switch tier {
	case "unverified":
		dailyLimit = params.UnverifiedDailyReceive
		maxBalance = params.UnverifiedMaxBalance
	case "verified":
		maxBalance = params.VerifiedMaxBalance
	}

	remaining := uint64(0)
	if dailyLimit > dailyReceived {
		remaining = dailyLimit - dailyReceived
	}
	_ = remaining

	return &types.QueryAddressStatusResponse{
		Tier:          tier,
		DailyReceived: fmt.Sprintf("%d", dailyReceived),
		DailyLimit:    fmt.Sprintf("%d", dailyLimit),
		MaxBalance:    fmt.Sprintf("%d", maxBalance),
	}, nil
}
