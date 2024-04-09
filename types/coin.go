package types

import (
	"math/big"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// TBD defines the default coin denomination used in TBDNetwork in:
	//
	// - Staking parameters: denomination used as stake in the dPoS chain
	// - Mint parameters: denomination minted due to fee distribution rewards
	// - Governance parameters: denomination used for spam prevention in proposal deposits
	// - Crisis parameters: constant fee denomination used for spam prevention to check broken invariant
	UTBD string = "utbd"

	// BaseDenomUnit defines the base denomination unit for TBDnetwork.
	// 1 tbd = 1x10^{BaseDenomUnit} utbd
	BaseDenomUnit = 18

	// DefaultGasPrice is default gas price for transactions
	DefaultGasPrice = 20
)

// PowerReduction defines the default power reduction value for staking
var PowerReduction = sdkmath.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(BaseDenomUnit), nil))

// NewTBDCoin is a utility function that returns an "utbd" coin with the given sdkmath.Int amount.
// The function will panic if the provided amount is negative.
func NewTBDCoin(amount sdkmath.Int) sdk.Coin {
	return sdk.NewCoin(UTBD, amount)
}

// NewTBDDecCoin is a utility function that returns an "utbd" decimal coin with the given sdkmath.Int amount.
// The function will panic if the provided amount is negative.
func NewTBDDecCoin(amount sdkmath.Int) sdk.DecCoin {
	return sdk.NewDecCoin(UTBD, amount)
}

// NewTBDCoinInt64 is a utility function that returns an "utbd" coin with the given int64 amount.
// The function will panic if the provided amount is negative.
func NewTBDCoinInt64(amount int64) sdk.Coin {
	return sdk.NewInt64Coin(UTBD, amount)
}
