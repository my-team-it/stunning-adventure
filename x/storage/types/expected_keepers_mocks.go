// Code generated by MockGen. DO NOT EDIT.
// Source: x/storage/types/expected_keepers.go

// Package types is a generated GoMock package.
package types

import (
	big "math/big"
	reflect "reflect"

	math "cosmossdk.io/math"
	types2 "github.com/cosmos/cosmos-sdk/types"
	types3 "github.com/cosmos/cosmos-sdk/x/auth/types"
	gomock "github.com/golang/mock/gomock"

	resource "github.com/bnb-chain/greenfield/types/resource"
	types "github.com/bnb-chain/greenfield/x/payment/types"
	types0 "github.com/bnb-chain/greenfield/x/permission/types"
	types1 "github.com/bnb-chain/greenfield/x/sp/types"
)

// MockAccountKeeper is a mock of AccountKeeper interface
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// GetAccount mocks base method
func (m *MockAccountKeeper) GetAccount(arg0 types2.Context, arg1 types2.AccAddress) types3.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", arg0, arg1)
	ret0, _ := ret[0].(types3.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount
func (mr *MockAccountKeeperMockRecorder) GetAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetAccount), arg0, arg1)
}

// GetModuleAddress mocks base method
func (m *MockAccountKeeper) GetModuleAddress(arg0 string) types2.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAddress", arg0)
	ret0, _ := ret[0].(types2.AccAddress)
	return ret0
}

// GetModuleAddress indicates an expected call of GetModuleAddress
func (mr *MockAccountKeeperMockRecorder) GetModuleAddress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAddress", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAddress), arg0)
}

// MockBankKeeper is a mock of BankKeeper interface
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// GetAllBalances mocks base method
func (m *MockBankKeeper) GetAllBalances(arg0 types2.Context, arg1 types2.AccAddress) types2.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBalances", arg0, arg1)
	ret0, _ := ret[0].(types2.Coins)
	return ret0
}

// GetAllBalances indicates an expected call of GetAllBalances
func (mr *MockBankKeeperMockRecorder) GetAllBalances(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBalances", reflect.TypeOf((*MockBankKeeper)(nil).GetAllBalances), arg0, arg1)
}

// GetBalance mocks base method
func (m *MockBankKeeper) GetBalance(arg0 types2.Context, arg1 types2.AccAddress, arg2 string) types2.Coin {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", arg0, arg1, arg2)
	ret0, _ := ret[0].(types2.Coin)
	return ret0
}

// GetBalance indicates an expected call of GetBalance
func (mr *MockBankKeeperMockRecorder) GetBalance(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockBankKeeper)(nil).GetBalance), arg0, arg1, arg2)
}

// SendCoinsFromModuleToAccount mocks base method
func (m *MockBankKeeper) SendCoinsFromModuleToAccount(arg0 types2.Context, arg1 string, arg2 types2.AccAddress, arg3 types2.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToAccount", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToAccount indicates an expected call of SendCoinsFromModuleToAccount
func (mr *MockBankKeeperMockRecorder) SendCoinsFromModuleToAccount(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToAccount", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromModuleToAccount), arg0, arg1, arg2, arg3)
}

// SpendableCoins mocks base method
func (m *MockBankKeeper) SpendableCoins(arg0 types2.Context, arg1 types2.AccAddress) types2.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpendableCoins", arg0, arg1)
	ret0, _ := ret[0].(types2.Coins)
	return ret0
}

// SpendableCoins indicates an expected call of SpendableCoins
func (mr *MockBankKeeperMockRecorder) SpendableCoins(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpendableCoins", reflect.TypeOf((*MockBankKeeper)(nil).SpendableCoins), arg0, arg1)
}

// MockSpKeeper is a mock of SpKeeper interface
type MockSpKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockSpKeeperMockRecorder
}

// MockSpKeeperMockRecorder is the mock recorder for MockSpKeeper
type MockSpKeeperMockRecorder struct {
	mock *MockSpKeeper
}

// NewMockSpKeeper creates a new mock instance
func NewMockSpKeeper(ctrl *gomock.Controller) *MockSpKeeper {
	mock := &MockSpKeeper{ctrl: ctrl}
	mock.recorder = &MockSpKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSpKeeper) EXPECT() *MockSpKeeperMockRecorder {
	return m.recorder
}

