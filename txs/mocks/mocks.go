// Code generated by MockGen. DO NOT EDIT.
// Source: ./interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	types "github.com/spacemeshos/go-spacemesh/common/types"
	system "github.com/spacemeshos/go-spacemesh/system"
	types0 "github.com/spacemeshos/go-spacemesh/txs/types"
)

// MockconservativeState is a mock of conservativeState interface.
type MockconservativeState struct {
	ctrl     *gomock.Controller
	recorder *MockconservativeStateMockRecorder
}

// MockconservativeStateMockRecorder is the mock recorder for MockconservativeState.
type MockconservativeStateMockRecorder struct {
	mock *MockconservativeState
}

// NewMockconservativeState creates a new mock instance.
func NewMockconservativeState(ctrl *gomock.Controller) *MockconservativeState {
	mock := &MockconservativeState{ctrl: ctrl}
	mock.recorder = &MockconservativeStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockconservativeState) EXPECT() *MockconservativeStateMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockconservativeState) Add(arg0 *types.Transaction, arg1 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockconservativeStateMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockconservativeState)(nil).Add), arg0, arg1)
}

// AddToCache mocks base method.
func (m *MockconservativeState) AddToCache(arg0 *types.Transaction, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToCache", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToCache indicates an expected call of AddToCache.
func (mr *MockconservativeStateMockRecorder) AddToCache(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToCache", reflect.TypeOf((*MockconservativeState)(nil).AddToCache), arg0, arg1)
}

// HasTx mocks base method.
func (m *MockconservativeState) HasTx(arg0 types.TransactionID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasTx", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasTx indicates an expected call of HasTx.
func (mr *MockconservativeStateMockRecorder) HasTx(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasTx", reflect.TypeOf((*MockconservativeState)(nil).HasTx), arg0)
}

// Validation mocks base method.
func (m *MockconservativeState) Validation(arg0 types.RawTx) system.ValidationRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validation", arg0)
	ret0, _ := ret[0].(system.ValidationRequest)
	return ret0
}

// Validation indicates an expected call of Validation.
func (mr *MockconservativeStateMockRecorder) Validation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validation", reflect.TypeOf((*MockconservativeState)(nil).Validation), arg0)
}

// MockvmState is a mock of vmState interface.
type MockvmState struct {
	ctrl     *gomock.Controller
	recorder *MockvmStateMockRecorder
}

// MockvmStateMockRecorder is the mock recorder for MockvmState.
type MockvmStateMockRecorder struct {
	mock *MockvmState
}

// NewMockvmState creates a new mock instance.
func NewMockvmState(ctrl *gomock.Controller) *MockvmState {
	mock := &MockvmState{ctrl: ctrl}
	mock.recorder = &MockvmStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockvmState) EXPECT() *MockvmStateMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockvmState) Apply(arg0 types.LayerID, arg1 []types.RawTx, arg2 []types.AnyReward) ([]types.TransactionID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1, arg2)
	ret0, _ := ret[0].([]types.TransactionID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockvmStateMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockvmState)(nil).Apply), arg0, arg1, arg2)
}

// GetAllAccounts mocks base method.
func (m *MockvmState) GetAllAccounts() ([]*types.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAccounts")
	ret0, _ := ret[0].([]*types.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAccounts indicates an expected call of GetAllAccounts.
func (mr *MockvmStateMockRecorder) GetAllAccounts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAccounts", reflect.TypeOf((*MockvmState)(nil).GetAllAccounts))
}

// GetBalance mocks base method.
func (m *MockvmState) GetBalance(arg0 types.Address) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockvmStateMockRecorder) GetBalance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockvmState)(nil).GetBalance), arg0)
}

