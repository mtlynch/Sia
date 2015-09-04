package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/NebulousLabs/Sia/mocks"
	"github.com/NebulousLabs/Sia/types"
)

func TestAddMinerFee(t *testing.T) {
	assert := assert.New(t)
	mockWalletSpender := new(mocks.WalletSpender)

	tb := transactionBuilder{walletSpender: mockWalletSpender}
	assert.Equal(uint64(0), tb.AddMinerFee(types.NewCurrency64(25)))
	assert.Equal(uint64(1), tb.AddMinerFee(types.NewCurrency64(15)))
	assert.Equal(uint64(2), tb.AddMinerFee(types.NewCurrency64(19)))

	tr, parents := tb.View()
	assert.Nil(parents)
	assert.Equal(3, len(tr.MinerFees))
	assert.Equal(types.NewCurrency64(25), tr.MinerFees[0])
	assert.Equal(types.NewCurrency64(15), tr.MinerFees[1])
	assert.Equal(types.NewCurrency64(19), tr.MinerFees[2])
}
