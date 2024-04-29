package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestWeatherEndpoint(t *testing.T) {
	e := echo.New()
	e.GET("/weather/:cep", WeatherHandler)

	req := httptest.NewRequest(http.MethodGet, "/weather/15910000", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var responseMap map[string]interface{}

	err := json.Unmarshal(rec.Body.Bytes(), &responseMap)
	assert.NoError(t, err)

	_, tempCExists := responseMap["temp_C"]
	_, tempFExists := responseMap["temp_F"]
	_, tempKExists := responseMap["temp_K"]

	assert.True(t, tempCExists, "Campo 'temp_C' não encontrado")
	assert.True(t, tempFExists, "Campo 'temp_F' não encontrado")
	assert.True(t, tempKExists, "Campo 'temp_K' não encontrado")
}
