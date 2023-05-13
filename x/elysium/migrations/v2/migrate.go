package v2

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furyaxyz/elysium/v2/x/elysium/exported"
	"github.com/furyaxyz/elysium/v2/x/elysium/types"
)

// Migrate migrates the x/elysium module state from the consensus version 1 to
// version 2. Specifically, it takes the parameters that are currently stored
// and managed by the x/params modules and stores them directly into the x/elysium
// module state.
func Migrate(ctx sdk.Context, store sdk.KVStore, legacySubspace exported.Subspace, cdc codec.BinaryCodec) error {
	var currParams types.Params
	legacySubspace.GetParamSet(ctx, &currParams)

	if err := currParams.Validate(); err != nil {
		return err
	}

	bz := cdc.MustMarshal(&currParams)
	store.Set(types.ParamsKey, bz)

	return nil
}
