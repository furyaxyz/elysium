package simulation

import (
	"errors"
	"math/rand"

	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/ethereum/go-ethereum/common"

	"github.com/furyaxyz/elysium/v2/x/elysium/keeper"
	"github.com/furyaxyz/elysium/v2/x/elysium/types"
)

const (
	/* #nosec */
	OpWeightMsgUpdateTokenMapping = "op_weight_msg_update_token_mapping"
)

const (
	WeightMsgUpdateTokenMapping = 50
)

// WeightedOperations generate SimulateUpdateTokenMapping operation.
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec,
	ak types.AccountKeeper, bk types.BankKeeper, k *keeper.Keeper,
) simulation.WeightedOperations {
	var weightMsgUpdateTokenMapping int

	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateTokenMapping, &weightMsgUpdateTokenMapping, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateTokenMapping = WeightMsgUpdateTokenMapping
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgUpdateTokenMapping,
			SimulateUpdateTokenMapping(ak, bk, k),
		),
	}
}

// SimulateUpdateTokenMapping generate mocked MsgUpdateTokenMapping message, apply the message and assert the results.
func SimulateUpdateTokenMapping(ak types.AccountKeeper, bk types.BankKeeper, k *keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		elysiumAdmin := k.GetParams(ctx).ElysiumAdmin
		var simAccount simtypes.Account

		if r.Intn(2) > 0 {
			var found bool
			simAccount, found = findElysiumAdmin(accs, elysiumAdmin)
			if !found {
				simAccount, _ = simtypes.RandomAcc(r, accs)
			}
		} else {
			simAccount, _ = simtypes.RandomAcc(r, accs)
		}

		denom := GenIbcElyDenom(r)
		contractBytes := make([]byte, 20)
		r.Read(contractBytes)
		contract := common.BytesToAddress(contractBytes).String()
		expendable := bk.SpendableCoins(ctx, simAccount.Address)

		msg := types.NewMsgUpdateTokenMapping(simAccount.Address.String(), denom, contract, "", 0)

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: expendable,
		}

		oper, ops, err := simulation.GenAndDeliverTxWithRandFees(txCtx)
		if simAccount.Address.String() != elysiumAdmin && errors.Is(err, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "msg sender is authorized")) {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unauthorized tx should fail"), nil, nil
		}
		return oper, ops, err
	}
}

func findElysiumAdmin(accs []simtypes.Account, elysiumAdmin string) (simtypes.Account, bool) {
	found := false
	for _, acc := range accs {
		if acc.Address.String() == elysiumAdmin {
			found = true
			return acc, found
		}
	}
	return simtypes.Account{}, false
}
