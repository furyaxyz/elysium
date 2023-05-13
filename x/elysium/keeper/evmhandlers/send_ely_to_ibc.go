package evmhandler

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	elysiumkeeper "github.com/furyaxyz/elysium/v2/x/elysium/keeper"
	"github.com/furyaxyz/elysium/v2/x/elysium/types"
)

var _ types.EvmLogHandler = SendElyToIbcHandler{}

const SendElyToIbcEventName = "__ElysiumSendElyToIbc"

// SendElyToIbcEvent represent the signature of
// `event __ElysiumSendElyToIbc(string recipient, uint256 amount)`
var SendElyToIbcEvent abi.Event

func init() {
	addressType, _ := abi.NewType("address", "", nil)
	uint256Type, _ := abi.NewType("uint256", "", nil)
	stringType, _ := abi.NewType("string", "", nil)

	SendElyToIbcEvent = abi.NewEvent(
		SendElyToIbcEventName,
		SendElyToIbcEventName,
		false,
		abi.Arguments{abi.Argument{
			Name:    "sender",
			Type:    addressType,
			Indexed: false,
		}, abi.Argument{
			Name:    "recipient",
			Type:    stringType,
			Indexed: false,
		}, abi.Argument{
			Name:    "amount",
			Type:    uint256Type,
			Indexed: false,
		}},
	)
}

// SendElyToIbcHandler handles `__ElysiumSendElyToIbc` log
type SendElyToIbcHandler struct {
	bankKeeper   types.BankKeeper
	elysiumKeeper elysiumkeeper.Keeper
}

func NewSendElyToIbcHandler(bankKeeper types.BankKeeper, elysiumKeeper elysiumkeeper.Keeper) *SendElyToIbcHandler {
	return &SendElyToIbcHandler{
		bankKeeper:   bankKeeper,
		elysiumKeeper: elysiumKeeper,
	}
}

func (h SendElyToIbcHandler) EventID() common.Hash {
	return SendElyToIbcEvent.ID
}

func (h SendElyToIbcHandler) Handle(
	ctx sdk.Context,
	contract common.Address,
	topics []common.Hash,
	data []byte,
	_ func(contractAddress common.Address, logSig common.Hash, logData []byte),
) error {
	unpacked, err := SendElyToIbcEvent.Inputs.Unpack(data)
	if err != nil {
		// log and ignore
		h.elysiumKeeper.Logger(ctx).Error("log signature matches but failed to decode", "error", err)
		return nil
	}

	contractAddr := sdk.AccAddress(contract.Bytes())
	sender := sdk.AccAddress(unpacked[0].(common.Address).Bytes())
	recipient := unpacked[1].(string)
	amount := sdk.NewIntFromBigInt(unpacked[2].(*big.Int))
	evmDenom := h.elysiumKeeper.GetEvmParams(ctx).EvmDenom
	coins := sdk.NewCoins(sdk.NewCoin(evmDenom, amount))
	// First, transfer IBC coin to user so that he will be the refunded address if transfer fails
	if err = h.bankKeeper.SendCoins(ctx, contractAddr, sender, coins); err != nil {
		return err
	}
	// Initiate IBC transfer from sender account
	if err = h.elysiumKeeper.IbcTransferCoins(ctx, sender.String(), recipient, coins, ""); err != nil {
		return err
	}
	return nil
}
