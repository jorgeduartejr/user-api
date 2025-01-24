package tests

import (
    "github.com/stretchr/testify/assert"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestCreateWallet(t *testing.T) {
    app := SetupApp()

    req := httptest.NewRequest(http.MethodPost, "/wallets", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestGetWallet(t *testing.T) {
    app := SetupApp()

    req := httptest.NewRequest(http.MethodGet, "/wallets/123", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestAddFunds(t *testing.T) {
    app := SetupApp()

    req := httptest.NewRequest(http.MethodPut, "/wallets/123/funds", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}