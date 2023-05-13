package types

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	// KeyIbcElyDenom is store's key for the IBC Ely denomination
	KeyIbcElyDenom = []byte("IbcElyDenom")
	// KeyIbcTimeout is store's key for the IBC Timeout
	KeyIbcTimeout = []byte("IbcTimeout")
	// KeyElysiumAdmin is store's key for the admin address
	KeyElysiumAdmin = []byte("ElysiumAdmin")
	// KeyEnableAutoDeployment is store's key for the EnableAutoDeployment
	KeyEnableAutoDeployment = []byte("EnableAutoDeployment")
)

const (
	IbcElyDenomDefaultValue = "ibc/6B5A664BF0AF4F71B2F0BAA33141E2F1321242FBD5D19762F541EC971ACB0865"
	IbcTimeoutDefaultValue  = uint64(86400000000000) // 1 day
)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new parameter configuration for the elysium module
func NewParams(ibcElyDenom string, ibcTimeout uint64, elysiumAdmin string, enableAutoDeployment bool) Params {
	return Params{
		IbcElyDenom:          ibcElyDenom,
		IbcTimeout:           ibcTimeout,
		ElysiumAdmin:          elysiumAdmin,
		EnableAutoDeployment: enableAutoDeployment,
	}
}

// DefaultParams is the default parameter configuration for the elysium module
func DefaultParams() Params {
	return Params{
		IbcElyDenom:          IbcElyDenomDefaultValue,
		IbcTimeout:           IbcTimeoutDefaultValue,
		ElysiumAdmin:          "",
		EnableAutoDeployment: false,
	}
}

// Validate all elysium module parameters
func (p Params) Validate() error {
	if err := validateIsUint64(p.IbcTimeout); err != nil {
		return err
	}
	if err := validateIsIbcDenom(p.IbcElyDenom); err != nil {
		return err
	}
	if len(p.ElysiumAdmin) > 0 {
		if _, err := sdk.AccAddressFromBech32(p.ElysiumAdmin); err != nil {
			return err
		}
	}
	return nil
}

// String implements the fmt.Stringer interface
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// ParamSetPairs implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyIbcElyDenom, &p.IbcElyDenom, validateIsIbcDenom),
		paramtypes.NewParamSetPair(KeyIbcTimeout, &p.IbcTimeout, validateIsUint64),
		paramtypes.NewParamSetPair(KeyElysiumAdmin, &p.ElysiumAdmin, validateIsAddress),
		paramtypes.NewParamSetPair(KeyEnableAutoDeployment, &p.EnableAutoDeployment, validateIsBool),
	}
}

func validateIsIbcDenom(i interface{}) error {
	s, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !IsValidIBCDenom(s) {
		return fmt.Errorf("invalid ibc denom: %T", i)
	}
	return nil
}

func validateIsUint64(i interface{}) error {
	if _, ok := i.(uint64); !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}

func validateIsAddress(i interface{}) error {
	s, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if len(s) > 0 {
		if _, err := sdk.AccAddressFromBech32(s); err != nil {
			return err
		}
	}
	return nil
}

func validateIsBool(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}
