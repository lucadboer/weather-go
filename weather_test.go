package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// Função de inicialização do servidor para testes
func setupServer() *echo.Echo {
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "App is running!")
	})
	e.GET("/weather/:cep", WeatherHandler)
	return e
}

func TestValidCEP(t *testing.T) {
	log.Println("Iniciando Teste: TestValidCEP")
	e := setupServer()

	req := httptest.NewRequest(http.MethodGet, "/weather/15910000", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/weather/:cep")
	c.SetParamNames("cep")
	c.SetParamValues("01153000")

	if assert.NoError(t, WeatherHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var response WeatherResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		if assert.NoError(t, err) {
			log.Printf("Resposta recebida: %+v\n", response)
			assert.Greater(t, response.TempC, 0.0)
			assert.Greater(t, response.TempF, 0.0)
			assert.Greater(t, response.TempK, 0.0)
		}
	} else {
		log.Println("Erro ao chamar WeatherHandler")
	}
	log.Println("Teste TestValidCEP finalizado")
}

func TestInvalidCEP(t *testing.T) {
	log.Println("Iniciando Teste: TestInvalidCEP")
	e := setupServer()

	req := httptest.NewRequest(http.MethodGet, "/weather/123", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/weather/:cep")
	c.SetParamNames("cep")
	c.SetParamValues("123")

	if assert.NoError(t, WeatherHandler(c)) {
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		if assert.NoError(t, err) {
			log.Printf("Resposta recebida: %+v\n", response)
			assert.Equal(t, "invalid zipcode", response["message"])
		}
	} else {
		log.Println("Erro ao chamar WeatherHandler")
	}
	log.Println("Teste TestInvalidCEP finalizado")
}

func TestNotFoundCEP(t *testing.T) {
	log.Println("Iniciando Teste: TestNotFoundCEP")
	e := setupServer()

	req := httptest.NewRequest(http.MethodGet, "/weather/00000000", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/weather/:cep")
	c.SetParamNames("cep")
	c.SetParamValues("00000000")

	if assert.NoError(t, WeatherHandler(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		if assert.NoError(t, err) {
			log.Printf("Resposta recebida: %+v\n", response)
			assert.Equal(t, "can not find zipcode", response["message"])
		}
	} else {
		log.Println("Erro ao chamar WeatherHandler")
	}
	log.Println("Teste TestNotFoundCEP finalizado")
}
