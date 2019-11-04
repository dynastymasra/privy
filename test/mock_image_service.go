package test

import (
	"context"

	"github.com/dynastymasra/privy/domain"
	"github.com/stretchr/testify/mock"
)

type MockImageService struct {
	mock.Mock
}

func (m *MockImageService) Create(ctx context.Context, image domain.Image) (*domain.Image, error) {
	args := m.Called(ctx)
	return args.Get(0).(*domain.Image), args.Error(1)
}

func (m *MockImageService) FindByID(ctx context.Context, id int) (*domain.Image, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Image), args.Error(1)
}

func (m *MockImageService) Fetch(ctx context.Context, from, size int) ([]domain.Image, error) {
	args := m.Called(ctx, from, size)
	return args.Get(0).([]domain.Image), args.Error(1)
}

func (m *MockImageService) Update(ctx context.Context, id int, image domain.Image) (*domain.Image, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Image), args.Error(1)
}

func (m *MockImageService) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
