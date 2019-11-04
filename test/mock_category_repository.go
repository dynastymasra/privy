package test

import (
	"context"

	"github.com/dynastymasra/privy/domain"
	"github.com/stretchr/testify/mock"
)

type MockCategoryRepository struct {
	mock.Mock
}

func (m *MockCategoryRepository) Create(ctx context.Context, category domain.Category) (*domain.Category, error) {
	args := m.Called(ctx)
	return args.Get(0).(*domain.Category), args.Error(0)
}

func (m *MockCategoryRepository) FindByID(ctx context.Context, id int) (*domain.Category, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Category), args.Error(1)
}

func (m *MockCategoryRepository) Fetch(ctx context.Context, offset, limit int) ([]domain.Category, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Category), args.Error(1)
}

func (m *MockCategoryRepository) Update(ctx context.Context, category domain.Category) error {
	args := m.Called(ctx, category)
	return args.Error(0)
}

func (m *MockCategoryRepository) Delete(ctx context.Context, category domain.Category) error {
	args := m.Called(ctx, category)
	return args.Error(0)
}
