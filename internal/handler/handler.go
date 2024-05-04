package handler

import (
	"encoding/json"
	"fmt"
	"github.com.br/jimmmisss/api/api-cep/internal"
	"net/http"
	"strings"
	"time"
)

func BuscarCep(w http.ResponseWriter, r *http.Request) {
	cep := r.PathValue("cep")
	w.Header().Set("Content-Type", "application/json")

	cepCleaned := strings.ReplaceAll(cep, "-", "")
	cepCleaned = strings.ReplaceAll(cepCleaned, ".", "")

	brasilCepChannel := make(chan *internal.BrasilApiModel)
	viaCepChannel := make(chan *internal.ViaCepModel)

	go internal.GetCepBrasilApi(cep, brasilCepChannel)
	go internal.GetCepViaCep(cep, viaCepChannel)

	var brasilCep *internal.BrasilApiModel
	var viaCep *internal.ViaCepModel

	select {
	case resp := <-brasilCepChannel:
		fmt.Println("CEP encontrado na API BrasilAPI com melhor performance")
		brasilCep = resp
		fmt.Println(brasilCep)
		json.NewEncoder(w).Encode(brasilCep)
	case via := <-viaCepChannel:
		fmt.Println("CEP encontrado na API ViaCep com melhor performance")
		viaCep = via
		fmt.Println(viaCep)
		json.NewEncoder(w).Encode(viaCep)
	case <-time.After(1 * time.Second):
		fmt.Println("Tempo limite excedido ao buscar CEP")
	}
}
