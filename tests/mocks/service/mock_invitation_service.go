// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\Dell\projects2\watchguard-proj\Project-WG\play-hub\internal\domain\interfaces\service\invitation_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	models "project2/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockInvitationService is a mock of InvitationService interface.
type MockInvitationService struct {
	ctrl     *gomock.Controller
	recorder *MockInvitationServiceMockRecorder
}

// MockInvitationServiceMockRecorder is the mock recorder for MockInvitationService.
type MockInvitationServiceMockRecorder struct {
	mock *MockInvitationService
}

// NewMockInvitationService creates a new mock instance.
func NewMockInvitationService(ctrl *gomock.Controller) *MockInvitationService {
	mock := &MockInvitationService{ctrl: ctrl}
	mock.recorder = &MockInvitationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInvitationService) EXPECT() *MockInvitationServiceMockRecorder {
	return m.recorder
}

// AcceptInvitation mocks base method.
func (m *MockInvitationService) AcceptInvitation(ctx context.Context, invitationID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptInvitation", ctx, invitationID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AcceptInvitation indicates an expected call of AcceptInvitation.
func (mr *MockInvitationServiceMockRecorder) AcceptInvitation(ctx, invitationID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptInvitation", reflect.TypeOf((*MockInvitationService)(nil).AcceptInvitation), ctx, invitationID)
}

// GetAllPendingInvitations mocks base method.
func (m *MockInvitationService) GetAllPendingInvitations(ctx context.Context, userID uuid.UUID) ([]models.Invitations, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPendingInvitations", ctx, userID)
	ret0, _ := ret[0].([]models.Invitations)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPendingInvitations indicates an expected call of GetAllPendingInvitations.
func (mr *MockInvitationServiceMockRecorder) GetAllPendingInvitations(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPendingInvitations", reflect.TypeOf((*MockInvitationService)(nil).GetAllPendingInvitations), ctx, userID)
}

// MakeInvitation mocks base method.
func (m *MockInvitationService) MakeInvitation(ctx context.Context, invitingUserID, invitedUserID, slotId uuid.UUID,gameId uuid.UUID) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeInvitation", ctx, invitingUserID, invitedUserID, slotId,gameId)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeInvitation indicates an expected call of MakeInvitation.
func (mr *MockInvitationServiceMockRecorder) MakeInvitation(ctx, invitingUserID, invitedUserID, slotId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeInvitation", reflect.TypeOf((*MockInvitationService)(nil).MakeInvitation), ctx, invitingUserID, invitedUserID, slotId)
}

// RejectInvitation mocks base method.
func (m *MockInvitationService) RejectInvitation(ctx context.Context, invitationID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RejectInvitation", ctx, invitationID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RejectInvitation indicates an expected call of RejectInvitation.
func (mr *MockInvitationServiceMockRecorder) RejectInvitation(ctx, invitationID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RejectInvitation", reflect.TypeOf((*MockInvitationService)(nil).RejectInvitation), ctx, invitationID)
}
