// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gopowerstore "github.com/dell/gopowerstore"
	api "github.com/dell/gopowerstore/api"
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// APIClient mocks base method
func (m *MockClient) APIClient() api.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIClient")
	ret0, _ := ret[0].(api.Client)
	return ret0
}

// APIClient indicates an expected call of APIClient
func (mr *MockClientMockRecorder) APIClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIClient", reflect.TypeOf((*MockClient)(nil).APIClient))
}

// SetTraceID mocks base method
func (m *MockClient) SetTraceID(ctx context.Context, value string) context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTraceID", ctx, value)
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// SetTraceID indicates an expected call of SetTraceID
func (mr *MockClientMockRecorder) SetTraceID(ctx, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTraceID", reflect.TypeOf((*MockClient)(nil).SetTraceID), ctx, value)
}

// SetCustomHTTPHeaders mocks base method
func (m *MockClient) SetCustomHTTPHeaders(headers http.Header) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCustomHTTPHeaders", headers)
}

// SetCustomHTTPHeaders indicates an expected call of SetCustomHTTPHeaders
func (mr *MockClientMockRecorder) SetCustomHTTPHeaders(headers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCustomHTTPHeaders", reflect.TypeOf((*MockClient)(nil).SetCustomHTTPHeaders), headers)
}

