// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/baetyl/baetyl-cloud/v2/service (interfaces: ObjectService)

// Package service is a generated GoMock package.
package service

import (
	models "github.com/baetyl/baetyl-cloud/v2/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockObjectService is a mock of ObjectService interface
type MockObjectService struct {
	ctrl     *gomock.Controller
	recorder *MockObjectServiceMockRecorder
}

// MockObjectServiceMockRecorder is the mock recorder for MockObjectService
type MockObjectServiceMockRecorder struct {
	mock *MockObjectService
}

// NewMockObjectService creates a new mock instance
func NewMockObjectService(ctrl *gomock.Controller) *MockObjectService {
	mock := &MockObjectService{ctrl: ctrl}
	mock.recorder = &MockObjectServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockObjectService) EXPECT() *MockObjectServiceMockRecorder {
	return m.recorder
}

// CreateBucketIfNotExist mocks base method
func (m *MockObjectService) CreateBucketIfNotExist(arg0, arg1, arg2, arg3 string) (*models.Bucket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBucketIfNotExist", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*models.Bucket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBucketIfNotExist indicates an expected call of CreateBucketIfNotExist
func (mr *MockObjectServiceMockRecorder) CreateBucketIfNotExist(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBucketIfNotExist", reflect.TypeOf((*MockObjectService)(nil).CreateBucketIfNotExist), arg0, arg1, arg2, arg3)
}

// GenObjectURL mocks base method
func (m *MockObjectService) GenObjectURL(arg0 string, arg1 models.ConfigObjectItem) (*models.ObjectURL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenObjectURL", arg0, arg1)
	ret0, _ := ret[0].(*models.ObjectURL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenObjectURL indicates an expected call of GenObjectURL
func (mr *MockObjectServiceMockRecorder) GenObjectURL(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenObjectURL", reflect.TypeOf((*MockObjectService)(nil).GenObjectURL), arg0, arg1)
}

// ListBucketObjects mocks base method
func (m *MockObjectService) ListBucketObjects(arg0, arg1, arg2 string) (*models.ListObjectsResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBucketObjects", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.ListObjectsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBucketObjects indicates an expected call of ListBucketObjects
func (mr *MockObjectServiceMockRecorder) ListBucketObjects(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBucketObjects", reflect.TypeOf((*MockObjectService)(nil).ListBucketObjects), arg0, arg1, arg2)
}

// ListBuckets mocks base method
func (m *MockObjectService) ListBuckets(arg0, arg1 string) ([]models.Bucket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBuckets", arg0, arg1)
	ret0, _ := ret[0].([]models.Bucket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBuckets indicates an expected call of ListBuckets
func (mr *MockObjectServiceMockRecorder) ListBuckets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBuckets", reflect.TypeOf((*MockObjectService)(nil).ListBuckets), arg0, arg1)
}

// ListSources mocks base method
func (m *MockObjectService) ListSources() []models.ObjectStorageSource {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSources")
	ret0, _ := ret[0].([]models.ObjectStorageSource)
	return ret0
}

// ListSources indicates an expected call of ListSources
func (mr *MockObjectServiceMockRecorder) ListSources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSources", reflect.TypeOf((*MockObjectService)(nil).ListSources))
}

// PutObjectFromURLIfNotExist mocks base method
func (m *MockObjectService) PutObjectFromURLIfNotExist(arg0, arg1, arg2, arg3, arg4 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutObjectFromURLIfNotExist", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutObjectFromURLIfNotExist indicates an expected call of PutObjectFromURLIfNotExist
func (mr *MockObjectServiceMockRecorder) PutObjectFromURLIfNotExist(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObjectFromURLIfNotExist", reflect.TypeOf((*MockObjectService)(nil).PutObjectFromURLIfNotExist), arg0, arg1, arg2, arg3, arg4)
}