// GetSpStoragePriceByTime mocks base method
func (m *MockSpKeeper) GetSpStoragePriceByTime(arg0 types2.Context, arg1 types2.AccAddress, arg2 int64) (types1.SpStoragePrice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSpStoragePriceByTime", arg0, arg1, arg2)
	ret0, _ := ret[0].(types1.SpStoragePrice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSpStoragePriceByTime indicates an expected call of GetSpStoragePriceByTime
func (mr *MockSpKeeperMockRecorder) GetSpStoragePriceByTime(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSpStoragePriceByTime", reflect.TypeOf((*MockSpKeeper)(nil).GetSpStoragePriceByTime), arg0, arg1, arg2)
}

// GetStorageProvider mocks base method
func (m *MockSpKeeper) GetStorageProvider(arg0 types2.Context, arg1 types2.AccAddress) (*types1.StorageProvider, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageProvider", arg0, arg1)
	ret0, _ := ret[0].(*types1.StorageProvider)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetStorageProvider indicates an expected call of GetStorageProvider
func (mr *MockSpKeeperMockRecorder) GetStorageProvider(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageProvider", reflect.TypeOf((*MockSpKeeper)(nil).GetStorageProvider), arg0, arg1)
}

// GetStorageProviderByGcAddr mocks base method
func (m *MockSpKeeper) GetStorageProviderByGcAddr(arg0 types2.Context, arg1 types2.AccAddress) (*types1.StorageProvider, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageProviderByGcAddr", arg0, arg1)
	ret0, _ := ret[0].(*types1.StorageProvider)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetStorageProviderByGcAddr indicates an expected call of GetStorageProviderByGcAddr
func (mr *MockSpKeeperMockRecorder) GetStorageProviderByGcAddr(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageProviderByGcAddr", reflect.TypeOf((*MockSpKeeper)(nil).GetStorageProviderByGcAddr), arg0, arg1)
}

// GetStorageProviderBySealAddr mocks base method
func (m *MockSpKeeper) GetStorageProviderBySealAddr(arg0 types2.Context, arg1 types2.AccAddress) (*types1.StorageProvider, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageProviderBySealAddr", arg0, arg1)
	ret0, _ := ret[0].(*types1.StorageProvider)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetStorageProviderBySealAddr indicates an expected call of GetStorageProviderBySealAddr
func (mr *MockSpKeeperMockRecorder) GetStorageProviderBySealAddr(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageProviderBySealAddr", reflect.TypeOf((*MockSpKeeper)(nil).GetStorageProviderBySealAddr), arg0, arg1)
}

// IsStorageProviderExistAndInService mocks base method
func (m *MockSpKeeper) IsStorageProviderExistAndInService(arg0 types2.Context, arg1 types2.AccAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsStorageProviderExistAndInService", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// IsStorageProviderExistAndInService indicates an expected call of IsStorageProviderExistAndInService
func (mr *MockSpKeeperMockRecorder) IsStorageProviderExistAndInService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsStorageProviderExistAndInService", reflect.TypeOf((*MockSpKeeper)(nil).IsStorageProviderExistAndInService), arg0, arg1)
}

// SetSecondarySpStorePrice mocks base method
func (m *MockSpKeeper) SetSecondarySpStorePrice(arg0 types2.Context, arg1 types1.SecondarySpStorePrice) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSecondarySpStorePrice", arg0, arg1)
}

// SetSecondarySpStorePrice indicates an expected call of SetSecondarySpStorePrice
func (mr *MockSpKeeperMockRecorder) SetSecondarySpStorePrice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSecondarySpStorePrice", reflect.TypeOf((*MockSpKeeper)(nil).SetSecondarySpStorePrice), arg0, arg1)
}

// SetSpStoragePrice mocks base method
func (m *MockSpKeeper) SetSpStoragePrice(arg0 types2.Context, arg1 types1.SpStoragePrice) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSpStoragePrice", arg0, arg1)
}

// SetSpStoragePrice indicates an expected call of SetSpStoragePrice
func (mr *MockSpKeeperMockRecorder) SetSpStoragePrice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSpStoragePrice", reflect.TypeOf((*MockSpKeeper)(nil).SetSpStoragePrice), arg0, arg1)
}

// MockPaymentKeeper is a mock of PaymentKeeper interface
type MockPaymentKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentKeeperMockRecorder
}

