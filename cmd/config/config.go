package cmd

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// func initSDKConfig() {
// 	// Set prefixes
// 	accountPubKeyPrefix := app.AccountAddressPrefix + "pub"
// 	validatorAddressPrefix := app.AccountAddressPrefix + "valoper"
// 	validatorPubKeyPrefix := app.AccountAddressPrefix + "valoperpub"
// 	consNodeAddressPrefix := app.AccountAddressPrefix + "valcons"
// 	consNodePubKeyPrefix := app.AccountAddressPrefix + "valconspub"

// 	// Set and seal config
// 	config := sdk.GetConfig()
// 	config.SetBech32PrefixForAccount(app.AccountAddressPrefix, accountPubKeyPrefix)
// 	config.SetBech32PrefixForValidator(validatorAddressPrefix, validatorPubKeyPrefix)
// 	config.SetBech32PrefixForConsensusNode(consNodeAddressPrefix, consNodePubKeyPrefix)
// 	config.Seal()
// }

const (
	// Bech32Prefix defines the Bech32 prefix for accounts
	Bech32Prefix = "tbd"

	// Bech32PrefixAccAddr defines the Bech32 prefix of an account's address
	Bech32PrefixAccAddr = Bech32Prefix
	// Bech32PrefixAccPub defines the Bech32 prefix of an account's public key
	Bech32PrefixAccPub = Bech32Prefix + sdk.PrefixPublic
	// Bech32PrefixValAddr defines the Bech32 prefix of a validator's operator address
	Bech32PrefixValAddr = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixOperator
	// Bech32PrefixValPub defines the Bech32 prefix of a validator's operator public key
	Bech32PrefixValPub = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixOperator + sdk.PrefixPublic
	// Bech32PrefixConsAddr defines the Bech32 prefix of a consensus node address
	Bech32PrefixConsAddr = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixConsensus
	// Bech32PrefixConsPub defines the Bech32 prefix of a consensus node public key
	Bech32PrefixConsPub = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixConsensus + sdk.PrefixPublic
)

const (
	// DisplayDenom defines the denomination displayed to users in client applications.
	DisplayDenom = "tbd"
	// BaseDenom defines to the default denomination used in Realio Network (staking, EVM, governance, etc.)
	BaseDenom = "utbd"
)

// SetBech32Prefixes sets the global prefixes to be used when serializing addresses and public keys to Bech32 strings.
func SetBech32Prefixes(config *sdk.Config) {
	config.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(Bech32PrefixValAddr, Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(Bech32PrefixConsAddr, Bech32PrefixConsPub)
}

// SetBip44CoinType sets the global coin type to be used in hierarchical deterministic wallets.
func SetBip44CoinType(config *sdk.Config) {
	config.SetPurpose(sdk.Purpose) // Shared
}

// RegisterDenoms registers the base and display denominations to the SDK.
func RegisterDenoms() {
	if err := sdk.RegisterDenom(DisplayDenom, sdk.OneDec()); err != nil {
		panic(err)
	}
}
