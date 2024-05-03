package internal

import (
	"encoding/json"
	"net/http"
)

type BrasilApiModel struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

type ViaCepModel struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func GetCepBrasilApi(cep string) (*BrasilApiModel, error) {

	endereco, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
	if err != nil {
		panic(err)
	}
	defer endereco.Body.Close()

	var model BrasilApiModel
	err = json.NewDecoder(endereco.Body).Decode(&model)
	if err != nil {
		panic(err)
	}

	return &model, nil
}

func GetCepViaCep(cep string) (*ViaCepModel, error) {

	endereco, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		panic(err)
	}
	defer endereco.Body.Close()

	var model ViaCepModel
	err = json.NewDecoder(endereco.Body).Decode(&model)
	if err != nil {
		panic(err)
	}

	return &model, nil
}
