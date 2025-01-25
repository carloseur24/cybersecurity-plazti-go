package repository

import (
	"context"
	"github-tracker/github-tracker-app/repository/entity"

	"github.com/stretchr/testify/mock"
)

type MockCommit struct {
	*mock.Mock
	Commit
}

func (m *MockCommit) Insert(ctx context.Context, commit *entity.Commit) (err error) {
	result := m.Called(ctx, commit)
	return result.Error(0)
}

func (m *MockCommit) GetCommitByAuthorEmail(ctx context.Context, email string) (commits []entity.Commit, err error) {
	result := m.Called(ctx, email)
	return result.Get(0).([]entity.Commit), result.Error(1)
}
