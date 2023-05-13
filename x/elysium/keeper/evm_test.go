package keeper_test

import (
	"math/big"

	"github.com/furyaxyz/elysium/v2/x/elysium/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/ethermint/crypto/ethsecp256k1"
)

func (suite *KeeperTestSuite) TestDeployContract() {
	suite.SetupTest()
	keeper := suite.app.ElysiumKeeper

	_, err := keeper.DeployModuleFRC21(suite.ctx, "test")
	suite.Require().NoError(err)
}

func (suite *KeeperTestSuite) TestTokenConversion() {
	suite.SetupTest()
	keeper := suite.app.ElysiumKeeper

	// generate test address
	priv, err := ethsecp256k1.GenerateKey()
	suite.Require().NoError(err)
	address := common.BytesToAddress(priv.PubKey().Address().Bytes())
	cosmosAddress := sdk.AccAddress(address.Bytes())

	denom := "ibc/0000000000000000000000000000000000000000000000000000000000000000"
	amount := big.NewInt(100)
	coins := sdk.NewCoins(sdk.NewCoin(denom, sdk.NewIntFromBigInt(amount)))

	// mint native tokens
	err = suite.MintCoins(sdk.AccAddress(address.Bytes()), coins)
	suite.Require().NoError(err)

	// send to erc20
	err = keeper.ConvertCoinsFromNativeToFRC21(suite.ctx, address, coins, true)
	suite.Require().NoError(err)

	// check erc20 balance
	contract, found := keeper.GetContractByDenom(suite.ctx, denom)
	suite.Require().True(found)

	ret, err := keeper.CallModuleFRC21(suite.ctx, contract, "balanceOf", address)
	suite.Require().NoError(err)
	suite.Require().Equal(amount, big.NewInt(0).SetBytes(ret))

	ret, err = keeper.CallModuleFRC21(suite.ctx, contract, "totalSupply")
	suite.Require().NoError(err)
	suite.Require().Equal(amount, big.NewInt(0).SetBytes(ret))

	// convert back to native
	err = keeper.ConvertCoinFromFRC21ToNative(suite.ctx, contract, address, coins[0].Amount)
	suite.Require().NoError(err)

	ret, err = keeper.CallModuleFRC21(suite.ctx, contract, "balanceOf", address)
	suite.Require().NoError(err)
	suite.Require().Equal(0, big.NewInt(0).Cmp(big.NewInt(0).SetBytes(ret)))

	ret, err = keeper.CallModuleFRC21(suite.ctx, contract, "totalSupply")
	suite.Require().NoError(err)
	suite.Require().Equal(0, big.NewInt(0).Cmp(big.NewInt(0).SetBytes(ret)))

	// native balance recovered
	coin := suite.app.BankKeeper.GetBalance(suite.ctx, cosmosAddress, denom)
	suite.Require().Equal(amount, coin.Amount.BigInt())
}

func (suite *KeeperTestSuite) TestSourceTokenConversion() {
	suite.SetupTest()
	keeper := suite.app.ElysiumKeeper

	// generate test address
	priv, err := ethsecp256k1.GenerateKey()
	suite.Require().NoError(err)
	address := common.BytesToAddress(priv.PubKey().Address().Bytes())
	cosmosAddress := sdk.AccAddress(address.Bytes())

	// Deploy FRC21 token
	contractAddress, err := keeper.DeployModuleFRC21(suite.ctx, "Test")
	suite.Require().NoError(err)

	// Register the token
	denom := "elysium" + contractAddress.Hex()
	msgUpdateTokenMapping := types.MsgUpdateTokenMapping{
		Sender:   cosmosAddress.String(),
		Denom:    denom,
		Contract: contractAddress.Hex(),
		Symbol:   "Test",
		Decimal:  0,
	}
	err = keeper.RegisterOrUpdateTokenMapping(suite.ctx, &msgUpdateTokenMapping)
	suite.Require().NoError(err)

	// Mint some FRC21 token
	amount := big.NewInt(100)
	_, err = suite.app.ElysiumKeeper.CallModuleFRC21(suite.ctx, contractAddress, "mint_by_elysium_module", address, amount)
	suite.Require().NoError(err)

	// Convert FRC21 to native
	err = keeper.ConvertCoinFromFRC21ToNative(suite.ctx, contractAddress, address, sdk.NewIntFromBigInt(amount))
	suite.Require().NoError(err)

	// Check balance
	coin := suite.app.BankKeeper.GetBalance(suite.ctx, cosmosAddress, denom)
	suite.Require().Equal(amount, coin.Amount.BigInt())

	// Convert native to FRC21
	coins := sdk.NewCoins(sdk.NewCoin(denom, sdk.NewIntFromBigInt(amount)))
	err = keeper.ConvertCoinsFromNativeToFRC21(suite.ctx, address, coins, false)
	suite.Require().NoError(err)

	// check balance
	coin = suite.app.BankKeeper.GetBalance(suite.ctx, cosmosAddress, denom)
	suite.Require().Equal(big.NewInt(0), coin.Amount.BigInt())
	ret, err := keeper.CallModuleFRC21(suite.ctx, contractAddress, "balanceOf", address)
	suite.Require().NoError(err)
	suite.Require().Equal(0, big.NewInt(100).Cmp(big.NewInt(0).SetBytes(ret)))
}
