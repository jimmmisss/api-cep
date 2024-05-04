package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
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

func GetCepBrasilApi(cep string, channel chan<- *BrasilApiModel) {
	start := time.Now()

	endereco, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
	if err != nil {
		fmt.Println("Erro ao criar requisição:", err)
		channel <- nil
		return
	}
	defer endereco.Body.Close()

	body, err := io.ReadAll(endereco.Body)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON do CEP Brasil:", err)
		channel <- nil
		return
	}

	var model BrasilApiModel
	err = json.Unmarshal(body, &model)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON do CEP Brasil:", err)
		channel <- nil
		return
	}

	elapsed := time.Since(start)
	fmt.Printf("API %s respondeu em %s\n", "CEP Brasil", elapsed)

	channel <- &model
}

func GetCepViaCep(cep string, channel chan<- *ViaCepModel) {
	start := time.Now()

	endereco, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		fmt.Println("Erro ao criar requisição:", err)
		channel <- nil
		return
	}
	defer endereco.Body.Close()

	body, err := io.ReadAll(endereco.Body)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON do Via CEP:", err)
		channel <- nil
		return
	}

	var model ViaCepModel
	err = json.Unmarshal(body, &model)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON do Via CEP:", err)
		channel <- nil
		return
	}

	elapsed := time.Since(start)
	fmt.Printf("API %s respondeu em %s\n", "Via CEP", elapsed)

	channel <- &model
}
