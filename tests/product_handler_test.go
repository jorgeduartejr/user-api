package tests

import (
    "user-api/database"
    "user-api/handlers"
    "github.com/gofiber/fiber/v2"
    "github.com/stretchr/testify/assert"
    "net/http"
    "net/http/httptest"
    "testing"
)

func setupApp() *fiber.App {
    app := fiber.New()

    // Conexão com o banco de dados de teste
    err := database.ConnectTestDB()
    if err != nil {
        panic(err)
    }

    // Inicializar a coleção de produtos
    handlers.InitProductCollection("testdb")

    // Configuração das rotas de produtos
    app.Get("/products", handlers.GetAllProducts)
    app.Get("/products/:id", handlers.GetProduct)
    app.Post("/products", handlers.CreateProduct)
    app.Put("/products/:id", handlers.UpdateProduct)
    app.Delete("/products/:id", handlers.DeleteProduct)

    return app
}

func TestGetAllProducts(t *testing.T) {
    app := setupApp()

    req := httptest.NewRequest(http.MethodGet, "/products", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestCreateProduct(t *testing.T) {
    app := setupApp()

    req := httptest.NewRequest(http.MethodPost, "/products", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestGetProduct(t *testing.T) {
    app := setupApp()

    req := httptest.NewRequest(http.MethodGet, "/products/123", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestUpdateProduct(t *testing.T) {
    app := setupApp()

    req := httptest.NewRequest(http.MethodPut, "/products/123", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestDeleteProduct(t *testing.T) {
    app := setupApp()

    req := httptest.NewRequest(http.MethodDelete, "/products/123", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestGetProductInvalidID(t *testing.T) {
    app := setupApp()

    req := httptest.NewRequest(http.MethodGet, "/products/invalid-id", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestUpdateProductInvalidID(t *testing.T) {
    app := setupApp()

    req := httptest.NewRequest(http.MethodPut, "/products/invalid-id", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestDeleteProductInvalidID(t *testing.T) {
    app := setupApp()

    req := httptest.NewRequest(http.MethodDelete, "/products/invalid-id", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestCreateProductInvalidData(t *testing.T) {
    app := setupApp()

    req := httptest.NewRequest(http.MethodPost, "/products", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestUpdateProductInvalidData(t *testing.T) {
    app := setupApp()

    req := httptest.NewRequest(http.MethodPut, "/products/123", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestDeleteProductInvalidData(t *testing.T) {
    app := setupApp()

    req := httptest.NewRequest(http.MethodDelete, "/products/123", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

