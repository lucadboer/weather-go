# Weather App

## Descrição
Este aplicativo fornece informações sobre o clima com base no CEP fornecido. Ele retorna a temperatura atual em graus Celsius, Fahrenheit e Kelvin.

## Pré-requisitos
- Go 1.x
- Docker (para uso com Docker)
- Chave de API do WeatherAPI
- Para utilizar este aplicativo, você precisa obter uma chave de API do WeatherAPI. Visite [WeatherAPI](https://www.weatherapi.com/) para obter sua chave.

## Como usar

### Sem Docker
1. Clone o repositório e navegue até a pasta do projeto.
2. Adicione sua chave de API do WeatherAPI ao código ou como uma variável de ambiente `WEATHER_API_KEY`.
3. Execute o comando `go run main.go` para iniciar o servidor.
4. Acesse `http://localhost:8080/weather/{cep}`, substituindo `{cep}` pelo CEP desejado, para obter as informações de clima.

### Com Docker
1. Clone o repositório e navegue até a pasta do projeto.
2. Adicione sua chave de API do WeatherAPI ao código ou configure-a no Dockerfile ou docker-compose.yml como uma variável de ambiente.
3. Construa a imagem Docker com o comando `docker build -t weather-app .`.
4. Inicie o container com o comando `docker-compose up`.
5. Acesse `http://localhost:8080/weather/{cep}`, substituindo `{cep}` pelo CEP desejado, para obter as informações de clima.

### Testando em Produção
Para testar em produção, você pode usar a URL fornecida: [https://temp-by-cep-zoqxxrkhyq-uc.a.run.app/weather/01153000](https://temp-by-cep-zoqxxrkhyq-uc.a.run.app/weather/01153000)

## Endpoints
- `GET /weather/:cep` - Retorna informações de clima para o CEP fornecido.

## Resposta de Exemplo
```json
{
  "temp_C": 22.0,
  "temp_F": 71.6,
  "temp_K": 295.15
}
```

## Tecnologias Utilizadas
- Go
- Echo Framework
- Docker
- API de CEP (viacep.com.br)
- API de Clima (weatherapi.com)
