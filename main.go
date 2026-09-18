package main

import(
	"fmt"
	"encoding/json"
	"net/http"
	"strings"
	"os"
	"os/exec"
	"runtime"
)

var moedasPermitidas = map[string] string{
	"BRL": "Real Brasileiro",
	"USD": "Dólar Americano",
	"EUR": "Euro",
	"GBP": "Libra Esterlina",
	"CAD": "Dólar Canadense",
	"AUD": "Dólar Australiano",
	"JPY": "Iene Japonês",
	"CHF": "Franco Suiço",
	"CNY": "Yuan Chinês",
	"ARS": "Peso Argentino",
	"CLP": "Peso Chileno",
	"UYU": "Peso Uruguaio",
}

type ExchangeRateResponse struct{
	Rates map[string]float64 `json:"rates"`
}

func main(){
	limparTerminal()
	moedaOrigem, moedaDestino, valor, nomeMoedaOrigem, nomeMoedaDestino := entradaDeDados()
	
	fmt.Printf("\nConvertendo %.2f em %s para %s...\n", valor, nomeMoedaOrigem, nomeMoedaDestino)

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

	fmt.Printf("\nResultado da conversão é: %.2f\n", resultado)
}

func limparTerminal(){
	var cmd *exec.Cmd

	if runtime.GOOS == "windons"{
		cmd = exec.Command("cmd", "/c", "cls")
	}else{
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func entradaDeDados() (string, string, float64, string, string){
	listarMoedas()

	var escolha string

	for{
		fmt.Print("\nDeseja listar os outros tipos de moedas? [sim/nao] ")
		fmt.Scanln(&escolha)

		if simOuNao(escolha) == "sim"{
			listarOutrasMoedas()
			break
		} else if simOuNao(escolha) == "nao"{
			break
		} else if simOuNao(escolha) == "input inválido"{
			continue
		}
	}

	var moedaOrigem string
	var nomeMoedaOrigem string

	for{
		fmt.Print("\nDigite o moeda de origem para conversão: [BRL, USD, EUR...]: ")
		fmt.Scanln(&moedaOrigem)
		moedaOrigem = strings.ToUpper(moedaOrigem)
		
		if _, existe := moedasPermitidas[moedaOrigem]; !existe{
			fmt.Printf("A moeda [%s] não foi encontrada, digite novamente!", moedaOrigem)
			continue
		}

		nomeMoedaOrigem = moedasPermitidas[moedaOrigem]
		break
	}

	var valor float64

	for{
		fmt.Printf("Digite o valor a ser comparado da moeda [%s] %s: ", moedaOrigem, nomeMoedaOrigem)
		fmt.Scanln(&valor)

		if valor <= 0.0{
			fmt.Print("O valor não pode ser igual ou menor que 0!")
			continue
		}

		break
	}

	var moedaDestino string
	var nomeMoedaDestino string

	for{
		fmt.Print("Digite o moeda de destino para conversão: [BRL, USD, EUR...]: ")
		fmt.Scanln(&moedaDestino)
		moedaDestino = strings.ToUpper(moedaDestino)

		if _, existe := moedasPermitidas[moedaDestino]; !existe{
			fmt.Printf("A moeda [%s] não foi encontrada, digite novamente!", moedaDestino)
			continue
		}

		if moedaDestino == moedaOrigem{
			fmt.Printf("A moeda de comparação não pode ser a mesma a ser comparada! Ambas são [%s] %s!", moedaOrigem, nomeMoedaOrigem)
			continue
		}

		nomeMoedaDestino = moedasPermitidas[moedaDestino]
		break
	}

	return moedaOrigem, moedaDestino, valor, nomeMoedaOrigem, nomeMoedaDestino
}

func listarMoedas(){
	fmt.Println("\n---LISTA DE MOEDAS---")
	fmt.Println("[BRL] Real Brasileiro")
	fmt.Println("[USD] Dólar Americano")
	fmt.Println("[EUR] Euro")
	fmt.Println("[GBP] Libra Esterlina")
}

func listarOutrasMoedas(){
	fmt.Println("\n---OUTRAS MOEDAS---")
	fmt.Println("[CAD] Dólar Canadense")
	fmt.Println("[AUD] Dólar Australiano")
	fmt.Println("[JPY] Iene Japonês")
	fmt.Println("[CHF] Franco Suiço")
	fmt.Println("[CNY] Yuan Chinês")
	fmt.Println("[ARS] Peso Argentino")
	fmt.Println("[CLP] Peso Chileno")
	fmt.Println("[UYU] Peso Uruguaio")
}

func simOuNao(escolha string) string{
	strings.ToLower(escolha)

		switch escolha{
			case "sim", "s":
				return "sim"
			case "nao", "não", "nn", "n":
				return "nao"
			default:
				fmt.Print("Erro: Escolha errada, digite novamente!")
				return "input inválido"
		}
}