// GetLayerApplied mocks base method.
func (m *MockvmState) GetLayerApplied(arg0 types.TransactionID) (types.LayerID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLayerApplied", arg0)
	ret0, _ := ret[0].(types.LayerID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLayerApplied indicates an expected call of GetLayerApplied.
func (mr *MockvmStateMockRecorder) GetLayerApplied(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLayerApplied", reflect.TypeOf((*MockvmState)(nil).GetLayerApplied), arg0)
}

// GetLayerStateRoot mocks base method.
func (m *MockvmState) GetLayerStateRoot(arg0 types.LayerID) (types.Hash32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLayerStateRoot", arg0)
	ret0, _ := ret[0].(types.Hash32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLayerStateRoot indicates an expected call of GetLayerStateRoot.
func (mr *MockvmStateMockRecorder) GetLayerStateRoot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLayerStateRoot", reflect.TypeOf((*MockvmState)(nil).GetLayerStateRoot), arg0)
}

// GetNonce mocks base method.
func (m *MockvmState) GetNonce(arg0 types.Address) (types.Nonce, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNonce", arg0)
	ret0, _ := ret[0].(types.Nonce)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNonce indicates an expected call of GetNonce.
func (mr *MockvmStateMockRecorder) GetNonce(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNonce", reflect.TypeOf((*MockvmState)(nil).GetNonce), arg0)
}

// GetStateRoot mocks base method.
func (m *MockvmState) GetStateRoot() (types.Hash32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStateRoot")
	ret0, _ := ret[0].(types.Hash32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStateRoot indicates an expected call of GetStateRoot.
func (mr *MockvmStateMockRecorder) GetStateRoot() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStateRoot", reflect.TypeOf((*MockvmState)(nil).GetStateRoot))
}

// Revert mocks base method.
func (m *MockvmState) Revert(arg0 types.LayerID) (types.Hash32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Revert", arg0)
	ret0, _ := ret[0].(types.Hash32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Revert indicates an expected call of Revert.
func (mr *MockvmStateMockRecorder) Revert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Revert", reflect.TypeOf((*MockvmState)(nil).Revert), arg0)
}

// Validation mocks base method.
func (m *MockvmState) Validation(arg0 types.RawTx) system.ValidationRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validation", arg0)
	ret0, _ := ret[0].(system.ValidationRequest)
	return ret0
}

// Validation indicates an expected call of Validation.
func (mr *MockvmStateMockRecorder) Validation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validation", reflect.TypeOf((*MockvmState)(nil).Validation), arg0)
}

// MockconStateCache is a mock of conStateCache interface.
type MockconStateCache struct {
	ctrl     *gomock.Controller
	recorder *MockconStateCacheMockRecorder
}

// MockconStateCacheMockRecorder is the mock recorder for MockconStateCache.
type MockconStateCacheMockRecorder struct {
	mock *MockconStateCache
}

// NewMockconStateCache creates a new mock instance.
func NewMockconStateCache(ctrl *gomock.Controller) *MockconStateCache {
	mock := &MockconStateCache{ctrl: ctrl}
	mock.recorder = &MockconStateCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockconStateCache) EXPECT() *MockconStateCacheMockRecorder {
	return m.recorder
}

// GetMempool mocks base method.
func (m *MockconStateCache) GetMempool() map[types.Address][]*types0.NanoTX {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMempool")
	ret0, _ := ret[0].(map[types.Address][]*types0.NanoTX)
	return ret0
}

// GetMempool indicates an expected call of GetMempool.
func (mr *MockconStateCacheMockRecorder) GetMempool() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMempool", reflect.TypeOf((*MockconStateCache)(nil).GetMempool))
}

// MocktxProvider is a mock of txProvider interface.
type MocktxProvider struct {
	ctrl     *gomock.Controller
	recorder *MocktxProviderMockRecorder
}

// MocktxProviderMockRecorder is the mock recorder for MocktxProvider.
type MocktxProviderMockRecorder struct {
	mock *MocktxProvider
}

// NewMocktxProvider creates a new mock instance.
func NewMocktxProvider(ctrl *gomock.Controller) *MocktxProvider {
	mock := &MocktxProvider{ctrl: ctrl}
	mock.recorder = &MocktxProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocktxProvider) EXPECT() *MocktxProviderMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MocktxProvider) Add(arg0 *types.Transaction, arg1 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MocktxProviderMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MocktxProvider)(nil).Add), arg0, arg1)
}

// AddToBlock mocks base method.
func (m *MocktxProvider) AddToBlock(arg0 types.LayerID, arg1 types.BlockID, arg2 []types.TransactionID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToBlock", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToBlock indicates an expected call of AddToBlock.
func (mr *MocktxProviderMockRecorder) AddToBlock(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToBlock", reflect.TypeOf((*MocktxProvider)(nil).AddToBlock), arg0, arg1, arg2)
}

// AddToProposal mocks base method.
func (m *MocktxProvider) AddToProposal(arg0 types.LayerID, arg1 types.ProposalID, arg2 []types.TransactionID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToProposal", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToProposal indicates an expected call of AddToProposal.
func (mr *MocktxProviderMockRecorder) AddToProposal(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToProposal", reflect.TypeOf((*MocktxProvider)(nil).AddToProposal), arg0, arg1, arg2)
}

// ApplyLayer mocks base method.
func (m *MocktxProvider) ApplyLayer(arg0 types.LayerID, arg1 types.BlockID, arg2 types.Address, arg3 map[uint64]types.TransactionID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyLayer", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyLayer indicates an expected call of ApplyLayer.
func (mr *MocktxProviderMockRecorder) ApplyLayer(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyLayer", reflect.TypeOf((*MocktxProvider)(nil).ApplyLayer), arg0, arg1, arg2, arg3)
}

// DiscardNonceBelow mocks base method.
func (m *MocktxProvider) DiscardNonceBelow(arg0 types.Address, arg1 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiscardNonceBelow", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DiscardNonceBelow indicates an expected call of DiscardNonceBelow.
func (mr *MocktxProviderMockRecorder) DiscardNonceBelow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscardNonceBelow", reflect.TypeOf((*MocktxProvider)(nil).DiscardNonceBelow), arg0, arg1)
}

// Get mocks base method.
func (m *MocktxProvider) Get(arg0 types.TransactionID) (*types.MeshTransaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*types.MeshTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MocktxProviderMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MocktxProvider)(nil).Get), arg0)
}