// MockPaymentKeeperMockRecorder is the mock recorder for MockPaymentKeeper
type MockPaymentKeeperMockRecorder struct {
	mock *MockPaymentKeeper
}

// NewMockPaymentKeeper creates a new mock instance
func NewMockPaymentKeeper(ctrl *gomock.Controller) *MockPaymentKeeper {
	mock := &MockPaymentKeeper{ctrl: ctrl}
	mock.recorder = &MockPaymentKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPaymentKeeper) EXPECT() *MockPaymentKeeperMockRecorder {
	return m.recorder
}

// ApplyUserFlowsList mocks base method
func (m *MockPaymentKeeper) ApplyUserFlowsList(arg0 types2.Context, arg1 []types.UserFlows) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyUserFlowsList", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyUserFlowsList indicates an expected call of ApplyUserFlowsList
func (mr *MockPaymentKeeperMockRecorder) ApplyUserFlowsList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyUserFlowsList", reflect.TypeOf((*MockPaymentKeeper)(nil).ApplyUserFlowsList), arg0, arg1)
}

// GetParams mocks base method
func (m *MockPaymentKeeper) GetParams(arg0 types2.Context) types.Params {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParams", arg0)
	ret0, _ := ret[0].(types.Params)
	return ret0
}

// GetParams indicates an expected call of GetParams
func (mr *MockPaymentKeeperMockRecorder) GetParams(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParams", reflect.TypeOf((*MockPaymentKeeper)(nil).GetParams), arg0)
}

