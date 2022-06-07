// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated by MockGen. DO NOT EDIT.
// Source: engine/workflow/installer.go

// Package workflow is a generated GoMock package.
package workflow

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockInstaller is a mock of Installer interface.
type MockInstaller struct {
	ctrl     *gomock.Controller
	recorder *MockInstallerMockRecorder
}

// MockInstallerMockRecorder is the mock recorder for MockInstaller.
type MockInstallerMockRecorder struct {
	mock *MockInstaller
}

// NewMockInstaller creates a new mock instance.
func NewMockInstaller(ctrl *gomock.Controller) *MockInstaller {
	mock := &MockInstaller{ctrl: ctrl}
	mock.recorder = &MockInstallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInstaller) EXPECT() *MockInstallerMockRecorder {
	return m.recorder
}

// Decompress mocks base method.
func (m *MockInstaller) Decompress(source, dest string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decompress", source, dest)
	ret0, _ := ret[0].(error)
	return ret0
}

// Decompress indicates an expected call of Decompress.
func (mr *MockInstallerMockRecorder) Decompress(source, dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decompress", reflect.TypeOf((*MockInstaller)(nil).Decompress), source, dest)
}

// Download mocks base method.
func (m *MockInstaller) Download(url, path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", url, path)
	ret0, _ := ret[0].(error)
	return ret0
}

// Download indicates an expected call of Download.
func (mr *MockInstallerMockRecorder) Download(url, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockInstaller)(nil).Download), url, path)
}

// Install mocks base method.
func (m *MockInstaller) Install(workingDir, installScriptPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Install", workingDir, installScriptPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// Install indicates an expected call of Install.
func (mr *MockInstallerMockRecorder) Install(workingDir, installScriptPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockInstaller)(nil).Install), workingDir, installScriptPath)
}