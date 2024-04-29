package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/weather/:cep", WeatherHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

type CEP struct {
	Localidade string `json:"localidade"`
}

type WeatherResponse struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}


func WeatherHandler(c echo.Context) error {
	cep := c.Param("cep")

	if len(cep) != 8 {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"message": "invalid zipcode"})
	}

	cepInfo, err := fetchCEPInfo(cep)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "can not find zipcode"})
	}

	weatherInfo, err := fetchWeatherInfo(cepInfo.Localidade)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	tempC := weatherInfo.TempC
	tempF := weatherInfo.TempF
	tempK := tempC + 273.15

	response := WeatherResponse{
		TempC: tempC,
		TempF: tempF,
		TempK: tempK,
	}

	return c.JSON(http.StatusOK, response)
}

func fetchCEPInfo(cep string) (*CEP, error) {
	resp, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var cepInfo CEP
	if err := json.NewDecoder(resp.Body).Decode(&cepInfo); err != nil {
		return nil, err
	}

	return &cepInfo, nil
}

func fetchWeatherInfo(cityName string) (*WeatherResponse, error) {
	apiKey := "2f9cee46bbca477b930110109242904"

	cityName = strings.ReplaceAll(cityName, " ", "%")

	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", apiKey, cityName)

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResponse struct {
		Current struct{
			TempC float64 `json:"temp_c"`
			TempF float64 `json:"temp_f"`
		} `json:"current"`
	}

	if err := json.Unmarshal(bodyBytes, &apiResponse); err != nil {
		return nil, err
	}

	return &WeatherResponse{
		TempC: apiResponse.Current.TempC,
		TempF: apiResponse.Current.TempF,
	}, nil
}