// GetStoragePrice mocks base method
func (m *MockPaymentKeeper) GetStoragePrice(arg0 types2.Context, arg1 types.StoragePriceParams) (types.StoragePrice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStoragePrice", arg0, arg1)
	ret0, _ := ret[0].(types.StoragePrice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStoragePrice indicates an expected call of GetStoragePrice
func (mr *MockPaymentKeeperMockRecorder) GetStoragePrice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStoragePrice", reflect.TypeOf((*MockPaymentKeeper)(nil).GetStoragePrice), arg0, arg1)
}

// GetStreamRecord mocks base method
func (m *MockPaymentKeeper) GetStreamRecord(arg0 types2.Context, arg1 types2.AccAddress) (*types.StreamRecord, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStreamRecord", arg0, arg1)
	ret0, _ := ret[0].(*types.StreamRecord)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetStreamRecord indicates an expected call of GetStreamRecord
func (mr *MockPaymentKeeperMockRecorder) GetStreamRecord(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStreamRecord", reflect.TypeOf((*MockPaymentKeeper)(nil).GetStreamRecord), arg0, arg1)
}

// IsPaymentAccountOwner mocks base method
func (m *MockPaymentKeeper) IsPaymentAccountOwner(arg0 types2.Context, arg1, arg2 types2.AccAddress) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPaymentAccountOwner", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsPaymentAccountOwner indicates an expected call of IsPaymentAccountOwner
func (mr *MockPaymentKeeperMockRecorder) IsPaymentAccountOwner(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPaymentAccountOwner", reflect.TypeOf((*MockPaymentKeeper)(nil).IsPaymentAccountOwner), arg0, arg1, arg2)
}

// UpdateStreamRecordByAddr mocks base method
func (m *MockPaymentKeeper) UpdateStreamRecordByAddr(arg0 types2.Context, arg1 *types.StreamRecordChange) (*types.StreamRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStreamRecordByAddr", arg0, arg1)
	ret0, _ := ret[0].(*types.StreamRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStreamRecordByAddr indicates an expected call of UpdateStreamRecordByAddr
func (mr *MockPaymentKeeperMockRecorder) UpdateStreamRecordByAddr(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStreamRecordByAddr", reflect.TypeOf((*MockPaymentKeeper)(nil).UpdateStreamRecordByAddr), arg0, arg1)
}

// MockPermissionKeeper is a mock of PermissionKeeper interface
type MockPermissionKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockPermissionKeeperMockRecorder
}

// MockPermissionKeeperMockRecorder is the mock recorder for MockPermissionKeeper
type MockPermissionKeeperMockRecorder struct {
	mock *MockPermissionKeeper
}

// NewMockPermissionKeeper creates a new mock instance
func NewMockPermissionKeeper(ctrl *gomock.Controller) *MockPermissionKeeper {
	mock := &MockPermissionKeeper{ctrl: ctrl}
	mock.recorder = &MockPermissionKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPermissionKeeper) EXPECT() *MockPermissionKeeperMockRecorder {
	return m.recorder
}

// AddGroupMember mocks base method
func (m *MockPermissionKeeper) AddGroupMember(arg0 types2.Context, arg1 math.Uint, arg2 types2.AccAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddGroupMember", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddGroupMember indicates an expected call of AddGroupMember
func (mr *MockPermissionKeeperMockRecorder) AddGroupMember(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddGroupMember", reflect.TypeOf((*MockPermissionKeeper)(nil).AddGroupMember), arg0, arg1, arg2)
}

// DeletePolicy mocks base method
func (m *MockPermissionKeeper) DeletePolicy(arg0 types2.Context, arg1 *types0.Principal, arg2 resource.ResourceType, arg3 math.Uint) (math.Uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePolicy", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(math.Uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePolicy indicates an expected call of DeletePolicy
func (mr *MockPermissionKeeperMockRecorder) DeletePolicy(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePolicy", reflect.TypeOf((*MockPermissionKeeper)(nil).DeletePolicy), arg0, arg1, arg2, arg3)
}

// GetGroupMember mocks base method
func (m *MockPermissionKeeper) GetGroupMember(arg0 types2.Context, arg1 math.Uint, arg2 types2.AccAddress) (*types0.GroupMember, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupMember", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types0.GroupMember)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetGroupMember indicates an expected call of GetGroupMember
func (mr *MockPermissionKeeperMockRecorder) GetGroupMember(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupMember", reflect.TypeOf((*MockPermissionKeeper)(nil).GetGroupMember), arg0, arg1, arg2)
}

// GetGroupMemberByID mocks base method
func (m *MockPermissionKeeper) GetGroupMemberByID(arg0 types2.Context, arg1 math.Uint) (*types0.GroupMember, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupMemberByID", arg0, arg1)
	ret0, _ := ret[0].(*types0.GroupMember)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetGroupMemberByID indicates an expected call of GetGroupMemberByID
func (mr *MockPermissionKeeperMockRecorder) GetGroupMemberByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupMemberByID", reflect.TypeOf((*MockPermissionKeeper)(nil).GetGroupMemberByID), arg0, arg1)
}

// GetPolicyByID mocks base method
func (m *MockPermissionKeeper) GetPolicyByID(arg0 types2.Context, arg1 math.Uint) (*types0.Policy, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPolicyByID", arg0, arg1)
	ret0, _ := ret[0].(*types0.Policy)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetPolicyByID indicates an expected call of GetPolicyByID
func (mr *MockPermissionKeeperMockRecorder) GetPolicyByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicyByID", reflect.TypeOf((*MockPermissionKeeper)(nil).GetPolicyByID), arg0, arg1)
}

// GetPolicyForAccount mocks base method
func (m *MockPermissionKeeper) GetPolicyForAccount(arg0 types2.Context, arg1 math.Uint, arg2 resource.ResourceType, arg3 types2.AccAddress) (*types0.Policy, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPolicyForAccount", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*types0.Policy)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetPolicyForAccount indicates an expected call of GetPolicyForAccount
func (mr *MockPermissionKeeperMockRecorder) GetPolicyForAccount(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicyForAccount", reflect.TypeOf((*MockPermissionKeeper)(nil).GetPolicyForAccount), arg0, arg1, arg2, arg3)
}

// GetPolicyForGroup mocks base method
func (m *MockPermissionKeeper) GetPolicyForGroup(arg0 types2.Context, arg1 math.Uint, arg2 resource.ResourceType, arg3 math.Uint) (*types0.Policy, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPolicyForGroup", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*types0.Policy)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetPolicyForGroup indicates an expected call of GetPolicyForGroup
func (mr *MockPermissionKeeperMockRecorder) GetPolicyForGroup(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicyForGroup", reflect.TypeOf((*MockPermissionKeeper)(nil).GetPolicyForGroup), arg0, arg1, arg2, arg3)
}

// PutPolicy mocks base method
func (m *MockPermissionKeeper) PutPolicy(arg0 types2.Context, arg1 *types0.Policy) (math.Uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutPolicy", arg0, arg1)
	ret0, _ := ret[0].(math.Uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutPolicy indicates an expected call of PutPolicy
func (mr *MockPermissionKeeperMockRecorder) PutPolicy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutPolicy", reflect.TypeOf((*MockPermissionKeeper)(nil).PutPolicy), arg0, arg1)
}

// RemoveGroupMember mocks base method
func (m *MockPermissionKeeper) RemoveGroupMember(arg0 types2.Context, arg1 math.Uint, arg2 types2.AccAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveGroupMember", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveGroupMember indicates an expected call of RemoveGroupMember
func (mr *MockPermissionKeeperMockRecorder) RemoveGroupMember(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveGroupMember", reflect.TypeOf((*MockPermissionKeeper)(nil).RemoveGroupMember), arg0, arg1, arg2)
}

// VerifyPolicy mocks base method
func (m *MockPermissionKeeper) VerifyPolicy(arg0 types2.Context, arg1 math.Uint, arg2 resource.ResourceType, arg3 types2.AccAddress, arg4 types0.ActionType, arg5 *types0.VerifyOptions) types0.Effect {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyPolicy", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(types0.Effect)
	return ret0
}

// VerifyPolicy indicates an expected call of VerifyPolicy
func (mr *MockPermissionKeeperMockRecorder) VerifyPolicy(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyPolicy", reflect.TypeOf((*MockPermissionKeeper)(nil).VerifyPolicy), arg0, arg1, arg2, arg3, arg4, arg5)
}

// MockCrossChainKeeper is a mock of CrossChainKeeper interface
type MockCrossChainKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockCrossChainKeeperMockRecorder
}

// MockCrossChainKeeperMockRecorder is the mock recorder for MockCrossChainKeeper
type MockCrossChainKeeperMockRecorder struct {
	mock *MockCrossChainKeeper
}

// NewMockCrossChainKeeper creates a new mock instance
func NewMockCrossChainKeeper(ctrl *gomock.Controller) *MockCrossChainKeeper {
	mock := &MockCrossChainKeeper{ctrl: ctrl}
	mock.recorder = &MockCrossChainKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCrossChainKeeper) EXPECT() *MockCrossChainKeeperMockRecorder {
	return m.recorder
}

// CreateRawIBCPackageWithFee mocks base method
func (m *MockCrossChainKeeper) CreateRawIBCPackageWithFee(arg0 types2.Context, arg1 types2.ChannelID, arg2 types2.CrossChainPackageType, arg3 []byte, arg4, arg5 *big.Int) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRawIBCPackageWithFee", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRawIBCPackageWithFee indicates an expected call of CreateRawIBCPackageWithFee
func (mr *MockCrossChainKeeperMockRecorder) CreateRawIBCPackageWithFee(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRawIBCPackageWithFee", reflect.TypeOf((*MockCrossChainKeeper)(nil).CreateRawIBCPackageWithFee), arg0, arg1, arg2, arg3, arg4, arg5)
}

// RegisterChannel mocks base method
func (m *MockCrossChainKeeper) RegisterChannel(arg0 string, arg1 types2.ChannelID, arg2 types2.CrossChainApplication) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterChannel", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterChannel indicates an expected call of RegisterChannel
func (mr *MockCrossChainKeeperMockRecorder) RegisterChannel(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterChannel", reflect.TypeOf((*MockCrossChainKeeper)(nil).RegisterChannel), arg0, arg1, arg2)
}

// ForceDeleteAccountPolicyForResource mocks base method.
func (m *MockPermissionKeeper) ForceDeleteAccountPolicyForResource(ctx types2.Context, maxDelete, deletedCount uint64, resourceType resource.ResourceType, resourceID math.Uint) (uint64, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForceDeleteAccountPolicyForResource", ctx, resourceType, resourceID)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ForceDeleteAccountPolicyForResource indicates an expected call of ForceDeleteAccountPolicyForResource.
func (mr *MockPermissionKeeperMockRecorder) ForceDeleteAccountPolicyForResource(ctx, maxDelete, deletedCount, resourceType, resourceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForceDeleteAccountPolicyForResource", reflect.TypeOf((*MockPermissionKeeper)(nil).ForceDeleteAccountPolicyForResource), ctx, maxDelete, deletedCount, resourceType, resourceID)
}

// ForceDeleteGroupPolicyForResource mocks base method.
func (m *MockPermissionKeeper) ForceDeleteGroupPolicyForResource(ctx types2.Context, maxDelete, deletedCount uint64, resourceType resource.ResourceType, resourceID math.Uint) (uint64, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForceDeleteGroupPolicyForResource", ctx, resourceType, resourceID)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ForceDeleteGroupPolicyForResource indicates an expected call of ForceDeleteGroupPolicyForResource.
func (mr *MockPermissionKeeperMockRecorder) ForceDeleteGroupPolicyForResource(ctx, maxDelete, deletedCount, resourceType, resourceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForceDeleteGroupPolicyForResource", reflect.TypeOf((*MockPermissionKeeper)(nil).ForceDeleteGroupPolicyForResource), ctx, maxDelete, deletedCount, resourceType, resourceID)
}

// ForceDeleteGroupMembers mocks base method.
func (m *MockPermissionKeeper) ForceDeleteGroupMembers(ctx types2.Context, maxDelete, deletedCount uint64, groupId math.Uint) (uint64, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForceDeleteGroupMembers", ctx, groupId)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ForceDeleteGroupMembers indicates an expected call of ForceDeleteGroupMembers.
func (mr *MockPermissionKeeperMockRecorder) ForceDeleteGroupMembers(ctx, maxDelete, deletedCount, groupId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForceDeleteGroupMembers", reflect.TypeOf((*MockPermissionKeeper)(nil).ForceDeleteGroupMembers), ctx, maxDelete, deletedCount, groupId)
}

// ExistAccountPolicyForResource mocks base method.
func (m *MockPermissionKeeper) ExistAccountPolicyForResource(ctx types2.Context, resourceType resource.ResourceType, resourceID math.Uint) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistAccountPolicyForResource", ctx, resourceType, resourceID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ExistAccountPolicyForResource indicates an expected call of ExistAccountPolicyForResource.
func (mr *MockPermissionKeeperMockRecorder) ExistAccountPolicyForResource(ctx, resourceType, resourceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistAccountPolicyForResource", reflect.TypeOf((*MockPermissionKeeper)(nil).ExistAccountPolicyForResource), ctx, resourceType, resourceID)
}

// ExistGroupPolicyForResource mocks base method.
func (m *MockPermissionKeeper) ExistGroupPolicyForResource(ctx types2.Context, resourceType resource.ResourceType, resourceID math.Uint) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistGroupPolicyForResource", ctx, resourceType, resourceID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ExistGroupPolicyForResource indicates an expected call of ExistGroupPolicyForResource.
func (mr *MockPermissionKeeperMockRecorder) ExistGroupPolicyForResource(ctx, resourceType, resourceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistGroupPolicyForResource", reflect.TypeOf((*MockPermissionKeeper)(nil).ExistGroupPolicyForResource), ctx, resourceType, resourceID)
}

// ExistGroupMemberForGroup mocks base method.
func (m *MockPermissionKeeper) ExistGroupMemberForGroup(ctx types2.Context, groupID math.Uint) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistGroupMemberForGroup", ctx, groupID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ExistGroupMemberForGroup indicates an expected call of ExistGroupMemberForGroup.
func (mr *MockPermissionKeeperMockRecorder) ExistGroupMemberForGroup(ctx, groupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistGroupMemberForGroup", reflect.TypeOf((*MockPermissionKeeper)(nil).ExistGroupMemberForGroup), ctx, groupID)
}