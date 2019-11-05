package product_test

import (
	"context"
	"log"
	"testing"

	"github.com/jinzhu/gorm"

	"github.com/dynastymasra/privy/domain"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"

	"github.com/dynastymasra/privy/config"
	"github.com/dynastymasra/privy/infrastructure/database/postgres"
	"github.com/dynastymasra/privy/product"
	"github.com/stretchr/testify/suite"
)

type RepositorySuite struct {
	suite.Suite
	*product.Repository
}

func Test_RepositorySuite(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}

func (p *RepositorySuite) SetupSuite() {
	config.Load()
	config.SetupTestLogger()
}

func (p *RepositorySuite) TearDownSuite() {
	db, _ := postgres.Connect(config.Postgres())
	postgres.Close(db)
	postgres.Reset()
}

func (p *RepositorySuite) SetupTest() {
	db, err := postgres.Connect(config.Postgres())
	if err != nil {
		log.Fatal(err)
	}

	productRepo := product.NewRepository(db)

	p.Repository = productRepo
}

func genProduct() domain.Product {
	return domain.Product{
		Name:        fake.ProductName(),
		Description: fake.Paragraphs(),
		Enable:      true,
	}
}

func (p *RepositorySuite) Test_Create_Success() {
	res, err := p.Repository.Create(context.Background(), genProduct())

	assert.NotNil(p.T(), res)
	assert.Greater(p.T(), res.ID, 0)
	assert.NoError(p.T(), err)
}

func (p *RepositorySuite) Test_Create_Failed_Image() {
	prod := genProduct()
	prod.ImageIDs = []int{1000000000000}
	res, err := p.Repository.Create(context.Background(), prod)

	assert.Nil(p.T(), res)
	assert.Error(p.T(), err)
}

func (p *RepositorySuite) Test_Create_Failed_Category() {
	prod := genProduct()
	prod.CategoryIDs = []int{1000000000000}
	res, err := p.Repository.Create(context.Background(), prod)

	assert.Nil(p.T(), res)
	assert.Error(p.T(), err)
}

func (p *RepositorySuite) Test_Create_Failed() {
	testProd := genProduct()
	testProd.Name = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

	res, err := p.Repository.Create(context.Background(), testProd)

	assert.Nil(p.T(), res)
	assert.Error(p.T(), err)
}

func (p *RepositorySuite) Test_FindByID_Success() {
	prod := genProduct()

	res, _ := p.Repository.Create(context.Background(), prod)

	resp, err := p.Repository.FindByID(context.Background(), res.ID)

	assert.NotNil(p.T(), resp)
	assert.NoError(p.T(), err)
}

func (p *RepositorySuite) Test_FindByID_Failed() {
	resp, err := p.Repository.FindByID(context.Background(), 10000000000)

	assert.Nil(p.T(), resp)
	assert.Error(p.T(), err)
}

func (p *RepositorySuite) Test_Fetch_Success() {
	prod := genProduct()

	p.Repository.Create(context.Background(), prod)

	resp, err := p.Repository.Fetch(context.Background(), 0, 20)

	assert.NotEmpty(p.T(), resp)
	assert.NoError(p.T(), err)
}

func (p *RepositorySuite) Test_Update_Success() {
	prod := genProduct()

	res, _ := p.Repository.Create(context.Background(), prod)

	res.Name = "Update"
	err := p.Repository.Update(context.Background(), *res)

	assert.NoError(p.T(), err)
}

func (p *RepositorySuite) Test_Update_Failed_NotFound() {

	err := p.Repository.Update(context.Background(), domain.Product{ID: 10000000000})

	assert.Error(p.T(), err)
}

func (p *RepositorySuite) Test_Update_Failed_Image() {
	prod := genProduct()

	res, _ := p.Repository.Create(context.Background(), prod)

	res.Name = "Update"
	res.ImageIDs = []int{100000000000}
	err := p.Repository.Update(context.Background(), *res)

	assert.Error(p.T(), err)
}

func (p *RepositorySuite) Test_Update_Failed_Category() {
	prod := genProduct()

	res, _ := p.Repository.Create(context.Background(), prod)

	res.Name = "Update"
	res.CategoryIDs = []int{100000000000}
	err := p.Repository.Update(context.Background(), *res)

	assert.Error(p.T(), err)
}

func (p *RepositorySuite) Test_Delete_Success() {
	prod := genProduct()

	res, _ := p.Repository.Create(context.Background(), prod)

	err := p.Repository.Delete(context.Background(), *res)

	assert.NoError(p.T(), err)
}

func (p *RepositorySuite) Test_Delete_Failed() {
	err := p.Repository.Delete(context.Background(), domain.Product{ID: 10000000001})

	assert.Error(p.T(), err)
	assert.EqualError(p.T(), err, gorm.ErrRecordNotFound.Error())
}
