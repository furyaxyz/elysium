package elysium_test

import (
	"github.com/furyaxyz/elysium/v2/x/elysium"
	"github.com/furyaxyz/elysium/v2/x/elysium/types"
)

func (suite *ElysiumTestSuite) TestInitGenesis() {
	testCases := []struct {
		name     string
		malleate func()
		genState *types.GenesisState
		expPanic bool
	}{
		{
			"default",
			func() {},
			types.DefaultGenesis(),
			false,
		},
		{
			"Wrong ibcElyDenom length",
			func() {},
			&types.GenesisState{
				Params: types.Params{
					IbcElyDenom: "ibc/6B5A664BF0AF4F71B2F0BAA33141E2F1321242FBD5D19762F541EC971ACB086534",
				},
			},
			true,
		},
		{
			"Wrong ibcElyDenom prefix",
			func() {},
			&types.GenesisState{
				Params: types.Params{
					IbcElyDenom: "aaa/6B5A664BF0AF4F71B2F0BAA33141E2F1321242FBD5D19762F541EC971ACB0865",
				},
			},
			true,
		},
		{
			"Wrong denom in external token mapping",
			func() {},
			&types.GenesisState{
				ExternalContracts: []types.TokenMapping{
					{
						Denom:    "aaa/6B5A664BF0AF4F71B2F0BAA33141E2F1321242FBD5D19762F541EC971ACB0865",
						Contract: "0x0000000000000000000000000000000000000000",
					},
				},
			},
			true,
		},
		{
			"Wrong denom in auto token mapping",
			func() {},
			&types.GenesisState{
				AutoContracts: []types.TokenMapping{
					{
						Denom:    "aaa/6B5A664BF0AF4F71B2F0BAA33141E2F1321242FBD5D19762F541EC971ACB0865",
						Contract: "0x0000000000000000000000000000000000000000",
					},
				},
			},
			true,
		},
		{
			"Wrong contract in external token mapping",
			func() {},
			&types.GenesisState{
				ExternalContracts: []types.TokenMapping{
					{
						Denom:    "ibc/6B5A664BF0AF4F71B2F0BAA33141E2F1321242FBD5D19762F541EC971ACB0865",
						Contract: "0x00000000000000000000000000000000000000",
					},
				},
			},
			true,
		},
		{
			"Wrong contract in auto token mapping",
			func() {},
			&types.GenesisState{
				AutoContracts: []types.TokenMapping{
					{
						Denom:    "ibc/6B5A664BF0AF4F71B2F0BAA33141E2F1321242FBD5D19762F541EC971ACB0865",
						Contract: "0x00000000000000000000000000000000000000",
					},
				},
			},
			true,
		},
		{
			"Correct token mapping",
			func() {},
			&types.GenesisState{
				Params: types.DefaultParams(),
				ExternalContracts: []types.TokenMapping{
					{
						Denom:    "ibc/6B5A664BF0AF4F71B2F0BAA33141E2F1321242FBD5D19762F541EC971ACB0865",
						Contract: "0x0000000000000000000000000000000000000000",
					},
				},
				AutoContracts: []types.TokenMapping{
					{
						Denom:    "gravity0x0000000000000000000000000000000000000000",
						Contract: "0x0000000000000000000000000000000000000000",
					},
				},
			},
			false,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			tc.malleate()
			if tc.expPanic {
				suite.Require().Panics(
					func() {
						elysium.InitGenesis(suite.ctx, suite.app.ElysiumKeeper, *tc.genState)
					},
				)
			} else {
				suite.Require().NotPanics(
					func() {
						elysium.InitGenesis(suite.ctx, suite.app.ElysiumKeeper, *tc.genState)
					},
				)
			}
		})
	}
}

func (suite *ElysiumTestSuite) TestExportGenesis() {
	genesisState := elysium.ExportGenesis(suite.ctx, suite.app.ElysiumKeeper)
	suite.Require().Equal(genesisState.Params.IbcElyDenom, types.DefaultParams().IbcElyDenom)
}
