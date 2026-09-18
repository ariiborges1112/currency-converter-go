package main

import(
	"fmt"
	"encoding/json"
	"net/http"
	"strings"
)

var moedasPermitidas = make(map[string] int)

type ExchangeRateResponse struct{
	Rates map[string]float64 `json:"rates"`
}

func main(){
	entradaDeDados()

	if !moedaValida(moedaOrigem) || !moedaValida(moedaDestino){
    	fmt.Println("Erro: Selecione apenas uma das moedas permitidas: [BRL, USD, EUR].")
    	return
	}

	fmt.Printf("\nConvertendo %.2f de %s para %s...\n", valor, moedaOrigem, moedaDestino)

	url := fmt.Sprintf("https://open.er-api.com/v6/latest/%s", moedaOrigem)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Erro ao buscar taxas de câmbio:", err)
		return
	}

	defer resp.Body.Close()

	var dados ExchangeRateResponse
	json.NewDecoder(resp.Body).Decode(&dados)

	taxa := dados.Rates[moedaDestino]
	
	resultado := valor * taxa

	fmt.Printf("Resultado da conversão é: %.2f", resultado)
}

