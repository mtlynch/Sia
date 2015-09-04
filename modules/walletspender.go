package modules

import (
	"github.com/NebulousLabs/Sia/types"
)

// WalletSpender allows clients to retrieve and modify keys and outputs in a
// wallet in order to spend siacoins.
type WalletSpender interface {
	// acquireLock acquires a mutex on walletSpender. Clients must call
	// acquireLock before calling any other function in walletSpender and must
	// call releaseLock before exiting scope.
	AcquireLock()
	ReleaseLock()

	SiacoinOutputs() map[types.SiacoinOutputID]types.SiacoinOutput
	SiafundOutputs() map[types.SiafundOutputID]types.SiafundOutput
	UnconfirmedProcessedTransactions() []ProcessedTransaction

	// spendHeight returns the block height at which a given output was spent.
	SpendHeight(types.OutputID) types.BlockHeight

	// setSpendHeight marks the output with the given ID as spent in the given
	// block.
	SetSpendHeight(types.OutputID, types.BlockHeight)

	// removeOutput removes a given output from the wallet.
	RemoveOutput(types.OutputID)

	// key returns the key associated with the given hash.
	Key(types.UnlockHash) types.SpendableKey

	// hasKey indicates whether a key with the given hash exists in the wallet.
	HasKey(types.UnlockHash) bool

	// consensusSetHeight returns the current block height of the wallet.
	ConsensusSetHeight() types.BlockHeight

	// nextPrimarySeedAddress fetches the next address from the primary seed.
	NextPrimarySeedAddress() (types.UnlockConditions, error)
}
