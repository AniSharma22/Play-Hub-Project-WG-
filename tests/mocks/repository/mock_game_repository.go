// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\Dell\projects2\watchguard-proj\Project-WG\play-hub\internal\domain\interfaces\repository\game_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	entities "project2/internal/domain/entities"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockGameRepository is a mock of GameRepository interface.
type MockGameRepository struct {
	ctrl     *gomock.Controller
	recorder *MockGameRepositoryMockRecorder
}

// MockGameRepositoryMockRecorder is the mock recorder for MockGameRepository.
type MockGameRepositoryMockRecorder struct {
	mock *MockGameRepository
}

// NewMockGameRepository creates a new mock instance.
func NewMockGameRepository(ctrl *gomock.Controller) *MockGameRepository {
	mock := &MockGameRepository{ctrl: ctrl}
	mock.recorder = &MockGameRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGameRepository) EXPECT() *MockGameRepositoryMockRecorder {
	return m.recorder
}

// CreateGame mocks base method.
func (m *MockGameRepository) CreateGame(ctx context.Context, game *entities.Game) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGame", ctx, game)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGame indicates an expected call of CreateGame.
func (mr *MockGameRepositoryMockRecorder) CreateGame(ctx, game interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGame", reflect.TypeOf((*MockGameRepository)(nil).CreateGame), ctx, game)
}

// DeleteGame mocks base method.
func (m *MockGameRepository) DeleteGame(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGame", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGame indicates an expected call of DeleteGame.
func (mr *MockGameRepositoryMockRecorder) DeleteGame(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGame", reflect.TypeOf((*MockGameRepository)(nil).DeleteGame), ctx, id)
}

// FetchAllGames mocks base method.
func (m *MockGameRepository) FetchAllGames(ctx context.Context) ([]entities.Game, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAllGames", ctx)
	ret0, _ := ret[0].([]entities.Game)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAllGames indicates an expected call of FetchAllGames.
func (mr *MockGameRepositoryMockRecorder) FetchAllGames(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAllGames", reflect.TypeOf((*MockGameRepository)(nil).FetchAllGames), ctx)
}

// FetchGameByID mocks base method.
func (m *MockGameRepository) FetchGameByID(ctx context.Context, id uuid.UUID) (*entities.Game, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchGameByID", ctx, id)
	ret0, _ := ret[0].(*entities.Game)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchGameByID indicates an expected call of FetchGameByID.
func (mr *MockGameRepositoryMockRecorder) FetchGameByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchGameByID", reflect.TypeOf((*MockGameRepository)(nil).FetchGameByID), ctx, id)
}

// UpdateGameStatus mocks base method.
func (m *MockGameRepository) UpdateGameStatus(ctx context.Context, gameID uuid.UUID, status bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGameStatus", ctx, gameID, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGameStatus indicates an expected call of UpdateGameStatus.
func (mr *MockGameRepositoryMockRecorder) UpdateGameStatus(ctx, gameID, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGameStatus", reflect.TypeOf((*MockGameRepository)(nil).UpdateGameStatus), ctx, gameID, status)
}