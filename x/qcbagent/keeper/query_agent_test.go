package keeper_test

import (
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "qcb/testutil/keeper"
	"qcb/testutil/nullify"
	"qcb/x/qcbagent/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestAgentQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.OanagentKeeper(t)
	msgs := createNAgent(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetAgentRequest
		response *types.QueryGetAgentResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetAgentRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetAgentResponse{Agent: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetAgentRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetAgentResponse{Agent: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetAgentRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Agent(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestAgentQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.OanagentKeeper(t)
	msgs := createNAgent(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllAgentRequest {
		return &types.QueryAllAgentRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.AgentAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Agent), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Agent),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.AgentAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Agent), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Agent),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.AgentAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Agent),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.AgentAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
