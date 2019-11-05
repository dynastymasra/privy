package product_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jinzhu/gorm"

	"github.com/dynastymasra/privy/delivery/http/handler/product"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"

	"github.com/dynastymasra/privy/config"
	"github.com/dynastymasra/privy/test"
	"github.com/stretchr/testify/suite"
)

type DeleteSuite struct {
	suite.Suite
	productService *test.MockProductService
}

func Test_DeleteSuite(t *testing.T) {
	suite.Run(t, new(DeleteSuite))
}

func (p *DeleteSuite) SetupSuite() {
	config.SetupTestLogger()
}

func (p *DeleteSuite) SetupTest() {
	p.productService = &test.MockProductService{}
}

func (p *DeleteSuite) Test_ProductDelete_Success() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/products/%d", 1), nil)

	r = mux.SetURLVars(r, map[string]string{
		"product_id": "1",
	})

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	p.productService.On("Delete", ctx, 1).Return(nil)

	product.DeleteHandler(p.productService)(w, r.WithContext(ctx))

	assert.Equal(p.T(), http.StatusOK, w.Code)
}

func (p *DeleteSuite) Test_ProductDelete_Failed_NotFound() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/products/%d", 1), nil)

	r = mux.SetURLVars(r, map[string]string{
		"product_id": "1",
	})

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	p.productService.On("Delete", ctx, 1).Return(gorm.ErrRecordNotFound)

	product.DeleteHandler(p.productService)(w, r.WithContext(ctx))

	assert.Equal(p.T(), http.StatusNotFound, w.Code)
}

func (p *DeleteSuite) Test_ProductDelete_Failed() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/products/%d", 1), nil)

	r = mux.SetURLVars(r, map[string]string{
		"product_id": "1",
	})

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	p.productService.On("Delete", ctx, 1).Return(assert.AnError)

	product.DeleteHandler(p.productService)(w, r.WithContext(ctx))

	assert.Equal(p.T(), http.StatusInternalServerError, w.Code)
}

func (p *DeleteSuite) Test_ProductDelete_Failed_Params() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/products/%s", "s"), nil)

	r = mux.SetURLVars(r, map[string]string{
		"product_id": "s",
	})

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	product.DeleteHandler(p.productService)(w, r.WithContext(ctx))

	assert.Equal(p.T(), http.StatusBadRequest, w.Code)
}
