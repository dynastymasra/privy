package product_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dynastymasra/privy/delivery/http/handler/product"
	"github.com/dynastymasra/privy/domain"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"

	"github.com/dynastymasra/privy/config"
	"github.com/dynastymasra/privy/test"
	"github.com/stretchr/testify/suite"
)

type FetchSuite struct {
	suite.Suite
	productService *test.MockProductService
}

func Test_FetchSuite(t *testing.T) {
	suite.Run(t, new(FetchSuite))
}

func (p *FetchSuite) SetupSuite() {
	config.SetupTestLogger()
}

func (p *FetchSuite) SetupTest() {
	p.productService = &test.MockProductService{}
}

func (p *FetchSuite) Test_ProductFindByID_Success() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/products/%d", 1), nil)

	r = mux.SetURLVars(r, map[string]string{
		"product_id": "1",
	})

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	p.productService.On("FindByID", ctx, 1).Return(&domain.Product{}, nil)

	product.FindByIDHandler(p.productService)(w, r.WithContext(ctx))

	assert.Equal(p.T(), http.StatusOK, w.Code)
}

func (p *FetchSuite) Test_ProductFindByID_Failed_Params() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/products/%s", "s"), nil)

	r = mux.SetURLVars(r, map[string]string{
		"product_id": "s",
	})

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	product.FindByIDHandler(p.productService)(w, r.WithContext(ctx))

	assert.Equal(p.T(), http.StatusBadRequest, w.Code)
}

func (p *FetchSuite) Test_ProductFindByID_Failed_NotFound() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/products/%d", 1), nil)

	r = mux.SetURLVars(r, map[string]string{
		"product_id": "1",
	})

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	p.productService.On("FindByID", ctx, 1).Return((*domain.Product)(nil), gorm.ErrRecordNotFound)

	product.FindByIDHandler(p.productService)(w, r.WithContext(ctx))

	assert.Equal(p.T(), http.StatusNotFound, w.Code)
}

func (p *FetchSuite) Test_ProductFindByID_Failed() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/products/%d", 1), nil)

	r = mux.SetURLVars(r, map[string]string{
		"product_id": "1",
	})

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	p.productService.On("FindByID", ctx, 1).Return((*domain.Product)(nil), assert.AnError)

	product.FindByIDHandler(p.productService)(w, r.WithContext(ctx))

	assert.Equal(p.T(), http.StatusInternalServerError, w.Code)
}

func (p *FetchSuite) Test_ProductFindAll_Success() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/products?from=20&size=40", nil)

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	p.productService.On("Fetch", ctx, 20, 40).Return([]domain.Product{{ID: 1}}, nil)

	product.FindAllHandler(p.productService)(w, r.WithContext(ctx))

	assert.Equal(p.T(), http.StatusOK, w.Code)
}

func (p *FetchSuite) Test_ProductFindAll_Failed() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/products?from=20&size=40", nil)

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	p.productService.On("Fetch", ctx, 20, 40).Return(([]domain.Product)(nil), assert.AnError)

	product.FindAllHandler(p.productService)(w, r.WithContext(ctx))

	assert.Equal(p.T(), http.StatusInternalServerError, w.Code)
}
