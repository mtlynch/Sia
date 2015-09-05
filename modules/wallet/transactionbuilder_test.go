package wallet

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/NebulousLabs/Sia/crypto"
	"github.com/NebulousLabs/Sia/mocks"
	"github.com/NebulousLabs/Sia/modules"
	"github.com/NebulousLabs/Sia/types"
)

func createSiacoinOutput(amount uint) types.SiacoinOutput {
	return types.SiacoinOutput{Value: types.NewCurrency64(uint64(amount))}
}

func createMockWalletSpender(currentHeight uint) (mockWalletSpender *mocks.WalletSpender) {
	mockWalletSpender = new(mocks.WalletSpender)
	mockWalletSpender.On("AcquireLock").Return()
	mockWalletSpender.On("ConsensusSetHeight").Return(types.BlockHeight(currentHeight))
	mockWalletSpender.On("UnconfirmedProcessedTransactions").Return([]modules.ProcessedTransaction{})
	mockWalletSpender.On("ReleaseLock").Return()
	return mockWalletSpender
}

func setSiacoinOutputs(ws *mocks.WalletSpender, values ...uint) {
	mockOutputs := map[types.SiacoinOutputID]types.SiacoinOutput{}
	for _, value := range values {
		output := createSiacoinOutput(value)
		oid := types.SiacoinOutputID(crypto.HashObject(output))
		mockOutputs[oid] = output
	}
	ws.On("SiacoinOutputs").Return(mockOutputs)
}

func TestAddMinerFee(t *testing.T) {
	assert := assert.New(t)
	tb := transactionBuilder{}

	// Add some miner fees,
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

func TestFundSiacoinsWhenNoOutputsExist(t *testing.T) {
	// Create a mock WalletSpender that contains no outputs or unconfirmed
	// transactions
	mockWalletSpender := createMockWalletSpender(1)
	setSiacoinOutputs(mockWalletSpender)
	mockWalletSpender.On("SiacoinOutputs").Return(make(map[types.SiacoinOutputID]types.SiacoinOutput))

	tb := transactionBuilder{walletSpender: mockWalletSpender}
	err := tb.FundSiacoins(types.NewCurrency64(100))
	assert.Equal(t, modules.ErrLowBalance, err)
}

func TestFundSiacoinsWhenBalanceIsTooLow(t *testing.T) {
	mockWalletSpender := createMockWalletSpender(100)
	setSiacoinOutputs(mockWalletSpender, 10)
	mockWalletSpender.On("SpendHeight", mock.Anything).Return(types.BlockHeight(1))
	mockWalletSpender.On("Key", mock.Anything).Return(types.SpendableKey{})

	tb := transactionBuilder{walletSpender: mockWalletSpender}
	err := tb.FundSiacoins(types.NewCurrency64(11))
	assert.Equal(t, modules.ErrLowBalance, err)
}

func TestFundSiacoinsWhenOutputIsNotYetTimeUnlocked(t *testing.T) {
	mockCurrentHeight := uint(100)
	mockWalletSpender := createMockWalletSpender(mockCurrentHeight)
	setSiacoinOutputs(mockWalletSpender, 10)
	mockWalletSpender.On("SpendHeight", mock.Anything).Return(types.BlockHeight(1))
	mockWalletSpender.On("Key", mock.Anything).Return(types.SpendableKey{UnlockConditions: types.UnlockConditions{Timelock: types.BlockHeight(mockCurrentHeight + 1)}})

	tb := transactionBuilder{walletSpender: mockWalletSpender}
	err := tb.FundSiacoins(types.NewCurrency64(10))
	assert.Equal(t, modules.ErrLowBalance, err, "Spending the output should fail because it is timelocked until the next block.")
}

func TestFundSiacoinsWhenOutputIsAlreadyTimeUnlocked(t *testing.T) {
	// TODO: implement this
}

func TestFundSiacoinsWhenOutputWasReceivedTooRecently(t *testing.T) {
	outputSpendHeight := 5
	walletSpendHeight := uint(outputSpendHeight + RespendTimeout - 1)
	mockWalletSpender := createMockWalletSpender(walletSpendHeight)
	setSiacoinOutputs(mockWalletSpender, 10)
	mockWalletSpender.On("SpendHeight", mock.Anything).Return(types.BlockHeight(outputSpendHeight))

	tb := transactionBuilder{walletSpender: mockWalletSpender}
	err := tb.FundSiacoins(types.NewCurrency64(10))
	assert.Equal(t, modules.ErrPotentialDoubleSpend, err)
}

func TestFundSiacoinsWhenBalanceIsSufficient(t *testing.T) {
	walletHeight := uint(100)
	mockWalletSpender := createMockWalletSpender(walletHeight)
	mockOutputs := map[types.SiacoinOutputID]types.SiacoinOutput{}
	output := createSiacoinOutput(10)
	oid := types.SiacoinOutputID(crypto.HashObject(output))
	log.Printf("oid=%v", oid)
	mockOutputs[oid] = output
	mockWalletSpender.On("SiacoinOutputs").Return(mockOutputs)
	setSiacoinOutputs(mockWalletSpender, 10)
	mockWalletSpender.On("SpendHeight", mock.Anything).Return(types.BlockHeight(1))
	log.Printf("Key(%v)", output.UnlockHash)
	mockWalletSpender.On("Key", output.UnlockHash).Return(types.SpendableKey{})
	mockWalletSpender.On("NextPrimarySeedAddress").Return(types.UnlockConditions{}, nil)
	mockWalletSpender.On("SetSpendHeight", mock.Anything, mock.Anything).Return()

	tb := transactionBuilder{walletSpender: mockWalletSpender}
	err := tb.FundSiacoins(types.NewCurrency64(10))
	assert.Nil(t, err)
}
