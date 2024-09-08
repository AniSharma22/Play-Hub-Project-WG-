// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\Dell\projects2\watchguard-proj\Project-WG\play-hub\internal\domain\interfaces\repository\notification_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	entities "project2/internal/domain/entities"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockNotificationRepository is a mock of NotificationRepository interface.
type MockNotificationRepository struct {
	ctrl     *gomock.Controller
	recorder *MockNotificationRepositoryMockRecorder
}

// MockNotificationRepositoryMockRecorder is the mock recorder for MockNotificationRepository.
type MockNotificationRepositoryMockRecorder struct {
	mock *MockNotificationRepository
}

// NewMockNotificationRepository creates a new mock instance.
func NewMockNotificationRepository(ctrl *gomock.Controller) *MockNotificationRepository {
	mock := &MockNotificationRepository{ctrl: ctrl}
	mock.recorder = &MockNotificationRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotificationRepository) EXPECT() *MockNotificationRepositoryMockRecorder {
	return m.recorder
}

// CreateNotification mocks base method.
func (m *MockNotificationRepository) CreateNotification(ctx context.Context, notification *entities.Notification) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNotification", ctx, notification)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNotification indicates an expected call of CreateNotification.
func (mr *MockNotificationRepositoryMockRecorder) CreateNotification(ctx, notification interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNotification", reflect.TypeOf((*MockNotificationRepository)(nil).CreateNotification), ctx, notification)
}

// DeleteNotificationByID mocks base method.
func (m *MockNotificationRepository) DeleteNotificationByID(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNotificationByID", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNotificationByID indicates an expected call of DeleteNotificationByID.
func (mr *MockNotificationRepositoryMockRecorder) DeleteNotificationByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNotificationByID", reflect.TypeOf((*MockNotificationRepository)(nil).DeleteNotificationByID), ctx, id)
}

// FetchNotificationByID mocks base method.
func (m *MockNotificationRepository) FetchNotificationByID(ctx context.Context, id uuid.UUID) (*entities.Notification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchNotificationByID", ctx, id)
	ret0, _ := ret[0].(*entities.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchNotificationByID indicates an expected call of FetchNotificationByID.
func (mr *MockNotificationRepositoryMockRecorder) FetchNotificationByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchNotificationByID", reflect.TypeOf((*MockNotificationRepository)(nil).FetchNotificationByID), ctx, id)
}

// FetchUserNotifications mocks base method.
func (m *MockNotificationRepository) FetchUserNotifications(ctx context.Context, userID uuid.UUID) ([]entities.Notification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchUserNotifications", ctx, userID)
	ret0, _ := ret[0].([]entities.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchUserNotifications indicates an expected call of FetchUserNotifications.
func (mr *MockNotificationRepositoryMockRecorder) FetchUserNotifications(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchUserNotifications", reflect.TypeOf((*MockNotificationRepository)(nil).FetchUserNotifications), ctx, userID)
}

// MarkNotificationAsRead mocks base method.
func (m *MockNotificationRepository) MarkNotificationAsRead(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkNotificationAsRead", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkNotificationAsRead indicates an expected call of MarkNotificationAsRead.
func (mr *MockNotificationRepositoryMockRecorder) MarkNotificationAsRead(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkNotificationAsRead", reflect.TypeOf((*MockNotificationRepository)(nil).MarkNotificationAsRead), ctx, id)
}