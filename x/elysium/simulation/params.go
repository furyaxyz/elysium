package simulation

// DONTCOVER

import (
	"fmt"
	"math/rand"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/furyaxyz/elysium/v2/x/elysium/types"
)

// ParamChanges defines the parameters that can be modified by param change proposals
// on the simulation.
func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyIbcElyDenom),
			func(r *rand.Rand) string {
				return fmt.Sprintf("%v", GenIbcElyDenom(r))
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyIbcTimeout),
			func(r *rand.Rand) string {
				return fmt.Sprintf("%v", GenIbcTimeout(r))
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyEnableAutoDeployment),
			func(r *rand.Rand) string {
				return fmt.Sprintf("%v", GenEnableAutoDeployment(r))
			},
		),
	}
}
