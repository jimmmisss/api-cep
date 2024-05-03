package handler

import (
	"encoding/json"
	"fmt"
	"github.com.br/jimmmisss/api/api-cep/internal"
	"net/http"
	"strings"
)

func getCepBrasilHandler(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if cep == "" {
		http.Error(w, "cep is required", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	cepCleaned := strings.ReplaceAll(cep, "-", "")
	cepCleaned = strings.ReplaceAll(cepCleaned, ".", "")

	endereco, err := internal.GetCepBrasilApi(cepCleaned)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Sprintf("Endereço: %v", endereco)
	json.NewEncoder(w).Encode(endereco)
}

func getCepViaCepHandler(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if cep == "" {
		http.Error(w, "cep is required", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	cepCleaned := strings.ReplaceAll(cep, "-", "")
	cepCleaned = strings.ReplaceAll(cepCleaned, ".", "")

	endereco, err := internal.GetCepViaCep(cepCleaned)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Sprintf("Endereço: %v", endereco)
	json.NewEncoder(w).Encode(endereco)
}
