package types_test

import (
	"testing"

	"optimizeglobalcoin/x/asset/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				ValueVotesList: []types.ValueVotes{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				ValueVotesCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated valueVotes",
			genState: &types.GenesisState{
				ValueVotesList: []types.ValueVotes{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid valueVotes count",
			genState: &types.GenesisState{
				ValueVotesList: []types.ValueVotes{
					{
						Id: 1,
					},
				},
				ValueVotesCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
