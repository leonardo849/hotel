package integration_test

import (
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
    "hotel/internal/router"  
)

func TestHelloEndpoint(t *testing.T) {
    app := router.SetupApp()  

    req := httptest.NewRequest("GET", "/hello", nil)
    resp, err := app.Test(req)

    assert.NoError(t, err)
    assert.Equal(t, 200, resp.StatusCode)
}
