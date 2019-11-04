package test

import (
	"context"

	"github.com/dynastymasra/privy/domain"
	"github.com/stretchr/testify/mock"
)

type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) Create(ctx context.Context, product domain.Product) (*domain.Product, error) {
	args := m.Called(ctx)
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *MockProductService) FindByID(ctx context.Context, id int) (*domain.Product, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *MockProductService) Fetch(ctx context.Context, from, size int) ([]domain.Product, error) {
	args := m.Called(ctx, from, size)
	return args.Get(0).([]domain.Product), args.Error(1)
}

func (m *MockProductService) Update(ctx context.Context, id int, product domain.Product) (*domain.Product, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *MockProductService) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
