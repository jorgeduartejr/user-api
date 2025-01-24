package tests

import (
    "github.com/stretchr/testify/assert"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestCreateUser(t *testing.T) {
    app := SetupApp()

    req := httptest.NewRequest(http.MethodPost, "/users", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestGetUser(t *testing.T) {
    app := SetupApp()

    req := httptest.NewRequest(http.MethodGet, "/users/123", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestUpdateUser(t *testing.T) {
    app := SetupApp()

    req := httptest.NewRequest(http.MethodPut, "/users/123", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestDeleteUser(t *testing.T) {
    app := SetupApp()

    req := httptest.NewRequest(http.MethodDelete, "/users/123", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}