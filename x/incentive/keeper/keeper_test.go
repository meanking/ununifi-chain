package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	// authexported "github.com/cosmos/cosmos-sdk/x/auth/exported"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	// supplyexported "github.com/cosmos/cosmos-sdk/x/supply/exported"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/lcnem/jpyx/app"
	// committeekeeper "github.com/lcnem/jpyx/x/committee/keeper"
	// hardkeeper "github.com/lcnem/jpyx/x/hard/keeper"
	"github.com/lcnem/jpyx/x/incentive/keeper"
	"github.com/lcnem/jpyx/x/incentive/types"
)

// Test suite used for all keeper tests
type KeeperTestSuite struct {
	suite.Suite

	keeper keeper.Keeper
	// hardKeeper      hardkeeper.Keeper
	stakingKeeper stakingkeeper.Keeper
	// committeeKeeper committeekeeper.Keeper
	app            app.TestApp
	ctx            sdk.Context
	addrs          []sdk.AccAddress
	validatorAddrs []sdk.ValAddress
}

// The default state used by each test
func (suite *KeeperTestSuite) SetupTest() {
	tApp := app.NewTestApp()
	ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})

	tApp.InitializeFromGenesisStates()

	_, addrs := app.GeneratePrivKeyAddressPairs(5)
	keeper := tApp.GetIncentiveKeeper()
	suite.app = tApp
	suite.ctx = ctx
	suite.keeper = keeper
	suite.addrs = addrs
}

func (suite *KeeperTestSuite) getAccount(addr sdk.AccAddress) authtypes.AccountI {
	ak := suite.app.GetAccountKeeper()
	return ak.GetAccount(suite.ctx, addr)
}

func (suite *KeeperTestSuite) getModuleAccount(name string) authtypes.ModuleAccountI {
	sk := suite.app.GetAccountKeeper()
	return sk.GetModuleAccount(suite.ctx, name)
}

func (suite *KeeperTestSuite) TestGetSetDeleteJpyxMintingClaim() {
	c := types.NewJpyxMintingClaim(suite.addrs[0], c("ujsmn", 1000000), types.RewardIndexes{types.NewRewardIndex("bnb-a", sdk.ZeroDec())})
	_, found := suite.keeper.GetJpyxMintingClaim(suite.ctx, suite.addrs[0])
	suite.Require().False(found)
	suite.Require().NotPanics(func() {
		suite.keeper.SetJpyxMintingClaim(suite.ctx, c)
	})
	testC, found := suite.keeper.GetJpyxMintingClaim(suite.ctx, suite.addrs[0])
	suite.Require().True(found)
	suite.Require().Equal(c, testC)
	suite.Require().NotPanics(func() {
		suite.keeper.DeleteJpyxMintingClaim(suite.ctx, suite.addrs[0])
	})
	_, found = suite.keeper.GetJpyxMintingClaim(suite.ctx, suite.addrs[0])
	suite.Require().False(found)
}

func (suite *KeeperTestSuite) TestIterateJPYXMintingClaims() {
	for i := 0; i < len(suite.addrs); i++ {
		c := types.NewJpyxMintingClaim(suite.addrs[i], c("ujsmn", 100000), types.RewardIndexes{types.NewRewardIndex("bnb-a", sdk.ZeroDec())})
		suite.Require().NotPanics(func() {
			suite.keeper.SetJpyxMintingClaim(suite.ctx, c)
		})
	}
	claims := types.JpyxMintingClaims{}
	suite.keeper.IterateJpyxMintingClaims(suite.ctx, func(c types.JpyxMintingClaim) bool {
		claims = append(claims, c)
		return false
	})
	suite.Require().Equal(len(suite.addrs), len(claims))

	claims = suite.keeper.GetAllJpyxMintingClaims(suite.ctx)
	suite.Require().Equal(len(suite.addrs), len(claims))
}

func (suite *KeeperTestSuite) createPeriodicVestingAccount(origVesting sdk.Coins, periods vestingtypes.Periods, startTime, endTime int64) (*vestingtypes.PeriodicVestingAccount, error) {
	_, addr := app.GeneratePrivKeyAddressPairs(1)
	bk := suite.app.GetBankKeeper()
	bacc := authtypes.NewBaseAccountWithAddress(addr[0])
	bk.AddCoins(suite.ctx, bacc.GetAddress(), origVesting)
	bva := vestingtypes.NewBaseVestingAccount(bacc, origVesting, endTime)
	err := bva.Validate()
	if err != nil {
		return &vestingtypes.PeriodicVestingAccount{}, err
	}
	pva := vestingtypes.NewPeriodicVestingAccountRaw(bva, startTime, periods)
	err = pva.Validate()
	if err != nil {
		return &vestingtypes.PeriodicVestingAccount{}, err
	}
	return pva, nil
}

// Avoid cluttering test cases with long function names
func i(in int64) sdk.Int                    { return sdk.NewInt(in) }
func d(str string) sdk.Dec                  { return sdk.MustNewDecFromStr(str) }
func c(denom string, amount int64) sdk.Coin { return sdk.NewInt64Coin(denom, amount) }
func cs(coins ...sdk.Coin) sdk.Coins        { return sdk.NewCoins(coins...) }

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}