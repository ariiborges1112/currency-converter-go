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

var moedasPermitidas = make(map[string] int)

type ExchangeRateResponse struct{
	Rates map[string]float64 `json:"rates"`
}

func main(){
	moedaOrigem, moedaDestino, valor := entradaDeDados()

	
	
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

func listarMoedas(){
	fmt.Print("\n---LISTA DE MOEDAS---")
	fmt.Print("[BRL] Real Brasileiro")
	fmt.Print("[USD] Dólar Americano")
	fmt.Print("[EUR] Euro")
	fmt.Print("[GBP] Libra Esterlina")
	
	if len(moedasPermitidas) == 0{
		moedasPermitidas["BRL"] = 1
		moedasPermitidas["USD"] = 2
		moedasPermitidas["EUR"] = 3
		moedasPermitidas["GBP"] = 4
	}
}

func listarOutrasMoedas(){
	fmt.Print("\n---OUTRAS MOEDAS---")
	fmt.Print("[CAD] Dólar Canadense")
	fmt.Print("[AUD] Dólar Australiano")
	fmt.Print("[JPY] Iene Japonês")
	fmt.Print("[CHF] Franco Suiço")
	fmt.Print("[CNY] Yuan Chinês")
	fmt.Print("[ARS] Peso Argentino")
	fmt.Print("[CLP] Peso Chileno")
	fmt.Print("[UYU] Peso Uruguaio")

	if len(moedasPermitidas) == 4{
		moedasPermitidas["CAD"] = 5
		moedasPermitidas["AUD"] = 6
		moedasPermitidas["JPY"] = 7
		moedasPermitidas["CHF"] = 8
		moedasPermitidas["CNY"] = 9
		moedasPermitidas["ARS"] = 10
		moedasPermitidas["CLP"] = 11
		moedasPermitidas["UYU"] = 12
	}
}

func entradaDeDados() (string, string, float64){
	listarMoedas()

	var escolha string
	fmt.Print("Deseja listar os outros tipos de moedas? [sim/nao] ")
	fmt.Scanln(&escolha)

	if simOuNao(escolha){
		listarOutrasMoedas()
	}

	var moedaOrigem string

	for{
		fmt.Print("\nDigite o moeda de origem para conversão: [BRL, USD, EUR...]: ")
		fmt.Scanln(&moedaOrigem)
		moedaOrigem = strings.ToUpper(moedaOrigem)
		
		if _, existe := moedasPermitidas[moedaOrigem]; !existe{
			fmt.Printf("A moeda [%s] não foi encontrada, digite novamente!", moedaOrigem)
			continue
		}

		break
	}

	var valor float64

	for{
		fmt.Printf("Digite o valor a ser comparado da moeda [%s]: ", moedaOrigem)
		fmt.Scanln(&valor)

		if valor <= 0.0{
			fmt.Print("O valor não pode ser igual ou menor que 0!")
			continue
		}

		break
	}

	var moedaDestino string

	for{
		fmt.Print("\nDigite o moeda de destino para conversão: [BRL, USD, EUR...]: ")
		fmt.Scanln(&moedaDestino)
		moedaDestino = strings.ToUpper(moedaDestino)

		if _, existe := moedasPermitidas[moedaDestino]; !existe{
			fmt.Printf("A moeda [%s] não foi encontrada, digite novamente!", moedaDestino)
			continue
		}

		if moedaDestino == moedaOrigem{
			fmt.Printf("A moeda de comparação não pode ser a mesma a ser comparada! Ambas são [%s]!", moedaOrigem)
			continue
		}

		break
	}

	return moedaOrigem, moedaDestino, valor
}

func limparTerminal(){
	var cmd *exec.Cmd
	
	if runtime.GOOS == "windowns"{
		cmd = exec.Command("cmd", "/c", "cls")
	}else{
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func simOuNao(escolha string) bool{
	strings.ToLower(escolha)

		switch escolha{
			case "sim", "s":
				return true
			case "nao", "não", "nn", "n":
				return false
			default:
				fmt.Print("Erro: Escolha errada, digite novamente!")
		}
}