package product_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dynastymasra/privy/delivery/http/handler/product"
	"github.com/dynastymasra/privy/domain"
	uuid "github.com/satori/go.uuid"

	"github.com/dynastymasra/privy/config"
	"github.com/dynastymasra/privy/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CreateSuite struct {
	suite.Suite
	productService *test.MockProductService
}

func Test_CreateSuite(t *testing.T) {
	suite.Run(t, new(CreateSuite))
}

func (p *CreateSuite) SetupSuite() {
	config.SetupTestLogger()
}

func (p *CreateSuite) SetupTest() {
	p.productService = &test.MockProductService{}
}

type errReader int

func (errReader) Read([]byte) (n int, err error) {
	return 0, assert.AnError
}

func productPayload() []byte {
	return []byte(`{
		"name": "Name",
		"description": "Description",
		"enable": true,
		"images": [
			1
		],
		"categories": [
			1
		]
	}`)
}

func (p *CreateSuite) Test_ProductCreate_Success() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/v1/products", bytes.NewReader(productPayload()))

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	p.productService.On("Create", ctx).Return(&domain.Product{}, nil)

	product.CreateHandler(p.productService)(w, r.WithContext(ctx))
	assert.Equal(p.T(), http.StatusCreated, w.Code)
}

func (p *CreateSuite) Test_ProductCreate_Failed_ReadBody() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/v1/products", errReader(0))

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	product.CreateHandler(p.productService)(w, r.WithContext(ctx))
	assert.Equal(p.T(), http.StatusBadRequest, w.Code)
}

func (p *CreateSuite) Test_ProductCreate_Failed_Unmarshal() {
	reqInBytes := []byte(`<- test chan`)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/v1/products", bytes.NewReader(reqInBytes))

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	product.CreateHandler(p.productService)(w, r.WithContext(ctx))
	assert.Equal(p.T(), http.StatusBadRequest, w.Code)
}

func (p *CreateSuite) Test_ProductCreate_Failed_Validation() {
	reqInBytes := []byte(`{
		"name": "test"
	}`)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/v1/products", bytes.NewReader(reqInBytes))

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	product.CreateHandler(p.productService)(w, r.WithContext(ctx))
	assert.Equal(p.T(), http.StatusBadRequest, w.Code)
}

func (p *CreateSuite) Test_ProductCreate_Failed() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/v1/products", bytes.NewReader(productPayload()))

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	p.productService.On("Create", ctx).Return((*domain.Product)(nil), assert.AnError)

	product.CreateHandler(p.productService)(w, r.WithContext(ctx))
	assert.Equal(p.T(), http.StatusInternalServerError, w.Code)
}
