package simulation_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/furyaxyz/elysium/v2/x/elysium/simulation"
	"github.com/furyaxyz/elysium/v2/x/elysium/types"
)

// TestRandomizedGenState tests the normal scenario of applying RandomizedGenState.
// Abonormal scenarios are not tested here.
func TestRandomizedGenState(t *testing.T) {
	registry := codectypes.NewInterfaceRegistry()
	types.RegisterInterfaces(registry)
	cdc := codec.NewProtoCodec(registry)

	s := rand.NewSource(1)
	r := rand.New(s)

	simState := module.SimulationState{
		AppParams:    make(simtypes.AppParams),
		Cdc:          cdc,
		Rand:         r,
		NumBonded:    3,
		Accounts:     simtypes.RandomAccounts(r, 3),
		InitialStake: math.NewInt(1000),
		GenState:     make(map[string]json.RawMessage),
	}

	simulation.RandomizedGenState(&simState)

	var elysiumGenesis types.GenesisState
	simState.Cdc.MustUnmarshalJSON(simState.GenState[types.ModuleName], &elysiumGenesis)

	require.Equal(t, "ibc/7939cb6694d2c422acd208a0072939487f6999eb9d18a44784045d87f3c67cf2", elysiumGenesis.Params.GetIbcElyDenom())
	require.Equal(t, uint64(0x68255aaf95e94627), elysiumGenesis.Params.GetIbcTimeout())
	require.Equal(t, "cosmos1tnh2q55v8wyygtt9srz5safamzdengsnqeycj3", elysiumGenesis.Params.GetElysiumAdmin())
	require.Equal(t, true, elysiumGenesis.Params.GetEnableAutoDeployment())

	require.Equal(t, len(elysiumGenesis.ExternalContracts), 0)
	require.Equal(t, len(elysiumGenesis.AutoContracts), 0)
}
