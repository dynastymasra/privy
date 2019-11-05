package product_test

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dynastymasra/privy/delivery/http/handler/product"
	"github.com/dynastymasra/privy/domain"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"

	"github.com/dynastymasra/privy/config"
	"github.com/dynastymasra/privy/test"
	"github.com/stretchr/testify/suite"
)

type UpdateSuite struct {
	suite.Suite
	productService *test.MockProductService
}

func Test_UpdateSuite(t *testing.T) {
	suite.Run(t, new(UpdateSuite))
}

func (p *UpdateSuite) SetupSuite() {
	config.SetupTestLogger()
}

func (p *UpdateSuite) SetupTest() {
	p.productService = &test.MockProductService{}
}

func (p *UpdateSuite) Test_ProductUpdate_Success() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/v1/products/%d", 1), bytes.NewReader(productPayload()))
	r = mux.SetURLVars(r, map[string]string{
		"product_id": "1",
	})

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	p.productService.On("Update", ctx, 1).Return(&domain.Product{}, nil)

	product.UpdateHandler(p.productService)(w, r.WithContext(ctx))
	assert.Equal(p.T(), http.StatusOK, w.Code)
}

func (p *UpdateSuite) Test_ProductUpdate_Failed_ReadBody() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/v1/products/%d", 1), errReader(0))
	r = mux.SetURLVars(r, map[string]string{
		"product_id": "1",
	})

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	product.UpdateHandler(p.productService)(w, r.WithContext(ctx))
	assert.Equal(p.T(), http.StatusBadRequest, w.Code)
}

func (p *UpdateSuite) Test_ProductUpdate_Failed_Unmarshal() {
	reqInBytes := []byte(`<- test chan`)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/v1/products/%d", 1), bytes.NewReader(reqInBytes))
	r = mux.SetURLVars(r, map[string]string{
		"product_id": "1",
	})

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	product.UpdateHandler(p.productService)(w, r.WithContext(ctx))
	assert.Equal(p.T(), http.StatusBadRequest, w.Code)
}

func (p *UpdateSuite) Test_ProductUpdate_Failed_Validation() {
	reqInBytes := []byte(`{
		"name": "test"
	}`)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/v1/products/%d", 1), bytes.NewReader(reqInBytes))
	r = mux.SetURLVars(r, map[string]string{
		"product_id": "1",
	})

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	product.UpdateHandler(p.productService)(w, r.WithContext(ctx))
	assert.Equal(p.T(), http.StatusBadRequest, w.Code)
}

func (p *UpdateSuite) Test_ProductUpdate_Failed() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/v1/products/%d", 1), bytes.NewReader(productPayload()))
	r = mux.SetURLVars(r, map[string]string{
		"product_id": "1",
	})

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	p.productService.On("Update", ctx, 1).Return((*domain.Product)(nil), assert.AnError)

	product.UpdateHandler(p.productService)(w, r.WithContext(ctx))
	assert.Equal(p.T(), http.StatusInternalServerError, w.Code)
}

func (p *UpdateSuite) Test_ProductUpdate_Failed_Param() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/v1/products/%s", "s"), bytes.NewReader(productPayload()))
	r = mux.SetURLVars(r, map[string]string{
		"product_id": "s",
	})

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	product.UpdateHandler(p.productService)(w, r.WithContext(ctx))
	assert.Equal(p.T(), http.StatusBadRequest, w.Code)
}
