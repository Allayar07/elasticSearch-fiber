// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	models "elasticSearch/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBooks is a mock of Books interface.
type MockBooks struct {
	ctrl     *gomock.Controller
	recorder *MockBooksMockRecorder
}

// MockBooksMockRecorder is the mock recorder for MockBooks.
type MockBooksMockRecorder struct {
	mock *MockBooks
}

// NewMockBooks creates a new mock instance.
func NewMockBooks(ctrl *gomock.Controller) *MockBooks {
	mock := &MockBooks{ctrl: ctrl}
	mock.recorder = &MockBooksMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBooks) EXPECT() *MockBooksMockRecorder {
	return m.recorder
}

// CreateBook mocks base method.
func (m *MockBooks) CreateBook(book models.Book) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBook", book)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBook indicates an expected call of CreateBook.
func (mr *MockBooksMockRecorder) CreateBook(book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockBooks)(nil).CreateBook), book)
}

// Delete mocks base method.
func (m *MockBooks) Delete(ids models.DeleteIds) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ids)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockBooksMockRecorder) Delete(ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBooks)(nil).Delete), ids)
}

// GetFormCache mocks base method.
func (m *MockBooks) GetFormCache(search interface{}) (models.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFormCache", search)
	ret0, _ := ret[0].(models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFormCache indicates an expected call of GetFormCache.
func (mr *MockBooksMockRecorder) GetFormCache(search interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFormCache", reflect.TypeOf((*MockBooks)(nil).GetFormCache), search)
}

// Search mocks base method.
func (m *MockBooks) Search(searchInput string) ([]models.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", searchInput)
	ret0, _ := ret[0].([]models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockBooksMockRecorder) Search(searchInput interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockBooks)(nil).Search), searchInput)
}

// Sync mocks base method.
func (m *MockBooks) Sync() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sync")
	ret0, _ := ret[0].(error)
	return ret0
}

// Sync indicates an expected call of Sync.
func (mr *MockBooksMockRecorder) Sync() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockBooks)(nil).Sync))
}

// Update mocks base method.
func (m *MockBooks) Update(book models.Book) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", book)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockBooksMockRecorder) Update(book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBooks)(nil).Update), book)
}
