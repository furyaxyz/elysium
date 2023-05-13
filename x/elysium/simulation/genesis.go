package simulation

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/furyaxyz/elysium/v2/x/elysium/types"
)

const (
	ibcElyDenomKey          = "ibc_ely_denom"
	ibcTimeoutKey           = "ibc_timeout"
	elysiumAdminKey          = "elysium_admin"
	enableAutoDeploymentKey = "enable_auto_deployment"
)

func GenIbcElyDenom(r *rand.Rand) string {
	randDenom := make([]byte, 32)
	r.Read(randDenom)
	return fmt.Sprintf("ibc/%s", hex.EncodeToString(randDenom))
}

func GenIbcTimeout(r *rand.Rand) uint64 {
	timeout := r.Uint64()
	return timeout
}

func GenElysiumAdmin(r *rand.Rand, simState *module.SimulationState) string {
	adminAccount, _ := simtypes.RandomAcc(r, simState.Accounts)
	return adminAccount.Address.String()
}

func GenEnableAutoDeployment(r *rand.Rand) bool {
	return r.Intn(2) > 0
}

// RandomizedGenState generates a random GenesisState for the elysium module
func RandomizedGenState(simState *module.SimulationState) {
	// elysium params
	var (
		ibcElyDenom          string
		ibcTimeout           uint64
		elysiumAdmin          string
		enableAutoDeployment bool
	)

	simState.AppParams.GetOrGenerate(
		simState.Cdc, ibcElyDenomKey, &ibcElyDenom, simState.Rand,
		func(r *rand.Rand) { ibcElyDenom = GenIbcElyDenom(r) },
	)

	simState.AppParams.GetOrGenerate(
		simState.Cdc, ibcTimeoutKey, &ibcTimeout, simState.Rand,
		func(r *rand.Rand) { ibcTimeout = GenIbcTimeout(r) },
	)

	simState.AppParams.GetOrGenerate(
		simState.Cdc, elysiumAdminKey, &elysiumAdmin, simState.Rand,
		func(r *rand.Rand) { elysiumAdmin = GenElysiumAdmin(r, simState) },
	)

	simState.AppParams.GetOrGenerate(
		simState.Cdc, enableAutoDeploymentKey, &enableAutoDeployment, simState.Rand,
		func(r *rand.Rand) { enableAutoDeployment = GenEnableAutoDeployment(r) },
	)

	params := types.NewParams(ibcElyDenom, ibcTimeout, elysiumAdmin, enableAutoDeployment)
	elysiumGenesis := &types.GenesisState{
		Params:            params,
		ExternalContracts: nil,
		AutoContracts:     nil,
	}

	bz, err := json.MarshalIndent(elysiumGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated %s parameters:\n%s\n", types.ModuleName, bz)

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(elysiumGenesis)
}
