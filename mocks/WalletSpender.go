package mocks

import "github.com/NebulousLabs/Sia/modules"
import "github.com/stretchr/testify/mock"

import "github.com/NebulousLabs/Sia/types"

type WalletSpender struct {
	mock.Mock
}

func (_m *WalletSpender) AcquireLock() {
	_m.Called()
}
func (_m *WalletSpender) ReleaseLock() {
	_m.Called()
}
func (_m *WalletSpender) SiacoinOutputs() map[types.SiacoinOutputID]types.SiacoinOutput {
	ret := _m.Called()

	var r0 map[types.SiacoinOutputID]types.SiacoinOutput
	if rf, ok := ret.Get(0).(func() map[types.SiacoinOutputID]types.SiacoinOutput); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[types.SiacoinOutputID]types.SiacoinOutput)
		}
	}

	return r0
}
func (_m *WalletSpender) SiafundOutputs() map[types.SiafundOutputID]types.SiafundOutput {
	ret := _m.Called()

	var r0 map[types.SiafundOutputID]types.SiafundOutput
	if rf, ok := ret.Get(0).(func() map[types.SiafundOutputID]types.SiafundOutput); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[types.SiafundOutputID]types.SiafundOutput)
		}
	}

	return r0
}
func (_m *WalletSpender) UnconfirmedProcessedTransactions() []modules.ProcessedTransaction {
	ret := _m.Called()

	var r0 []modules.ProcessedTransaction
	if rf, ok := ret.Get(0).(func() []modules.ProcessedTransaction); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]modules.ProcessedTransaction)
		}
	}

	return r0
}
func (_m *WalletSpender) SpendHeight(_a0 types.OutputID) types.BlockHeight {
	ret := _m.Called(_a0)

	var r0 types.BlockHeight
	if rf, ok := ret.Get(0).(func(types.OutputID) types.BlockHeight); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(types.BlockHeight)
	}

	return r0
}
func (_m *WalletSpender) SetSpendHeight(_a0 types.OutputID, _a1 types.BlockHeight) {
	_m.Called(_a0, _a1)
}
func (_m *WalletSpender) RemoveOutput(_a0 types.OutputID) {
	_m.Called(_a0)
}
func (_m *WalletSpender) Key(_a0 types.UnlockHash) types.SpendableKey {
	ret := _m.Called(_a0)

	var r0 types.SpendableKey
	if rf, ok := ret.Get(0).(func(types.UnlockHash) types.SpendableKey); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(types.SpendableKey)
	}

	return r0
}
func (_m *WalletSpender) HasKey(_a0 types.UnlockHash) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(types.UnlockHash) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
func (_m *WalletSpender) ConsensusSetHeight() types.BlockHeight {
	ret := _m.Called()

	var r0 types.BlockHeight
	if rf, ok := ret.Get(0).(func() types.BlockHeight); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.BlockHeight)
	}

	return r0
}
func (_m *WalletSpender) NextPrimarySeedAddress() (types.UnlockConditions, error) {
	ret := _m.Called()

	var r0 types.UnlockConditions
	if rf, ok := ret.Get(0).(func() types.UnlockConditions); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.UnlockConditions)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
