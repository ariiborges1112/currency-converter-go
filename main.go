package main

import(
	"fmt"
	"encoding/json"
	"net/http"
	"strings"
)

type ExchangeRateResponse struct{
	Rates map[string]float64 `json:"rates"`
}

func moedaValida(moeda string) bool{
	moedasPermitidas := []string{"BRL", "USD", "EUR"}
	for _, m := range moedasPermitidas{
		if m == moeda{
			return true
		}
	}
	return false
}

func main(){
	var moedaOrigem string
	fmt.Print("Digite o moeda de origem para conversão: [BRL, USD, EUR]: ")
	fmt.Scanln(&moedaOrigem)
	moedaOrigem = strings.ToUpper(moedaOrigem)

	var valor float64
	fmt.Printf("Digite o valor da moeda %s: ", moedaOrigem)
	fmt.Scanln(&valor)

	var moedaDestino string
	fmt.Print("\nDigite o moeda de destino para conversão: [BRL, USD, EUR]: ")
	fmt.Scanln(&moedaDestino)
	moedaDestino = strings.ToUpper(moedaDestino)

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