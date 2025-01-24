package tests

import (
    "github.com/stretchr/testify/assert"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestGetAllProducts(t *testing.T) {
    app := SetupApp()

    req := httptest.NewRequest(http.MethodGet, "/products", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestCreateProduct(t *testing.T) {
    app := SetupApp()

    req := httptest.NewRequest(http.MethodPost, "/products", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestGetProduct(t *testing.T) {
    app := SetupApp()

    req := httptest.NewRequest(http.MethodGet, "/products/123", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestUpdateProduct(t *testing.T) {
    app := SetupApp()

    req := httptest.NewRequest(http.MethodPut, "/products/123", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestDeleteProduct(t *testing.T) {
    app := SetupApp()

    req := httptest.NewRequest(http.MethodDelete, "/products/123", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestGetProductInvalidID(t *testing.T) {
    app := SetupApp()

    req := httptest.NewRequest(http.MethodGet, "/products/invalid-id", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestUpdateProductInvalidID(t *testing.T) {
    app := SetupApp()

    req := httptest.NewRequest(http.MethodPut, "/products/invalid-id", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestDeleteProductInvalidID(t *testing.T) {
    app := SetupApp()

    req := httptest.NewRequest(http.MethodDelete, "/products/invalid-id", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestCreateProductInvalidData(t *testing.T) {
    app := SetupApp()

    req := httptest.NewRequest(http.MethodPost, "/products", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestUpdateProductInvalidData(t *testing.T) {
    app := SetupApp()

    req := httptest.NewRequest(http.MethodPut, "/products/123", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestDeleteProductInvalidData(t *testing.T) {
    app := SetupApp()

    req := httptest.NewRequest(http.MethodDelete, "/products/123", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestGetProductNotFound(t *testing.T) {
    app := SetupApp()

    req := httptest.NewRequest(http.MethodGet, "/products/123", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestUpdateProductNotFound(t *testing.T) {
    app := SetupApp()

    req := httptest.NewRequest(http.MethodPut, "/products/123", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDeleteProductNotFound(t *testing.T) {
    app := SetupApp()

    req := httptest.NewRequest(http.MethodDelete, "/products/123", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}