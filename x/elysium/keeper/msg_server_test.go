package keeper_test

import (
	elysiummodulekeeper "github.com/furyaxyz/elysium/v2/x/elysium/keeper"
	"github.com/furyaxyz/elysium/v2/x/elysium/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func (suite *KeeperTestSuite) TestUpdateParams() {
	testCases := []struct {
		name      string
		req       *types.MsgUpdateParams
		expectErr bool
		expErrMsg string
	}{
		{
			name: "gov module account address as valid authority",
			req: &types.MsgUpdateParams{
				Authority: authtypes.NewModuleAddress(govtypes.ModuleName).String(),
				Params: types.Params{
					IbcElyDenom:          types.IbcElyDenomDefaultValue,
					IbcTimeout:           10,
					ElysiumAdmin:          sdk.AccAddress(suite.address.Bytes()).String(),
					EnableAutoDeployment: true,
				},
			},
			expectErr: false,
			expErrMsg: "",
		},
		{
			name: "set invalid authority",
			req: &types.MsgUpdateParams{
				Authority: "foo",
			},
			expectErr: true,
			expErrMsg: "invalid authority",
		},
		{
			name: "set invalid ibc ely denomination",
			req: &types.MsgUpdateParams{
				Authority: suite.app.ElysiumKeeper.GetAuthority(),
				Params: types.Params{
					IbcElyDenom:          "foo",
					IbcTimeout:           10,
					ElysiumAdmin:          sdk.AccAddress(suite.address.Bytes()).String(),
					EnableAutoDeployment: true,
				},
			},
			expectErr: true,
			expErrMsg: "invalid ibc denom",
		},
		{
			name: "set invalid elysium admin address",
			req: &types.MsgUpdateParams{
				Authority: suite.app.ElysiumKeeper.GetAuthority(),
				Params: types.Params{
					IbcElyDenom:          types.IbcElyDenomDefaultValue,
					IbcTimeout:           10,
					ElysiumAdmin:          "foo",
					EnableAutoDeployment: true,
				},
			},
			expectErr: true,
			expErrMsg: "invalid bech32 string",
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			msgServer := elysiummodulekeeper.NewMsgServerImpl(suite.app.ElysiumKeeper)
			_, err := msgServer.UpdateParams(suite.ctx, tc.req)
			if tc.expectErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}
