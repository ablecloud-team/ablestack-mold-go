//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: ./cloudstack/HostService.go

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHostServiceIface is a mock of HostServiceIface interface.
type MockHostServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockHostServiceIfaceMockRecorder
}

// MockHostServiceIfaceMockRecorder is the mock recorder for MockHostServiceIface.
type MockHostServiceIfaceMockRecorder struct {
	mock *MockHostServiceIface
}

// NewMockHostServiceIface creates a new mock instance.
func NewMockHostServiceIface(ctrl *gomock.Controller) *MockHostServiceIface {
	mock := &MockHostServiceIface{ctrl: ctrl}
	mock.recorder = &MockHostServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHostServiceIface) EXPECT() *MockHostServiceIfaceMockRecorder {
	return m.recorder
}

// AddBaremetalHost mocks base method.
func (m *MockHostServiceIface) AddBaremetalHost(p *AddBaremetalHostParams) (*AddBaremetalHostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBaremetalHost", p)
	ret0, _ := ret[0].(*AddBaremetalHostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddBaremetalHost indicates an expected call of AddBaremetalHost.
func (mr *MockHostServiceIfaceMockRecorder) AddBaremetalHost(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBaremetalHost", reflect.TypeOf((*MockHostServiceIface)(nil).AddBaremetalHost), p)
}

// AddGloboDnsHost mocks base method.
func (m *MockHostServiceIface) AddGloboDnsHost(p *AddGloboDnsHostParams) (*AddGloboDnsHostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddGloboDnsHost", p)
	ret0, _ := ret[0].(*AddGloboDnsHostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddGloboDnsHost indicates an expected call of AddGloboDnsHost.
func (mr *MockHostServiceIfaceMockRecorder) AddGloboDnsHost(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddGloboDnsHost", reflect.TypeOf((*MockHostServiceIface)(nil).AddGloboDnsHost), p)
}

// AddHost mocks base method.
func (m *MockHostServiceIface) AddHost(p *AddHostParams) (*AddHostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddHost", p)
	ret0, _ := ret[0].(*AddHostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddHost indicates an expected call of AddHost.
func (mr *MockHostServiceIfaceMockRecorder) AddHost(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddHost", reflect.TypeOf((*MockHostServiceIface)(nil).AddHost), p)
}

// AddSecondaryStorage mocks base method.
func (m *MockHostServiceIface) AddSecondaryStorage(p *AddSecondaryStorageParams) (*AddSecondaryStorageResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSecondaryStorage", p)
	ret0, _ := ret[0].(*AddSecondaryStorageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSecondaryStorage indicates an expected call of AddSecondaryStorage.
func (mr *MockHostServiceIfaceMockRecorder) AddSecondaryStorage(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSecondaryStorage", reflect.TypeOf((*MockHostServiceIface)(nil).AddSecondaryStorage), p)
}

// CancelHostMaintenance mocks base method.
func (m *MockHostServiceIface) CancelHostMaintenance(p *CancelHostMaintenanceParams) (*CancelHostMaintenanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelHostMaintenance", p)
	ret0, _ := ret[0].(*CancelHostMaintenanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelHostMaintenance indicates an expected call of CancelHostMaintenance.
func (mr *MockHostServiceIfaceMockRecorder) CancelHostMaintenance(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelHostMaintenance", reflect.TypeOf((*MockHostServiceIface)(nil).CancelHostMaintenance), p)
}

// ConfigureHAForHost mocks base method.
func (m *MockHostServiceIface) ConfigureHAForHost(p *ConfigureHAForHostParams) (*HAForHostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigureHAForHost", p)
	ret0, _ := ret[0].(*HAForHostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigureHAForHost indicates an expected call of ConfigureHAForHost.
func (mr *MockHostServiceIfaceMockRecorder) ConfigureHAForHost(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigureHAForHost", reflect.TypeOf((*MockHostServiceIface)(nil).ConfigureHAForHost), p)
}

// DedicateHost mocks base method.
func (m *MockHostServiceIface) DedicateHost(p *DedicateHostParams) (*DedicateHostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DedicateHost", p)
	ret0, _ := ret[0].(*DedicateHostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DedicateHost indicates an expected call of DedicateHost.
func (mr *MockHostServiceIfaceMockRecorder) DedicateHost(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DedicateHost", reflect.TypeOf((*MockHostServiceIface)(nil).DedicateHost), p)
}

// DeleteHost mocks base method.
func (m *MockHostServiceIface) DeleteHost(p *DeleteHostParams) (*DeleteHostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteHost", p)
	ret0, _ := ret[0].(*DeleteHostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteHost indicates an expected call of DeleteHost.
func (mr *MockHostServiceIfaceMockRecorder) DeleteHost(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHost", reflect.TypeOf((*MockHostServiceIface)(nil).DeleteHost), p)
}

// DisableOutOfBandManagementForHost mocks base method.
func (m *MockHostServiceIface) DisableOutOfBandManagementForHost(p *DisableOutOfBandManagementForHostParams) (*DisableOutOfBandManagementForHostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableOutOfBandManagementForHost", p)
	ret0, _ := ret[0].(*DisableOutOfBandManagementForHostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableOutOfBandManagementForHost indicates an expected call of DisableOutOfBandManagementForHost.
func (mr *MockHostServiceIfaceMockRecorder) DisableOutOfBandManagementForHost(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableOutOfBandManagementForHost", reflect.TypeOf((*MockHostServiceIface)(nil).DisableOutOfBandManagementForHost), p)
}

// EnableHAForHost mocks base method.
func (m *MockHostServiceIface) EnableHAForHost(p *EnableHAForHostParams) (*EnableHAForHostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableHAForHost", p)
	ret0, _ := ret[0].(*EnableHAForHostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableHAForHost indicates an expected call of EnableHAForHost.
func (mr *MockHostServiceIfaceMockRecorder) EnableHAForHost(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableHAForHost", reflect.TypeOf((*MockHostServiceIface)(nil).EnableHAForHost), p)
}

// EnableOutOfBandManagementForHost mocks base method.
func (m *MockHostServiceIface) EnableOutOfBandManagementForHost(p *EnableOutOfBandManagementForHostParams) (*EnableOutOfBandManagementForHostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableOutOfBandManagementForHost", p)
	ret0, _ := ret[0].(*EnableOutOfBandManagementForHostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableOutOfBandManagementForHost indicates an expected call of EnableOutOfBandManagementForHost.
func (mr *MockHostServiceIfaceMockRecorder) EnableOutOfBandManagementForHost(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableOutOfBandManagementForHost", reflect.TypeOf((*MockHostServiceIface)(nil).EnableOutOfBandManagementForHost), p)
}

// FindHostsForMigration mocks base method.
func (m *MockHostServiceIface) FindHostsForMigration(p *FindHostsForMigrationParams) (*FindHostsForMigrationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindHostsForMigration", p)
	ret0, _ := ret[0].(*FindHostsForMigrationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindHostsForMigration indicates an expected call of FindHostsForMigration.
func (mr *MockHostServiceIfaceMockRecorder) FindHostsForMigration(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindHostsForMigration", reflect.TypeOf((*MockHostServiceIface)(nil).FindHostsForMigration), p)
}

// GetHostByID mocks base method.
func (m *MockHostServiceIface) GetHostByID(id string, opts ...OptionFunc) (*Host, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetHostByID", varargs...)
	ret0, _ := ret[0].(*Host)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetHostByID indicates an expected call of GetHostByID.
func (mr *MockHostServiceIfaceMockRecorder) GetHostByID(id interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostByID", reflect.TypeOf((*MockHostServiceIface)(nil).GetHostByID), varargs...)
}

// GetHostByName mocks base method.
func (m *MockHostServiceIface) GetHostByName(name string, opts ...OptionFunc) (*Host, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetHostByName", varargs...)
	ret0, _ := ret[0].(*Host)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetHostByName indicates an expected call of GetHostByName.
func (mr *MockHostServiceIfaceMockRecorder) GetHostByName(name interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostByName", reflect.TypeOf((*MockHostServiceIface)(nil).GetHostByName), varargs...)
}

// GetHostID mocks base method.
func (m *MockHostServiceIface) GetHostID(name string, opts ...OptionFunc) (string, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetHostID", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetHostID indicates an expected call of GetHostID.
func (mr *MockHostServiceIfaceMockRecorder) GetHostID(name interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostID", reflect.TypeOf((*MockHostServiceIface)(nil).GetHostID), varargs...)
}

// GetHostTagID mocks base method.
func (m *MockHostServiceIface) GetHostTagID(keyword string, opts ...OptionFunc) (string, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{keyword}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetHostTagID", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetHostTagID indicates an expected call of GetHostTagID.
func (mr *MockHostServiceIfaceMockRecorder) GetHostTagID(keyword interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{keyword}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostTagID", reflect.TypeOf((*MockHostServiceIface)(nil).GetHostTagID), varargs...)
}

// GetHostsMetricByID mocks base method.
func (m *MockHostServiceIface) GetHostsMetricByID(id string, opts ...OptionFunc) (*HostsMetric, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetHostsMetricByID", varargs...)
	ret0, _ := ret[0].(*HostsMetric)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetHostsMetricByID indicates an expected call of GetHostsMetricByID.
func (mr *MockHostServiceIfaceMockRecorder) GetHostsMetricByID(id interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostsMetricByID", reflect.TypeOf((*MockHostServiceIface)(nil).GetHostsMetricByID), varargs...)
}

// GetHostsMetricByName mocks base method.
func (m *MockHostServiceIface) GetHostsMetricByName(name string, opts ...OptionFunc) (*HostsMetric, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetHostsMetricByName", varargs...)
	ret0, _ := ret[0].(*HostsMetric)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetHostsMetricByName indicates an expected call of GetHostsMetricByName.
func (mr *MockHostServiceIfaceMockRecorder) GetHostsMetricByName(name interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostsMetricByName", reflect.TypeOf((*MockHostServiceIface)(nil).GetHostsMetricByName), varargs...)
}

// GetHostsMetricID mocks base method.
func (m *MockHostServiceIface) GetHostsMetricID(name string, opts ...OptionFunc) (string, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetHostsMetricID", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetHostsMetricID indicates an expected call of GetHostsMetricID.
func (mr *MockHostServiceIfaceMockRecorder) GetHostsMetricID(name interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostsMetricID", reflect.TypeOf((*MockHostServiceIface)(nil).GetHostsMetricID), varargs...)
}

// ListDedicatedHosts mocks base method.
func (m *MockHostServiceIface) ListDedicatedHosts(p *ListDedicatedHostsParams) (*ListDedicatedHostsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDedicatedHosts", p)
	ret0, _ := ret[0].(*ListDedicatedHostsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDedicatedHosts indicates an expected call of ListDedicatedHosts.
func (mr *MockHostServiceIfaceMockRecorder) ListDedicatedHosts(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDedicatedHosts", reflect.TypeOf((*MockHostServiceIface)(nil).ListDedicatedHosts), p)
}

// ListHostTags mocks base method.
func (m *MockHostServiceIface) ListHostTags(p *ListHostTagsParams) (*ListHostTagsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListHostTags", p)
	ret0, _ := ret[0].(*ListHostTagsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListHostTags indicates an expected call of ListHostTags.
func (mr *MockHostServiceIfaceMockRecorder) ListHostTags(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHostTags", reflect.TypeOf((*MockHostServiceIface)(nil).ListHostTags), p)
}

// ListHosts mocks base method.
func (m *MockHostServiceIface) ListHosts(p *ListHostsParams) (*ListHostsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListHosts", p)
	ret0, _ := ret[0].(*ListHostsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListHosts indicates an expected call of ListHosts.
func (mr *MockHostServiceIfaceMockRecorder) ListHosts(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHosts", reflect.TypeOf((*MockHostServiceIface)(nil).ListHosts), p)
}

// ListHostsMetrics mocks base method.
func (m *MockHostServiceIface) ListHostsMetrics(p *ListHostsMetricsParams) (*ListHostsMetricsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListHostsMetrics", p)
	ret0, _ := ret[0].(*ListHostsMetricsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListHostsMetrics indicates an expected call of ListHostsMetrics.
func (mr *MockHostServiceIfaceMockRecorder) ListHostsMetrics(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHostsMetrics", reflect.TypeOf((*MockHostServiceIface)(nil).ListHostsMetrics), p)
}

// NewAddBaremetalHostParams mocks base method.
func (m *MockHostServiceIface) NewAddBaremetalHostParams(hypervisor, password, podid, url, username, zoneid string) *AddBaremetalHostParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAddBaremetalHostParams", hypervisor, password, podid, url, username, zoneid)
	ret0, _ := ret[0].(*AddBaremetalHostParams)
	return ret0
}

// NewAddBaremetalHostParams indicates an expected call of NewAddBaremetalHostParams.
func (mr *MockHostServiceIfaceMockRecorder) NewAddBaremetalHostParams(hypervisor, password, podid, url, username, zoneid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAddBaremetalHostParams", reflect.TypeOf((*MockHostServiceIface)(nil).NewAddBaremetalHostParams), hypervisor, password, podid, url, username, zoneid)
}

// NewAddGloboDnsHostParams mocks base method.
func (m *MockHostServiceIface) NewAddGloboDnsHostParams(password, physicalnetworkid, url, username string) *AddGloboDnsHostParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAddGloboDnsHostParams", password, physicalnetworkid, url, username)
	ret0, _ := ret[0].(*AddGloboDnsHostParams)
	return ret0
}

// NewAddGloboDnsHostParams indicates an expected call of NewAddGloboDnsHostParams.
func (mr *MockHostServiceIfaceMockRecorder) NewAddGloboDnsHostParams(password, physicalnetworkid, url, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAddGloboDnsHostParams", reflect.TypeOf((*MockHostServiceIface)(nil).NewAddGloboDnsHostParams), password, physicalnetworkid, url, username)
}

// NewAddHostParams mocks base method.
func (m *MockHostServiceIface) NewAddHostParams(hypervisor, password, podid, url, username, zoneid string) *AddHostParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAddHostParams", hypervisor, password, podid, url, username, zoneid)
	ret0, _ := ret[0].(*AddHostParams)
	return ret0
}

// NewAddHostParams indicates an expected call of NewAddHostParams.
func (mr *MockHostServiceIfaceMockRecorder) NewAddHostParams(hypervisor, password, podid, url, username, zoneid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAddHostParams", reflect.TypeOf((*MockHostServiceIface)(nil).NewAddHostParams), hypervisor, password, podid, url, username, zoneid)
}

// NewAddSecondaryStorageParams mocks base method.
func (m *MockHostServiceIface) NewAddSecondaryStorageParams(url string) *AddSecondaryStorageParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAddSecondaryStorageParams", url)
	ret0, _ := ret[0].(*AddSecondaryStorageParams)
	return ret0
}

// NewAddSecondaryStorageParams indicates an expected call of NewAddSecondaryStorageParams.
func (mr *MockHostServiceIfaceMockRecorder) NewAddSecondaryStorageParams(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAddSecondaryStorageParams", reflect.TypeOf((*MockHostServiceIface)(nil).NewAddSecondaryStorageParams), url)
}

// NewCancelHostMaintenanceParams mocks base method.
func (m *MockHostServiceIface) NewCancelHostMaintenanceParams(id string) *CancelHostMaintenanceParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCancelHostMaintenanceParams", id)
	ret0, _ := ret[0].(*CancelHostMaintenanceParams)
	return ret0
}

// NewCancelHostMaintenanceParams indicates an expected call of NewCancelHostMaintenanceParams.
func (mr *MockHostServiceIfaceMockRecorder) NewCancelHostMaintenanceParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCancelHostMaintenanceParams", reflect.TypeOf((*MockHostServiceIface)(nil).NewCancelHostMaintenanceParams), id)
}

// NewConfigureHAForHostParams mocks base method.
func (m *MockHostServiceIface) NewConfigureHAForHostParams(hostid, provider string) *ConfigureHAForHostParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewConfigureHAForHostParams", hostid, provider)
	ret0, _ := ret[0].(*ConfigureHAForHostParams)
	return ret0
}

// NewConfigureHAForHostParams indicates an expected call of NewConfigureHAForHostParams.
func (mr *MockHostServiceIfaceMockRecorder) NewConfigureHAForHostParams(hostid, provider interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewConfigureHAForHostParams", reflect.TypeOf((*MockHostServiceIface)(nil).NewConfigureHAForHostParams), hostid, provider)
}

// NewDedicateHostParams mocks base method.
func (m *MockHostServiceIface) NewDedicateHostParams(domainid, hostid string) *DedicateHostParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDedicateHostParams", domainid, hostid)
	ret0, _ := ret[0].(*DedicateHostParams)
	return ret0
}

// NewDedicateHostParams indicates an expected call of NewDedicateHostParams.
func (mr *MockHostServiceIfaceMockRecorder) NewDedicateHostParams(domainid, hostid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDedicateHostParams", reflect.TypeOf((*MockHostServiceIface)(nil).NewDedicateHostParams), domainid, hostid)
}

// NewDeleteHostParams mocks base method.
func (m *MockHostServiceIface) NewDeleteHostParams(id string) *DeleteHostParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDeleteHostParams", id)
	ret0, _ := ret[0].(*DeleteHostParams)
	return ret0
}

// NewDeleteHostParams indicates an expected call of NewDeleteHostParams.
func (mr *MockHostServiceIfaceMockRecorder) NewDeleteHostParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDeleteHostParams", reflect.TypeOf((*MockHostServiceIface)(nil).NewDeleteHostParams), id)
}

// NewDisableOutOfBandManagementForHostParams mocks base method.
func (m *MockHostServiceIface) NewDisableOutOfBandManagementForHostParams(hostid string) *DisableOutOfBandManagementForHostParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDisableOutOfBandManagementForHostParams", hostid)
	ret0, _ := ret[0].(*DisableOutOfBandManagementForHostParams)
	return ret0
}

// NewDisableOutOfBandManagementForHostParams indicates an expected call of NewDisableOutOfBandManagementForHostParams.
func (mr *MockHostServiceIfaceMockRecorder) NewDisableOutOfBandManagementForHostParams(hostid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDisableOutOfBandManagementForHostParams", reflect.TypeOf((*MockHostServiceIface)(nil).NewDisableOutOfBandManagementForHostParams), hostid)
}

// NewEnableHAForHostParams mocks base method.
func (m *MockHostServiceIface) NewEnableHAForHostParams(hostid string) *EnableHAForHostParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewEnableHAForHostParams", hostid)
	ret0, _ := ret[0].(*EnableHAForHostParams)
	return ret0
}

// NewEnableHAForHostParams indicates an expected call of NewEnableHAForHostParams.
func (mr *MockHostServiceIfaceMockRecorder) NewEnableHAForHostParams(hostid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewEnableHAForHostParams", reflect.TypeOf((*MockHostServiceIface)(nil).NewEnableHAForHostParams), hostid)
}

// NewEnableOutOfBandManagementForHostParams mocks base method.
func (m *MockHostServiceIface) NewEnableOutOfBandManagementForHostParams(hostid string) *EnableOutOfBandManagementForHostParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewEnableOutOfBandManagementForHostParams", hostid)
	ret0, _ := ret[0].(*EnableOutOfBandManagementForHostParams)
	return ret0
}

// NewEnableOutOfBandManagementForHostParams indicates an expected call of NewEnableOutOfBandManagementForHostParams.
func (mr *MockHostServiceIfaceMockRecorder) NewEnableOutOfBandManagementForHostParams(hostid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewEnableOutOfBandManagementForHostParams", reflect.TypeOf((*MockHostServiceIface)(nil).NewEnableOutOfBandManagementForHostParams), hostid)
}

// NewFindHostsForMigrationParams mocks base method.
func (m *MockHostServiceIface) NewFindHostsForMigrationParams(virtualmachineid string) *FindHostsForMigrationParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewFindHostsForMigrationParams", virtualmachineid)
	ret0, _ := ret[0].(*FindHostsForMigrationParams)
	return ret0
}

// NewFindHostsForMigrationParams indicates an expected call of NewFindHostsForMigrationParams.
func (mr *MockHostServiceIfaceMockRecorder) NewFindHostsForMigrationParams(virtualmachineid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewFindHostsForMigrationParams", reflect.TypeOf((*MockHostServiceIface)(nil).NewFindHostsForMigrationParams), virtualmachineid)
}

// NewListDedicatedHostsParams mocks base method.
func (m *MockHostServiceIface) NewListDedicatedHostsParams() *ListDedicatedHostsParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListDedicatedHostsParams")
	ret0, _ := ret[0].(*ListDedicatedHostsParams)
	return ret0
}

// NewListDedicatedHostsParams indicates an expected call of NewListDedicatedHostsParams.
func (mr *MockHostServiceIfaceMockRecorder) NewListDedicatedHostsParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListDedicatedHostsParams", reflect.TypeOf((*MockHostServiceIface)(nil).NewListDedicatedHostsParams))
}

// NewListHostTagsParams mocks base method.
func (m *MockHostServiceIface) NewListHostTagsParams() *ListHostTagsParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListHostTagsParams")
	ret0, _ := ret[0].(*ListHostTagsParams)
	return ret0
}

// NewListHostTagsParams indicates an expected call of NewListHostTagsParams.
func (mr *MockHostServiceIfaceMockRecorder) NewListHostTagsParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListHostTagsParams", reflect.TypeOf((*MockHostServiceIface)(nil).NewListHostTagsParams))
}

// NewListHostsMetricsParams mocks base method.
func (m *MockHostServiceIface) NewListHostsMetricsParams() *ListHostsMetricsParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListHostsMetricsParams")
	ret0, _ := ret[0].(*ListHostsMetricsParams)
	return ret0
}

// NewListHostsMetricsParams indicates an expected call of NewListHostsMetricsParams.
func (mr *MockHostServiceIfaceMockRecorder) NewListHostsMetricsParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListHostsMetricsParams", reflect.TypeOf((*MockHostServiceIface)(nil).NewListHostsMetricsParams))
}

// NewListHostsParams mocks base method.
func (m *MockHostServiceIface) NewListHostsParams() *ListHostsParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListHostsParams")
	ret0, _ := ret[0].(*ListHostsParams)
	return ret0
}

// NewListHostsParams indicates an expected call of NewListHostsParams.
func (mr *MockHostServiceIfaceMockRecorder) NewListHostsParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListHostsParams", reflect.TypeOf((*MockHostServiceIface)(nil).NewListHostsParams))
}

// NewPrepareHostForMaintenanceParams mocks base method.
func (m *MockHostServiceIface) NewPrepareHostForMaintenanceParams(id string) *PrepareHostForMaintenanceParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewPrepareHostForMaintenanceParams", id)
	ret0, _ := ret[0].(*PrepareHostForMaintenanceParams)
	return ret0
}

// NewPrepareHostForMaintenanceParams indicates an expected call of NewPrepareHostForMaintenanceParams.
func (mr *MockHostServiceIfaceMockRecorder) NewPrepareHostForMaintenanceParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewPrepareHostForMaintenanceParams", reflect.TypeOf((*MockHostServiceIface)(nil).NewPrepareHostForMaintenanceParams), id)
}

// NewReconnectHostParams mocks base method.
func (m *MockHostServiceIface) NewReconnectHostParams(id string) *ReconnectHostParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewReconnectHostParams", id)
	ret0, _ := ret[0].(*ReconnectHostParams)
	return ret0
}

// NewReconnectHostParams indicates an expected call of NewReconnectHostParams.
func (mr *MockHostServiceIfaceMockRecorder) NewReconnectHostParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewReconnectHostParams", reflect.TypeOf((*MockHostServiceIface)(nil).NewReconnectHostParams), id)
}

// NewReleaseDedicatedHostParams mocks base method.
func (m *MockHostServiceIface) NewReleaseDedicatedHostParams(hostid string) *ReleaseDedicatedHostParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewReleaseDedicatedHostParams", hostid)
	ret0, _ := ret[0].(*ReleaseDedicatedHostParams)
	return ret0
}

// NewReleaseDedicatedHostParams indicates an expected call of NewReleaseDedicatedHostParams.
func (mr *MockHostServiceIfaceMockRecorder) NewReleaseDedicatedHostParams(hostid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewReleaseDedicatedHostParams", reflect.TypeOf((*MockHostServiceIface)(nil).NewReleaseDedicatedHostParams), hostid)
}

// NewReleaseHostReservationParams mocks base method.
func (m *MockHostServiceIface) NewReleaseHostReservationParams(id string) *ReleaseHostReservationParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewReleaseHostReservationParams", id)
	ret0, _ := ret[0].(*ReleaseHostReservationParams)
	return ret0
}

// NewReleaseHostReservationParams indicates an expected call of NewReleaseHostReservationParams.
func (mr *MockHostServiceIfaceMockRecorder) NewReleaseHostReservationParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewReleaseHostReservationParams", reflect.TypeOf((*MockHostServiceIface)(nil).NewReleaseHostReservationParams), id)
}

// NewUpdateHostParams mocks base method.
func (m *MockHostServiceIface) NewUpdateHostParams(id string) *UpdateHostParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUpdateHostParams", id)
	ret0, _ := ret[0].(*UpdateHostParams)
	return ret0
}

// NewUpdateHostParams indicates an expected call of NewUpdateHostParams.
func (mr *MockHostServiceIfaceMockRecorder) NewUpdateHostParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUpdateHostParams", reflect.TypeOf((*MockHostServiceIface)(nil).NewUpdateHostParams), id)
}

// NewUpdateHostPasswordParams mocks base method.
func (m *MockHostServiceIface) NewUpdateHostPasswordParams(password, username string) *UpdateHostPasswordParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUpdateHostPasswordParams", password, username)
	ret0, _ := ret[0].(*UpdateHostPasswordParams)
	return ret0
}

// NewUpdateHostPasswordParams indicates an expected call of NewUpdateHostPasswordParams.
func (mr *MockHostServiceIfaceMockRecorder) NewUpdateHostPasswordParams(password, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUpdateHostPasswordParams", reflect.TypeOf((*MockHostServiceIface)(nil).NewUpdateHostPasswordParams), password, username)
}

// PrepareHostForMaintenance mocks base method.
func (m *MockHostServiceIface) PrepareHostForMaintenance(p *PrepareHostForMaintenanceParams) (*PrepareHostForMaintenanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareHostForMaintenance", p)
	ret0, _ := ret[0].(*PrepareHostForMaintenanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareHostForMaintenance indicates an expected call of PrepareHostForMaintenance.
func (mr *MockHostServiceIfaceMockRecorder) PrepareHostForMaintenance(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareHostForMaintenance", reflect.TypeOf((*MockHostServiceIface)(nil).PrepareHostForMaintenance), p)
}

// ReconnectHost mocks base method.
func (m *MockHostServiceIface) ReconnectHost(p *ReconnectHostParams) (*ReconnectHostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconnectHost", p)
	ret0, _ := ret[0].(*ReconnectHostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconnectHost indicates an expected call of ReconnectHost.
func (mr *MockHostServiceIfaceMockRecorder) ReconnectHost(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconnectHost", reflect.TypeOf((*MockHostServiceIface)(nil).ReconnectHost), p)
}

// ReleaseDedicatedHost mocks base method.
func (m *MockHostServiceIface) ReleaseDedicatedHost(p *ReleaseDedicatedHostParams) (*ReleaseDedicatedHostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReleaseDedicatedHost", p)
	ret0, _ := ret[0].(*ReleaseDedicatedHostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReleaseDedicatedHost indicates an expected call of ReleaseDedicatedHost.
func (mr *MockHostServiceIfaceMockRecorder) ReleaseDedicatedHost(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseDedicatedHost", reflect.TypeOf((*MockHostServiceIface)(nil).ReleaseDedicatedHost), p)
}

// ReleaseHostReservation mocks base method.
func (m *MockHostServiceIface) ReleaseHostReservation(p *ReleaseHostReservationParams) (*ReleaseHostReservationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReleaseHostReservation", p)
	ret0, _ := ret[0].(*ReleaseHostReservationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReleaseHostReservation indicates an expected call of ReleaseHostReservation.
func (mr *MockHostServiceIfaceMockRecorder) ReleaseHostReservation(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseHostReservation", reflect.TypeOf((*MockHostServiceIface)(nil).ReleaseHostReservation), p)
}

// UpdateHost mocks base method.
func (m *MockHostServiceIface) UpdateHost(p *UpdateHostParams) (*UpdateHostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHost", p)
	ret0, _ := ret[0].(*UpdateHostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateHost indicates an expected call of UpdateHost.
func (mr *MockHostServiceIfaceMockRecorder) UpdateHost(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHost", reflect.TypeOf((*MockHostServiceIface)(nil).UpdateHost), p)
}

// UpdateHostPassword mocks base method.
func (m *MockHostServiceIface) UpdateHostPassword(p *UpdateHostPasswordParams) (*UpdateHostPasswordResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHostPassword", p)
	ret0, _ := ret[0].(*UpdateHostPasswordResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateHostPassword indicates an expected call of UpdateHostPassword.
func (mr *MockHostServiceIfaceMockRecorder) UpdateHostPassword(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHostPassword", reflect.TypeOf((*MockHostServiceIface)(nil).UpdateHostPassword), p)
}
