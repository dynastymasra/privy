package product_test

import (
	"context"
	"testing"

	"github.com/dynastymasra/privy/domain"
	"github.com/stretchr/testify/assert"

	"github.com/dynastymasra/privy/config"
	"github.com/dynastymasra/privy/product"
	"github.com/dynastymasra/privy/test"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite
	productRepo    *test.MockProductRepository
	productService *product.Service
}

func Test_ServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) SetupSuite() {
	config.SetupTestLogger()
}

func (s *ServiceSuite) SetupTest() {
	s.productRepo = &test.MockProductRepository{}
	productService := product.NewService(s.productRepo)
	s.productService = &productService
}

func (s *ServiceSuite) Test_Create_Success() {
	s.productRepo.On("Create", context.Background()).Return(&domain.Product{ID: 1}, nil)
	s.productRepo.On("FindByID", context.Background(), 1).Return(&domain.Product{}, nil)

	product, err := s.productService.Create(context.Background(), domain.Product{})

	assert.NotNil(s.T(), product)
	assert.NoError(s.T(), err)
}

func (s *ServiceSuite) Test_Create_Failed() {
	s.productRepo.On("Create", context.Background()).Return((*domain.Product)(nil), assert.AnError)

	product, err := s.productService.Create(context.Background(), domain.Product{})

	assert.Nil(s.T(), product)
	assert.Error(s.T(), err)
}

func (s *ServiceSuite) Test_Create_Failed_Find() {
	s.productRepo.On("Create", context.Background()).Return(&domain.Product{ID: 1}, nil)
	s.productRepo.On("FindByID", context.Background(), 1).Return((*domain.Product)(nil), assert.AnError)

	product, err := s.productService.Create(context.Background(), domain.Product{})

	assert.Nil(s.T(), product)
	assert.Error(s.T(), err)
}

func (s *ServiceSuite) Test_FindByID_Success() {
	s.productRepo.On("FindByID", context.Background(), 1).Return(&domain.Product{}, nil)

	product, err := s.productService.FindByID(context.Background(), 1)

	assert.NotNil(s.T(), product)
	assert.NoError(s.T(), err)
}

func (s *ServiceSuite) Test_FindByID_Failed() {
	s.productRepo.On("FindByID", context.Background(), 1).Return((*domain.Product)(nil), assert.AnError)

	product, err := s.productService.FindByID(context.Background(), 1)

	assert.Nil(s.T(), product)
	assert.Error(s.T(), err)
}

func (s *ServiceSuite) Test_Fetch_Success() {
	s.productRepo.On("Fetch", context.Background(), 0, 20).Return([]domain.Product{{ID: 1}}, nil)

	products, err := s.productService.Fetch(context.Background(), 0, 20)

	assert.NotEmpty(s.T(), products)
	assert.NoError(s.T(), err)
}

func (s *ServiceSuite) Test_Fetch_Failed() {
	s.productRepo.On("Fetch", context.Background(), 0, 20).Return([]domain.Product{}, assert.AnError)

	products, err := s.productService.Fetch(context.Background(), 0, 20)

	assert.Empty(s.T(), products)
	assert.Error(s.T(), err)
}

func (s *ServiceSuite) Test_Update_Success() {
	s.productRepo.On("FindByID", context.Background(), 1).Return(&domain.Product{ID: 1}, nil)
	s.productRepo.On("Update", context.Background(), domain.Product{ID: 1}).Return(nil)

	product, err := s.productService.Update(context.Background(), 1, domain.Product{})

	assert.NotNil(s.T(), product)
	assert.Equal(s.T(), 1, product.ID)
	assert.NoError(s.T(), err)
}

func (s *ServiceSuite) Test_Update_Failed_FindByID() {
	s.productRepo.On("Update", context.Background(), domain.Product{ID: 1}).Return(nil)
	s.productRepo.On("FindByID", context.Background(), 1).Return((*domain.Product)(nil), assert.AnError)

	product, err := s.productService.Update(context.Background(), 1, domain.Product{})

	assert.Nil(s.T(), product)
	assert.Error(s.T(), err)
}

func (s *ServiceSuite) Test_Update_Failed_Update() {
	s.productRepo.On("Update", context.Background(), domain.Product{ID: 1}).Return(assert.AnError)

	product, err := s.productService.Update(context.Background(), 1, domain.Product{})

	assert.Nil(s.T(), product)
	assert.Error(s.T(), err)
}

func (s *ServiceSuite) Test_Delete_Success() {
	s.productRepo.On("FindByID", context.Background(), 1).Return(&domain.Product{ID: 1}, nil)
	s.productRepo.On("Delete", context.Background(), domain.Product{ID: 1}).Return(nil)

	err := s.productService.Delete(context.Background(), 1)

	assert.NoError(s.T(), err)
}

func (s *ServiceSuite) Test_Delete_Failed_FindByID() {
	s.productRepo.On("FindByID", context.Background(), 1).Return((*domain.Product)(nil), assert.AnError)

	err := s.productService.Delete(context.Background(), 1)

	assert.Error(s.T(), err)
}

func (s *ServiceSuite) Test_Delete_Failed_Delete() {
	s.productRepo.On("FindByID", context.Background(), 1).Return(&domain.Product{ID: 1}, nil)
	s.productRepo.On("Delete", context.Background(), domain.Product{ID: 1}).Return(assert.AnError)

	err := s.productService.Delete(context.Background(), 1)

	assert.Error(s.T(), err)
}
