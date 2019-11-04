package test

import (
	"context"

	"github.com/dynastymasra/privy/domain"
	"github.com/stretchr/testify/mock"
)

type MockImageRepository struct {
	mock.Mock
}

func (m *MockImageRepository) Create(ctx context.Context, image domain.Image) (*domain.Image, error) {
	args := m.Called(ctx)
	return args.Get(0).(*domain.Image), args.Error(0)
}

func (m *MockImageRepository) FindByID(ctx context.Context, id int) (*domain.Image, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Image), args.Error(1)
}

func (m *MockImageRepository) Fetch(ctx context.Context, offset, limit int) ([]domain.Image, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Image), args.Error(1)
}

func (m *MockImageRepository) Update(ctx context.Context, image domain.Image) error {
	args := m.Called(ctx, image)
	return args.Error(0)
}

func (m *MockImageRepository) Delete(ctx context.Context, image domain.Image) error {
	args := m.Called(ctx, image)
	return args.Error(0)
}