// GetAcctPendingFromNonce mocks base method.
func (m *MocktxProvider) GetAcctPendingFromNonce(arg0 types.Address, arg1 uint64) ([]*types.MeshTransaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAcctPendingFromNonce", arg0, arg1)
	ret0, _ := ret[0].([]*types.MeshTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAcctPendingFromNonce indicates an expected call of GetAcctPendingFromNonce.
func (mr *MocktxProviderMockRecorder) GetAcctPendingFromNonce(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAcctPendingFromNonce", reflect.TypeOf((*MocktxProvider)(nil).GetAcctPendingFromNonce), arg0, arg1)
}

// GetAllPending mocks base method.
func (m *MocktxProvider) GetAllPending() ([]*types.MeshTransaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPending")
	ret0, _ := ret[0].([]*types.MeshTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPending indicates an expected call of GetAllPending.
func (mr *MocktxProviderMockRecorder) GetAllPending() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPending", reflect.TypeOf((*MocktxProvider)(nil).GetAllPending))
}

// GetBlob mocks base method.
func (m *MocktxProvider) GetBlob(arg0 types.TransactionID) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlob", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlob indicates an expected call of GetBlob.
func (mr *MocktxProviderMockRecorder) GetBlob(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlob", reflect.TypeOf((*MocktxProvider)(nil).GetBlob), arg0)
}

// GetByAddress mocks base method.
func (m *MocktxProvider) GetByAddress(arg0, arg1 types.LayerID, arg2 types.Address) ([]*types.MeshTransaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByAddress", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*types.MeshTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByAddress indicates an expected call of GetByAddress.
func (mr *MocktxProviderMockRecorder) GetByAddress(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByAddress", reflect.TypeOf((*MocktxProvider)(nil).GetByAddress), arg0, arg1, arg2)
}

// GetMeshHash mocks base method.
func (m *MocktxProvider) GetMeshHash(arg0 types.LayerID) (types.Hash32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMeshHash", arg0)
	ret0, _ := ret[0].(types.Hash32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeshHash indicates an expected call of GetMeshHash.
func (mr *MocktxProviderMockRecorder) GetMeshHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeshHash", reflect.TypeOf((*MocktxProvider)(nil).GetMeshHash), arg0)
}

// Has mocks base method.
func (m *MocktxProvider) Has(arg0 types.TransactionID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Has indicates an expected call of Has.
func (mr *MocktxProviderMockRecorder) Has(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MocktxProvider)(nil).Has), arg0)
}

// LastAppliedLayer mocks base method.
func (m *MocktxProvider) LastAppliedLayer() (types.LayerID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastAppliedLayer")
	ret0, _ := ret[0].(types.LayerID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastAppliedLayer indicates an expected call of LastAppliedLayer.
func (mr *MocktxProviderMockRecorder) LastAppliedLayer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastAppliedLayer", reflect.TypeOf((*MocktxProvider)(nil).LastAppliedLayer))
}

// SetNextLayerBlock mocks base method.
func (m *MocktxProvider) SetNextLayerBlock(arg0 types.TransactionID, arg1 types.LayerID) (types.LayerID, types.BlockID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetNextLayerBlock", arg0, arg1)
	ret0, _ := ret[0].(types.LayerID)
	ret1, _ := ret[1].(types.BlockID)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SetNextLayerBlock indicates an expected call of SetNextLayerBlock.
func (mr *MocktxProviderMockRecorder) SetNextLayerBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNextLayerBlock", reflect.TypeOf((*MocktxProvider)(nil).SetNextLayerBlock), arg0, arg1)
}

// UndoLayers mocks base method.
func (m *MocktxProvider) UndoLayers(arg0 types.LayerID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UndoLayers", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UndoLayers indicates an expected call of UndoLayers.
func (mr *MocktxProviderMockRecorder) UndoLayers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UndoLayers", reflect.TypeOf((*MocktxProvider)(nil).UndoLayers), arg0)
}