// GetVolume mocks base method
func (m *MockClient) GetVolume(ctx context.Context, id string) (gopowerstore.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolume", ctx, id)
	ret0, _ := ret[0].(gopowerstore.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVolume indicates an expected call of GetVolume
func (mr *MockClientMockRecorder) GetVolume(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolume", reflect.TypeOf((*MockClient)(nil).GetVolume), ctx, id)
}

// GetVolumeByName mocks base method
func (m *MockClient) GetVolumeByName(ctx context.Context, name string) (gopowerstore.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolumeByName", ctx, name)
	ret0, _ := ret[0].(gopowerstore.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVolumeByName indicates an expected call of GetVolumeByName
func (mr *MockClientMockRecorder) GetVolumeByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolumeByName", reflect.TypeOf((*MockClient)(nil).GetVolumeByName), ctx, name)
}

// GetVolumes mocks base method
func (m *MockClient) GetVolumes(ctx context.Context) ([]gopowerstore.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolumes", ctx)
	ret0, _ := ret[0].([]gopowerstore.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVolumes indicates an expected call of GetVolumes
func (mr *MockClientMockRecorder) GetVolumes(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolumes", reflect.TypeOf((*MockClient)(nil).GetVolumes), ctx)
}

// CreateVolume mocks base method
func (m *MockClient) CreateVolume(ctx context.Context, createParams *gopowerstore.VolumeCreate) (gopowerstore.CreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVolume", ctx, createParams)
	ret0, _ := ret[0].(gopowerstore.CreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVolume indicates an expected call of CreateVolume
func (mr *MockClientMockRecorder) CreateVolume(ctx, createParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVolume", reflect.TypeOf((*MockClient)(nil).CreateVolume), ctx, createParams)
}

// DeleteVolume mocks base method
func (m *MockClient) DeleteVolume(ctx context.Context, deleteParams *gopowerstore.VolumeDelete, id string) (gopowerstore.EmptyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVolume", ctx, deleteParams, id)
	ret0, _ := ret[0].(gopowerstore.EmptyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteVolume indicates an expected call of DeleteVolume
func (mr *MockClientMockRecorder) DeleteVolume(ctx, deleteParams, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVolume", reflect.TypeOf((*MockClient)(nil).DeleteVolume), ctx, deleteParams, id)
}

// GetHost mocks base method
func (m *MockClient) GetHost(ctx context.Context, id string) (gopowerstore.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHost", ctx, id)
	ret0, _ := ret[0].(gopowerstore.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHost indicates an expected call of GetHost
func (mr *MockClientMockRecorder) GetHost(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHost", reflect.TypeOf((*MockClient)(nil).GetHost), ctx, id)
}

// GetHostByName mocks base method
func (m *MockClient) GetHostByName(ctx context.Context, name string) (gopowerstore.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostByName", ctx, name)
	ret0, _ := ret[0].(gopowerstore.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHostByName indicates an expected call of GetHostByName
func (mr *MockClientMockRecorder) GetHostByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostByName", reflect.TypeOf((*MockClient)(nil).GetHostByName), ctx, name)
}

// GetHosts mocks base method
func (m *MockClient) GetHosts(ctx context.Context) ([]gopowerstore.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHosts", ctx)
	ret0, _ := ret[0].([]gopowerstore.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHosts indicates an expected call of GetHosts
func (mr *MockClientMockRecorder) GetHosts(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHosts", reflect.TypeOf((*MockClient)(nil).GetHosts), ctx)
}

// CreateHost mocks base method
func (m *MockClient) CreateHost(ctx context.Context, createParams *gopowerstore.HostCreate) (gopowerstore.CreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHost", ctx, createParams)
	ret0, _ := ret[0].(gopowerstore.CreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateHost indicates an expected call of CreateHost
func (mr *MockClientMockRecorder) CreateHost(ctx, createParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHost", reflect.TypeOf((*MockClient)(nil).CreateHost), ctx, createParams)
}

// DeleteHost mocks base method
func (m *MockClient) DeleteHost(ctx context.Context, deleteParams *gopowerstore.HostDelete, id string) (gopowerstore.EmptyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteHost", ctx, deleteParams, id)
	ret0, _ := ret[0].(gopowerstore.EmptyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteHost indicates an expected call of DeleteHost
func (mr *MockClientMockRecorder) DeleteHost(ctx, deleteParams, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHost", reflect.TypeOf((*MockClient)(nil).DeleteHost), ctx, deleteParams, id)
}

// ModifyHost mocks base method
func (m *MockClient) ModifyHost(ctx context.Context, modifyParams *gopowerstore.HostModify, id string) (gopowerstore.CreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyHost", ctx, modifyParams, id)
	ret0, _ := ret[0].(gopowerstore.CreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyHost indicates an expected call of ModifyHost
func (mr *MockClientMockRecorder) ModifyHost(ctx, modifyParams, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyHost", reflect.TypeOf((*MockClient)(nil).ModifyHost), ctx, modifyParams, id)
}

// GetHostVolumeMappings mocks base method
func (m *MockClient) GetHostVolumeMappings(ctx context.Context) ([]gopowerstore.HostVolumeMapping, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostVolumeMappings", ctx)
	ret0, _ := ret[0].([]gopowerstore.HostVolumeMapping)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHostVolumeMappings indicates an expected call of GetHostVolumeMappings
func (mr *MockClientMockRecorder) GetHostVolumeMappings(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostVolumeMappings", reflect.TypeOf((*MockClient)(nil).GetHostVolumeMappings), ctx)
}

// GetHostVolumeMapping mocks base method
func (m *MockClient) GetHostVolumeMapping(ctx context.Context, id string) (gopowerstore.HostVolumeMapping, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostVolumeMapping", ctx, id)
	ret0, _ := ret[0].(gopowerstore.HostVolumeMapping)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHostVolumeMapping indicates an expected call of GetHostVolumeMapping
func (mr *MockClientMockRecorder) GetHostVolumeMapping(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostVolumeMapping", reflect.TypeOf((*MockClient)(nil).GetHostVolumeMapping), ctx, id)
}

// GetHostVolumeMappingByVolumeID mocks base method
func (m *MockClient) GetHostVolumeMappingByVolumeID(ctx context.Context, volumeID string) ([]gopowerstore.HostVolumeMapping, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostVolumeMappingByVolumeID", ctx, volumeID)
	ret0, _ := ret[0].([]gopowerstore.HostVolumeMapping)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHostVolumeMappingByVolumeID indicates an expected call of GetHostVolumeMappingByVolumeID
func (mr *MockClientMockRecorder) GetHostVolumeMappingByVolumeID(ctx, volumeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostVolumeMappingByVolumeID", reflect.TypeOf((*MockClient)(nil).GetHostVolumeMappingByVolumeID), ctx, volumeID)
}

// AttachVolumeToHost mocks base method
func (m *MockClient) AttachVolumeToHost(ctx context.Context, hostID string, attachParams *gopowerstore.HostVolumeAttach) (gopowerstore.EmptyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachVolumeToHost", ctx, hostID, attachParams)
	ret0, _ := ret[0].(gopowerstore.EmptyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AttachVolumeToHost indicates an expected call of AttachVolumeToHost
func (mr *MockClientMockRecorder) AttachVolumeToHost(ctx, hostID, attachParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachVolumeToHost", reflect.TypeOf((*MockClient)(nil).AttachVolumeToHost), ctx, hostID, attachParams)
}

// DetachVolumeFromHost mocks base method
func (m *MockClient) DetachVolumeFromHost(ctx context.Context, hostID string, detachParams *gopowerstore.HostVolumeDetach) (gopowerstore.EmptyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetachVolumeFromHost", ctx, hostID, detachParams)
	ret0, _ := ret[0].(gopowerstore.EmptyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetachVolumeFromHost indicates an expected call of DetachVolumeFromHost
func (mr *MockClientMockRecorder) DetachVolumeFromHost(ctx, hostID, detachParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachVolumeFromHost", reflect.TypeOf((*MockClient)(nil).DetachVolumeFromHost), ctx, hostID, detachParams)
}

// GetStorageISCSITargetAddresses mocks base method
func (m *MockClient) GetStorageISCSITargetAddresses(ctx context.Context) ([]gopowerstore.IPPoolAddress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageISCSITargetAddresses", ctx)
	ret0, _ := ret[0].([]gopowerstore.IPPoolAddress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStorageISCSITargetAddresses indicates an expected call of GetStorageISCSITargetAddresses
func (mr *MockClientMockRecorder) GetStorageISCSITargetAddresses(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageISCSITargetAddresses", reflect.TypeOf((*MockClient)(nil).GetStorageISCSITargetAddresses), ctx)
}

// GetCapacity mocks base method
func (m *MockClient) GetCapacity(ctx context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCapacity", ctx)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCapacity indicates an expected call of GetCapacity
func (mr *MockClientMockRecorder) GetCapacity(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCapacity", reflect.TypeOf((*MockClient)(nil).GetCapacity), ctx)
}

// GetFCPorts mocks base method
func (m *MockClient) GetFCPorts(ctx context.Context) ([]gopowerstore.FcPort, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFCPorts", ctx)
	ret0, _ := ret[0].([]gopowerstore.FcPort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFCPorts indicates an expected call of GetFCPorts
func (mr *MockClientMockRecorder) GetFCPorts(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFCPorts", reflect.TypeOf((*MockClient)(nil).GetFCPorts), ctx)
}

// GetFCPort mocks base method
func (m *MockClient) GetFCPort(ctx context.Context, id string) (gopowerstore.FcPort, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFCPort", ctx, id)
	ret0, _ := ret[0].(gopowerstore.FcPort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFCPort indicates an expected call of GetFCPort
func (mr *MockClientMockRecorder) GetFCPort(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFCPort", reflect.TypeOf((*MockClient)(nil).GetFCPort), ctx, id)
}

// SetLogger mocks base method
func (m *MockClient) SetLogger(logger gopowerstore.Logger) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLogger", logger)
}

// SetLogger indicates an expected call of SetLogger
func (mr *MockClientMockRecorder) SetLogger(logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLogger", reflect.TypeOf((*MockClient)(nil).SetLogger), logger)
}

// CreateSnapshot mocks base method
func (m *MockClient) CreateSnapshot(ctx context.Context, createSnapParams *gopowerstore.SnapshotCreate, id string) (gopowerstore.CreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSnapshot", ctx, createSnapParams, id)
	ret0, _ := ret[0].(gopowerstore.CreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSnapshot indicates an expected call of CreateSnapshot
func (mr *MockClientMockRecorder) CreateSnapshot(ctx, createSnapParams, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSnapshot", reflect.TypeOf((*MockClient)(nil).CreateSnapshot), ctx, createSnapParams, id)
}

// DeleteSnapshot mocks base method
func (m *MockClient) DeleteSnapshot(ctx context.Context, deleteParams *gopowerstore.VolumeDelete, id string) (gopowerstore.EmptyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSnapshot", ctx, deleteParams, id)
	ret0, _ := ret[0].(gopowerstore.EmptyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSnapshot indicates an expected call of DeleteSnapshot
func (mr *MockClientMockRecorder) DeleteSnapshot(ctx, deleteParams, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSnapshot", reflect.TypeOf((*MockClient)(nil).DeleteSnapshot), ctx, deleteParams, id)
}

// GetSnapshotsByVolumeID mocks base method
func (m *MockClient) GetSnapshotsByVolumeID(ctx context.Context, volID string) ([]gopowerstore.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSnapshotsByVolumeID", ctx, volID)
	ret0, _ := ret[0].([]gopowerstore.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSnapshotsByVolumeID indicates an expected call of GetSnapshotsByVolumeID
func (mr *MockClientMockRecorder) GetSnapshotsByVolumeID(ctx, volID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSnapshotsByVolumeID", reflect.TypeOf((*MockClient)(nil).GetSnapshotsByVolumeID), ctx, volID)
}

// GetSnapshots mocks base method
func (m *MockClient) GetSnapshots(ctx context.Context) ([]gopowerstore.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSnapshots", ctx)
	ret0, _ := ret[0].([]gopowerstore.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSnapshots indicates an expected call of GetSnapshots
func (mr *MockClientMockRecorder) GetSnapshots(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSnapshots", reflect.TypeOf((*MockClient)(nil).GetSnapshots), ctx)
}

// GetSnapshot mocks base method
func (m *MockClient) GetSnapshot(ctx context.Context, snapID string) (gopowerstore.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSnapshot", ctx, snapID)
	ret0, _ := ret[0].(gopowerstore.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSnapshot indicates an expected call of GetSnapshot
func (mr *MockClientMockRecorder) GetSnapshot(ctx, snapID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSnapshot", reflect.TypeOf((*MockClient)(nil).GetSnapshot), ctx, snapID)
}

// CreateVolumeFromSnapshot mocks base method
func (m *MockClient) CreateVolumeFromSnapshot(ctx context.Context, createParams *gopowerstore.VolumeClone, snapID string) (gopowerstore.CreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVolumeFromSnapshot", ctx, createParams, snapID)
	ret0, _ := ret[0].(gopowerstore.CreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVolumeFromSnapshot indicates an expected call of CreateVolumeFromSnapshot
func (mr *MockClientMockRecorder) CreateVolumeFromSnapshot(ctx, createParams, snapID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVolumeFromSnapshot", reflect.TypeOf((*MockClient)(nil).CreateVolumeFromSnapshot), ctx, createParams, snapID)
}

// GetNAS mocks base method
func (m *MockClient) GetNAS(ctx context.Context, id string) (gopowerstore.NAS, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNAS", ctx, id)
	ret0, _ := ret[0].(gopowerstore.NAS)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNAS indicates an expected call of GetNAS
func (mr *MockClientMockRecorder) GetNAS(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNAS", reflect.TypeOf((*MockClient)(nil).GetNAS), ctx, id)
}

// GetNASByName mocks base method
func (m *MockClient) GetNASByName(ctx context.Context, name string) (gopowerstore.NAS, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNASByName", ctx, name)
	ret0, _ := ret[0].(gopowerstore.NAS)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNASByName indicates an expected call of GetNASByName
func (mr *MockClientMockRecorder) GetNASByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNASByName", reflect.TypeOf((*MockClient)(nil).GetNASByName), ctx, name)
}

// GetFSByName mocks base method
func (m *MockClient) GetFSByName(ctx context.Context, name string) (gopowerstore.FileSystem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFSByName", ctx, name)
	ret0, _ := ret[0].(gopowerstore.FileSystem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFSByName indicates an expected call of GetFSByName
func (mr *MockClientMockRecorder) GetFSByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFSByName", reflect.TypeOf((*MockClient)(nil).GetFSByName), ctx, name)
}

// GetFS mocks base method
func (m *MockClient) GetFS(ctx context.Context, id string) (gopowerstore.FileSystem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFS", ctx, id)
	ret0, _ := ret[0].(gopowerstore.FileSystem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFS indicates an expected call of GetFS
func (mr *MockClientMockRecorder) GetFS(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFS", reflect.TypeOf((*MockClient)(nil).GetFS), ctx, id)
}

// GetFileInterface mocks base method
func (m *MockClient) GetFileInterface(ctx context.Context, id string) (gopowerstore.FileInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileInterface", ctx, id)
	ret0, _ := ret[0].(gopowerstore.FileInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileInterface indicates an expected call of GetFileInterface
func (mr *MockClientMockRecorder) GetFileInterface(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileInterface", reflect.TypeOf((*MockClient)(nil).GetFileInterface), ctx, id)
}

// GetNFSExportByName mocks base method
func (m *MockClient) GetNFSExportByName(ctx context.Context, name string) (gopowerstore.NFSExport, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNFSExportByName", ctx, name)
	ret0, _ := ret[0].(gopowerstore.NFSExport)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNFSExportByName indicates an expected call of GetNFSExportByName
func (mr *MockClientMockRecorder) GetNFSExportByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNFSExportByName", reflect.TypeOf((*MockClient)(nil).GetNFSExportByName), ctx, name)
}

// GetNFSExportByFileSystemID mocks base method
func (m *MockClient) GetNFSExportByFileSystemID(ctx context.Context, fsID string) (gopowerstore.NFSExport, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNFSExportByFileSystemID", ctx, fsID)
	ret0, _ := ret[0].(gopowerstore.NFSExport)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNFSExportByFileSystemID indicates an expected call of GetNFSExportByFileSystemID
func (mr *MockClientMockRecorder) GetNFSExportByFileSystemID(ctx, fsID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNFSExportByFileSystemID", reflect.TypeOf((*MockClient)(nil).GetNFSExportByFileSystemID), ctx, fsID)
}

// CreateNAS mocks base method
func (m *MockClient) CreateNAS(ctx context.Context, createParams *gopowerstore.NASCreate) (gopowerstore.CreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNAS", ctx, createParams)
	ret0, _ := ret[0].(gopowerstore.CreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNAS indicates an expected call of CreateNAS
func (mr *MockClientMockRecorder) CreateNAS(ctx, createParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNAS", reflect.TypeOf((*MockClient)(nil).CreateNAS), ctx, createParams)
}

// DeleteNAS mocks base method
func (m *MockClient) DeleteNAS(ctx context.Context, id string) (gopowerstore.EmptyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNAS", ctx, id)
	ret0, _ := ret[0].(gopowerstore.EmptyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteNAS indicates an expected call of DeleteNAS
func (mr *MockClientMockRecorder) DeleteNAS(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNAS", reflect.TypeOf((*MockClient)(nil).DeleteNAS), ctx, id)
}

// CreateFS mocks base method
func (m *MockClient) CreateFS(ctx context.Context, createParams *gopowerstore.FsCreate) (gopowerstore.CreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFS", ctx, createParams)
	ret0, _ := ret[0].(gopowerstore.CreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFS indicates an expected call of CreateFS
func (mr *MockClientMockRecorder) CreateFS(ctx, createParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFS", reflect.TypeOf((*MockClient)(nil).CreateFS), ctx, createParams)
}

// DeleteFS mocks base method
func (m *MockClient) DeleteFS(ctx context.Context, id string) (gopowerstore.EmptyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFS", ctx, id)
	ret0, _ := ret[0].(gopowerstore.EmptyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFS indicates an expected call of DeleteFS
func (mr *MockClientMockRecorder) DeleteFS(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFS", reflect.TypeOf((*MockClient)(nil).DeleteFS), ctx, id)
}

// CreateNFSExport mocks base method
func (m *MockClient) CreateNFSExport(ctx context.Context, createParams *gopowerstore.NFSExportCreate) (gopowerstore.CreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNFSExport", ctx, createParams)
	ret0, _ := ret[0].(gopowerstore.CreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNFSExport indicates an expected call of CreateNFSExport
func (mr *MockClientMockRecorder) CreateNFSExport(ctx, createParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNFSExport", reflect.TypeOf((*MockClient)(nil).CreateNFSExport), ctx, createParams)
}

// DeleteNFSExport mocks base method
func (m *MockClient) DeleteNFSExport(ctx context.Context, id string) (gopowerstore.EmptyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNFSExport", ctx, id)
	ret0, _ := ret[0].(gopowerstore.EmptyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteNFSExport indicates an expected call of DeleteNFSExport
func (mr *MockClientMockRecorder) DeleteNFSExport(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNFSExport", reflect.TypeOf((*MockClient)(nil).DeleteNFSExport), ctx, id)
}

// ModifyNFSExport mocks base method
func (m *MockClient) ModifyNFSExport(ctx context.Context, modifyParams *gopowerstore.NFSExportModify, id string) (gopowerstore.CreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyNFSExport", ctx, modifyParams, id)
	ret0, _ := ret[0].(gopowerstore.CreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyNFSExport indicates an expected call of ModifyNFSExport
func (mr *MockClientMockRecorder) ModifyNFSExport(ctx, modifyParams, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyNFSExport", reflect.TypeOf((*MockClient)(nil).ModifyNFSExport), ctx, modifyParams, id)
}

// CreateNFSServer mocks base method
func (m *MockClient) CreateNFSServer(ctx context.Context, createParams *gopowerstore.NFSServerCreate) (gopowerstore.CreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNFSServer", ctx, createParams)
	ret0, _ := ret[0].(gopowerstore.CreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNFSServer indicates an expected call of CreateNFSServer
func (mr *MockClientMockRecorder) CreateNFSServer(ctx, createParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNFSServer", reflect.TypeOf((*MockClient)(nil).CreateNFSServer), ctx, createParams)
}

// CreateFsSnapshot mocks base method
func (m *MockClient) CreateFsSnapshot(ctx context.Context, createSnapParams *gopowerstore.SnapshotFSCreate, id string) (gopowerstore.CreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFsSnapshot", ctx, createSnapParams, id)
	ret0, _ := ret[0].(gopowerstore.CreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFsSnapshot indicates an expected call of CreateFsSnapshot
func (mr *MockClientMockRecorder) CreateFsSnapshot(ctx, createSnapParams, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFsSnapshot", reflect.TypeOf((*MockClient)(nil).CreateFsSnapshot), ctx, createSnapParams, id)
}

// DeleteFsSnapshot mocks base method
func (m *MockClient) DeleteFsSnapshot(ctx context.Context, id string) (gopowerstore.EmptyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFsSnapshot", ctx, id)
	ret0, _ := ret[0].(gopowerstore.EmptyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFsSnapshot indicates an expected call of DeleteFsSnapshot
func (mr *MockClientMockRecorder) DeleteFsSnapshot(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFsSnapshot", reflect.TypeOf((*MockClient)(nil).DeleteFsSnapshot), ctx, id)
}

// GetFsSnapshotsByVolumeID mocks base method
func (m *MockClient) GetFsSnapshotsByVolumeID(ctx context.Context, volID string) ([]gopowerstore.FileSystem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFsSnapshotsByVolumeID", ctx, volID)
	ret0, _ := ret[0].([]gopowerstore.FileSystem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFsSnapshotsByVolumeID indicates an expected call of GetFsSnapshotsByVolumeID
func (mr *MockClientMockRecorder) GetFsSnapshotsByVolumeID(ctx, volID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFsSnapshotsByVolumeID", reflect.TypeOf((*MockClient)(nil).GetFsSnapshotsByVolumeID), ctx, volID)
}

// GetFsSnapshots mocks base method
func (m *MockClient) GetFsSnapshots(ctx context.Context) ([]gopowerstore.FileSystem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFsSnapshots", ctx)
	ret0, _ := ret[0].([]gopowerstore.FileSystem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFsSnapshots indicates an expected call of GetFsSnapshots
func (mr *MockClientMockRecorder) GetFsSnapshots(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFsSnapshots", reflect.TypeOf((*MockClient)(nil).GetFsSnapshots), ctx)
}

// GetFsSnapshot mocks base method
func (m *MockClient) GetFsSnapshot(ctx context.Context, snapID string) (gopowerstore.FileSystem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFsSnapshot", ctx, snapID)
	ret0, _ := ret[0].(gopowerstore.FileSystem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFsSnapshot indicates an expected call of GetFsSnapshot
func (mr *MockClientMockRecorder) GetFsSnapshot(ctx, snapID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFsSnapshot", reflect.TypeOf((*MockClient)(nil).GetFsSnapshot), ctx, snapID)
}

// CreateFsFromSnapshot mocks base method
func (m *MockClient) CreateFsFromSnapshot(ctx context.Context, createParams *gopowerstore.FsClone, snapID string) (gopowerstore.CreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFsFromSnapshot", ctx, createParams, snapID)
	ret0, _ := ret[0].(gopowerstore.CreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFsFromSnapshot indicates an expected call of CreateFsFromSnapshot
func (mr *MockClientMockRecorder) CreateFsFromSnapshot(ctx, createParams, snapID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFsFromSnapshot", reflect.TypeOf((*MockClient)(nil).CreateFsFromSnapshot), ctx, createParams, snapID)
}

// CloneVolume mocks base method
func (m *MockClient) CloneVolume(ctx context.Context, createParams *gopowerstore.VolumeClone, volID string) (gopowerstore.CreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloneVolume", ctx, createParams, volID)
	ret0, _ := ret[0].(gopowerstore.CreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloneVolume indicates an expected call of CloneVolume
func (mr *MockClientMockRecorder) CloneVolume(ctx, createParams, volID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloneVolume", reflect.TypeOf((*MockClient)(nil).CloneVolume), ctx, createParams, volID)
}

// ModifyVolume mocks base method
func (m *MockClient) ModifyVolume(ctx context.Context, modifyParams *gopowerstore.VolumeModify, volID string) (gopowerstore.EmptyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyVolume", ctx, modifyParams, volID)
	ret0, _ := ret[0].(gopowerstore.EmptyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyVolume indicates an expected call of ModifyVolume
func (mr *MockClientMockRecorder) ModifyVolume(ctx, modifyParams, volID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyVolume", reflect.TypeOf((*MockClient)(nil).ModifyVolume), ctx, modifyParams, volID)
}

// ModifyFS mocks base method
func (m *MockClient) ModifyFS(ctx context.Context, modifyParams *gopowerstore.FSModify, volID string) (gopowerstore.EmptyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyFS", ctx, modifyParams, volID)
	ret0, _ := ret[0].(gopowerstore.EmptyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyFS indicates an expected call of ModifyFS
func (mr *MockClientMockRecorder) ModifyFS(ctx, modifyParams, volID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyFS", reflect.TypeOf((*MockClient)(nil).ModifyFS), ctx, modifyParams, volID)
}

// CloneFS mocks base method
func (m *MockClient) CloneFS(ctx context.Context, createParams *gopowerstore.FsClone, fsID string) (gopowerstore.CreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloneFS", ctx, createParams, fsID)
	ret0, _ := ret[0].(gopowerstore.CreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloneFS indicates an expected call of CloneFS
func (mr *MockClientMockRecorder) CloneFS(ctx, createParams, fsID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloneFS", reflect.TypeOf((*MockClient)(nil).CloneFS), ctx, createParams, fsID)
}
