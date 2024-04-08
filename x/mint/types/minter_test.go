package types

import (
	"math/rand"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestNextAnnualProvision(t *testing.T) {
	minter := DefaultInitialMinter()
	params := DefaultParams()

	tests := []struct {
		totalSupply  string
		setInflation sdk.Dec
		expected     string
	}{
		// with 0 total staking token supply, next annual inflation should increase by InflationRate
		{"0", sdk.NewDecWithPrec(13, 2), "9750000000000000000000000"},

		// with 75 mil total staking token supply, next annual inflation should be 0
		{"75000000000000000000000000", sdk.NewDecWithPrec(13, 2), "0"},

		// with 35mil total staking token supply, next annual inflation should increase by InflationRate
		{"35000000000000000000000000", sdk.NewDecWithPrec(13, 2), "5200000000000000000000000"},
	}
	for i, tc := range tests {
		minter.Inflation = tc.setInflation
		expected, _ := sdk.NewDecFromStr(tc.expected)
		totalSupplyConverted, _ := sdkmath.NewIntFromString(tc.totalSupply)
		annualProv := minter.NextAnnualProvisions(params, totalSupplyConverted)
		require.True(t, annualProv.Equal(expected),
			"test: %v\n\tExp: %v\n\tGot: %v\n",
			i, tc.expected, annualProv)
	}
}

func TestBlockProvision(t *testing.T) {
	minter := InitialMinter(sdk.NewDecWithPrec(1, 1))
	params := DefaultParams()

	secondsPerYear := int64(60 * 60 * 8766)

	tests := []struct {
		annualProvisions int64
		expProvisions    int64
	}{
		{secondsPerYear / 5, 1},
		{secondsPerYear/5 + 1, 1},
		{(secondsPerYear / 5) * 2, 2},
		{(secondsPerYear / 5) / 2, 0},
	}
	for i, tc := range tests {
		minter.AnnualProvisions = sdk.NewDec(tc.annualProvisions)
		provisions := minter.BlockProvision(params)

		expProvisions := sdk.NewCoin(params.MintDenom,
			sdk.NewInt(tc.expProvisions))

		require.True(t, expProvisions.IsEqual(provisions),
			"test: %v\n\tExp: %v\n\tGot: %v\n",
			i, tc.expProvisions, provisions)
	}
}

// Benchmarking :)
// previously using math.Int operations:
// BenchmarkBlockProvision-4 5000000 220 ns/op
//
// using sdk.Dec operations: (current implementation)
// BenchmarkBlockProvision-4 3000000 429 ns/op
func BenchmarkBlockProvision(b *testing.B) {
	b.ReportAllocs()
	minter := InitialMinter(sdk.NewDecWithPrec(1, 1))
	params := DefaultParams()

	s1 := rand.NewSource(100)
	r1 := rand.New(s1) //nolint:gosec // this is a benchmark and is not relevant to security
	minter.AnnualProvisions = sdk.NewDec(r1.Int63n(1000000))

	// run the BlockProvision function b.N times
	for n := 0; n < b.N; n++ {
		minter.BlockProvision(params)
	}
}

// Next annual provisions benchmarking
// BenchmarkNextAnnualProvisions-4 5000000 251 ns/op
func BenchmarkNextAnnualProvisions(b *testing.B) {
	b.ReportAllocs()
	minter := InitialMinter(sdk.NewDecWithPrec(1, 1))
	params := DefaultParams()
	totalSupply := sdk.NewInt(100000000000000)

	// run the NextAnnualProvisions function b.N times
	for n := 0; n < b.N; n++ {
		minter.NextAnnualProvisions(params, totalSupply)
	}
}
