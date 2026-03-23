package keeper

import (
	"context"

	"qcb/x/qcbagent/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AgentAll(ctx context.Context, req *types.QueryAllAgentRequest) (*types.QueryAllAgentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var agents []types.Agent

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	agentStore := prefix.NewStore(store, types.KeyPrefix(types.AgentKeyPrefix))

	pageRes, err := query.Paginate(agentStore, req.Pagination, func(key []byte, value []byte) error {
		var agent types.Agent
		if err := k.cdc.Unmarshal(value, &agent); err != nil {
			return err
		}

		agents = append(agents, agent)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAgentResponse{Agent: agents, Pagination: pageRes}, nil
}

func (k Keeper) Agent(ctx context.Context, req *types.QueryGetAgentRequest) (*types.QueryGetAgentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetAgent(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetAgentResponse{Agent: val}, nil
}
