package test

import (
	"context"

	"github.com/dynastymasra/privy/domain"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) Create(ctx context.Context, product domain.Product) (*domain.Product, error) {
	args := m.Called(ctx)
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *MockProductRepository) FindByID(ctx context.Context, id int) (*domain.Product, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *MockProductRepository) Fetch(ctx context.Context, offset, limit int) ([]domain.Product, error) {
	args := m.Called(ctx, offset, limit)
	return args.Get(0).([]domain.Product), args.Error(1)
}

func (m *MockProductRepository) Update(ctx context.Context, product domain.Product) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}

func (m *MockProductRepository) Delete(ctx context.Context, product domain.Product) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}
