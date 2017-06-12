// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/sclevine/cflocal/cf/cmd (interfaces: FS)

package mocks

import (
	gomock "github.com/golang/mock/gomock"
	fs "github.com/sclevine/cflocal/fs"
	io "io"
	time "time"
)

// Mock of FS interface
type MockFS struct {
	ctrl     *gomock.Controller
	recorder *_MockFSRecorder
}

// Recorder for MockFS (not exported)
type _MockFSRecorder struct {
	mock *MockFS
}

func NewMockFS(ctrl *gomock.Controller) *MockFS {
	mock := &MockFS{ctrl: ctrl}
	mock.recorder = &_MockFSRecorder{mock}
	return mock
}

func (_m *MockFS) EXPECT() *_MockFSRecorder {
	return _m.recorder
}

func (_m *MockFS) Abs(_param0 string) (string, error) {
	ret := _m.ctrl.Call(_m, "Abs", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFSRecorder) Abs(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Abs", arg0)
}

func (_m *MockFS) MakeDirAll(_param0 string) error {
	ret := _m.ctrl.Call(_m, "MakeDirAll", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockFSRecorder) MakeDirAll(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MakeDirAll", arg0)
}

func (_m *MockFS) OpenFile(_param0 string) (fs.ReadResetWriteCloser, int64, error) {
	ret := _m.ctrl.Call(_m, "OpenFile", _param0)
	ret0, _ := ret[0].(fs.ReadResetWriteCloser)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockFSRecorder) OpenFile(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "OpenFile", arg0)
}

func (_m *MockFS) ReadFile(_param0 string) (io.ReadCloser, int64, error) {
	ret := _m.ctrl.Call(_m, "ReadFile", _param0)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockFSRecorder) ReadFile(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadFile", arg0)
}

func (_m *MockFS) TarApp(_param0 string) (io.ReadCloser, error) {
	ret := _m.ctrl.Call(_m, "TarApp", _param0)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFSRecorder) TarApp(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TarApp", arg0)
}

func (_m *MockFS) Watch(_param0 string, _param1 time.Duration) (<-chan time.Time, chan<- struct{}, error) {
	ret := _m.ctrl.Call(_m, "Watch", _param0, _param1)
	ret0, _ := ret[0].(<-chan time.Time)
	ret1, _ := ret[1].(chan<- struct{})
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockFSRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Watch", arg0, arg1)
}

func (_m *MockFS) WriteFile(_param0 string) (io.WriteCloser, error) {
	ret := _m.ctrl.Call(_m, "WriteFile", _param0)
	ret0, _ := ret[0].(io.WriteCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFSRecorder) WriteFile(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WriteFile", arg0)
}
