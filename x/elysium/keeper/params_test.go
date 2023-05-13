package keeper_test

import (
	"errors"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/furyaxyz/elysium/v2/app"
	elysiummodulekeeper "github.com/furyaxyz/elysium/v2/x/elysium/keeper"
	keepertest "github.com/furyaxyz/elysium/v2/x/elysium/keeper/mock"
	"github.com/furyaxyz/elysium/v2/x/elysium/types"
)

func (suite *KeeperTestSuite) TestGetSourceChannelID() {
	testCases := []struct {
		name          string
		ibcDenom      string
		expectedError error
		postCheck     func(channelID string)
	}{
		{
			"wrong ibc denom",
			"test",
			errors.New("test is invalid: ibc ely denom is invalid"),
			func(channelID string) {},
		},
		{
			"correct ibc denom",
			types.IbcElyDenomDefaultValue,
			nil,
			func(channelID string) {
				suite.Require().Equal(channelID, "channel-0")
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest() // reset
			// Create Elysium Keeper with mock transfer keeper
			elysiumKeeper := *elysiummodulekeeper.NewKeeper(
				app.MakeEncodingConfig().Codec,
				suite.app.GetKey(types.StoreKey),
				suite.app.GetKey(types.MemStoreKey),
				suite.app.BankKeeper,
				keepertest.IbcKeeperMock{},
				suite.app.GravityKeeper,
				suite.app.EvmKeeper,
				suite.app.AccountKeeper,
				authtypes.NewModuleAddress(govtypes.ModuleName).String(),
			)
			suite.app.ElysiumKeeper = elysiumKeeper

			channelID, err := suite.app.ElysiumKeeper.GetSourceChannelID(suite.ctx, tc.ibcDenom)
			if tc.expectedError != nil {
				suite.Require().EqualError(err, tc.expectedError.Error())
			} else {
				suite.Require().NoError(err)
				tc.postCheck(channelID)
			}
		})
	}
}
