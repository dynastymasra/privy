package test

import (
	"context"

	"github.com/dynastymasra/privy/domain"
	"github.com/stretchr/testify/mock"
)

type MockCategoryService struct {
	mock.Mock
}

func (m *MockCategoryService) Create(ctx context.Context, category domain.Category) (*domain.Category, error) {
	args := m.Called(ctx)
	return args.Get(0).(*domain.Category), args.Error(1)
}

func (m *MockCategoryService) FindByID(ctx context.Context, id int) (*domain.Category, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Category), args.Error(1)
}

func (m *MockCategoryService) Fetch(ctx context.Context, from, size int) ([]domain.Category, error) {
	args := m.Called(ctx, from, size)
	return args.Get(0).([]domain.Category), args.Error(1)
}

func (m *MockCategoryService) Update(ctx context.Context, id int, category domain.Category) (*domain.Category, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Category), args.Error(1)
}

func (m *MockCategoryService) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
