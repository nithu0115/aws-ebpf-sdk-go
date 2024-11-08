// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-ebpf-sdk-go/pkg/progs (interfaces: BpfProgAPIs)

// Package mock_progs is a generated GoMock package.
package mock_progs

import (
	reflect "reflect"

	progs "github.com/aws/aws-ebpf-sdk-go/pkg/progs"
	gomock "github.com/golang/mock/gomock"
)

// MockBpfProgAPIs is a mock of BpfProgAPIs interface.
type MockBpfProgAPIs struct {
	ctrl     *gomock.Controller
	recorder *MockBpfProgAPIsMockRecorder
}

// MockBpfProgAPIsMockRecorder is the mock recorder for MockBpfProgAPIs.
type MockBpfProgAPIsMockRecorder struct {
	mock *MockBpfProgAPIs
}

// NewMockBpfProgAPIs creates a new mock instance.
func NewMockBpfProgAPIs(ctrl *gomock.Controller) *MockBpfProgAPIs {
	mock := &MockBpfProgAPIs{ctrl: ctrl}
	mock.recorder = &MockBpfProgAPIsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBpfProgAPIs) EXPECT() *MockBpfProgAPIsMockRecorder {
	return m.recorder
}

// GetBPFProgAssociatedMapsIDs mocks base method.
func (m *MockBpfProgAPIs) GetBPFProgAssociatedMapsIDs(arg0 int) ([]uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBPFProgAssociatedMapsIDs", arg0)
	ret0, _ := ret[0].([]uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBPFProgAssociatedMapsIDs indicates an expected call of GetBPFProgAssociatedMapsIDs.
func (mr *MockBpfProgAPIsMockRecorder) GetBPFProgAssociatedMapsIDs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBPFProgAssociatedMapsIDs", reflect.TypeOf((*MockBpfProgAPIs)(nil).GetBPFProgAssociatedMapsIDs), arg0)
}

// GetProgFromPinPath mocks base method.
func (m *MockBpfProgAPIs) GetProgFromPinPath(arg0 string) (progs.BpfProgInfo, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProgFromPinPath", arg0)
	ret0, _ := ret[0].(progs.BpfProgInfo)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetProgFromPinPath indicates an expected call of GetProgFromPinPath.
func (mr *MockBpfProgAPIsMockRecorder) GetProgFromPinPath(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProgFromPinPath", reflect.TypeOf((*MockBpfProgAPIs)(nil).GetProgFromPinPath), arg0)
}

// LoadProg mocks base method.
func (m *MockBpfProgAPIs) LoadProg(arg0 progs.CreateEBPFProgInput) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadProg", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadProg indicates an expected call of LoadProg.
func (mr *MockBpfProgAPIsMockRecorder) LoadProg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadProg", reflect.TypeOf((*MockBpfProgAPIs)(nil).LoadProg), arg0)
}

// LoadProgWithNonePinType mocks base method.
func (m *MockBpfProgAPIs) LoadProgWithNonePinType(arg0 progs.CreateEBPFProgInput, arg1 uint32) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadProgWithNonePinType", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadProgWithNonePinType indicates an expected call of LoadProgWithNonePinType.
func (mr *MockBpfProgAPIsMockRecorder) LoadProgWithNonePinType(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadProgWithNonePinType", reflect.TypeOf((*MockBpfProgAPIs)(nil).LoadProgWithNonePinType), arg0, arg1)
}

// PinProg mocks base method.
func (m *MockBpfProgAPIs) PinProg(arg0 uint32, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PinProg", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PinProg indicates an expected call of PinProg.
func (mr *MockBpfProgAPIsMockRecorder) PinProg(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PinProg", reflect.TypeOf((*MockBpfProgAPIs)(nil).PinProg), arg0, arg1)
}

// UnPinProg mocks base method.
func (m *MockBpfProgAPIs) UnPinProg(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnPinProg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnPinProg indicates an expected call of UnPinProg.
func (mr *MockBpfProgAPIsMockRecorder) UnPinProg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnPinProg", reflect.TypeOf((*MockBpfProgAPIs)(nil).UnPinProg), arg0)
}
