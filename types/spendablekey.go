package types

import (
	"github.com/NebulousLabs/Sia/crypto"
)

// SpendableKey is a set of secret keys plus the corresponding unlock
// conditions.  The public key can be derived from the secret key and then
// matched to the corresponding public keys in the unlock conditions. All
// addresses that are to be used in 'FundSiacoins' or 'FundSiafunds' in the
// transaction builder must conform to this form of spendable key.
type SpendableKey struct {
	UnlockConditions UnlockConditions
	SecretKeys       []crypto.SecretKey
}
