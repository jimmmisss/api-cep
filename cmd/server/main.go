package main

import (
	"github.com.br/jimmmisss/api/api-cep/internal/handler"
	"net/http"
)

func main() {

	http.HandleFunc("/{cep}", handler.BuscarCep)

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		panic(err)
	}

}
