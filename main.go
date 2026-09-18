package main

import(
	"fmt"
	"encoding/json"
	"net/http"
	"strings"
	"os"
	"os/exec"
	"runtime"
	"time"
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

	for{
		sucesso := realizarConversao()

		if !sucesso{
			break
		}

		var escolha string

		for{
			fmt.Print("Deseja fazer outra conversão? [sim/nao]: ")
			fmt.Scanln(&escolha)

			if simOuNao(escolha) == "sim"{
				break
			}else if simOuNao(escolha) == "nao"{
				goto Finalizar
			}
		}
	}

	Finalizar:
		fmt.Println("\nFinalizando sistema...")
}

func limparTerminal(){
	var cmd *exec.Cmd

	if runtime.GOOS == "windows"{
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
			fmt.Printf("A moeda [%s] não foi encontrada, digite novamente!\n", moedaOrigem)
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
			fmt.Print("O valor não pode ser igual ou menor que 0!\n")
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
			fmt.Printf("A moeda [%s] não foi encontrada, digite novamente!\n", moedaDestino)
			continue
		}

		if moedaDestino == moedaOrigem{
			fmt.Printf("A moeda de comparação não pode ser a mesma a ser comparada! Ambas são [%s] %s!\n", moedaOrigem, nomeMoedaOrigem)
			continue
		}

		nomeMoedaDestino = moedasPermitidas[moedaDestino]
		break
	}

	return moedaOrigem, moedaDestino, valor, nomeMoedaOrigem, nomeMoedaDestino
}

func realizarConversao() bool{
	moedaOrigem, moedaDestino, valor, nomeMoedaOrigem, nomeMoedaDestino := entradaDeDados()
	
		fmt.Printf("\nConvertendo %.2f em %s para %s...\n", valor, nomeMoedaOrigem, nomeMoedaDestino)

		cliente := &http.Client{
			Timeout: 5 * time.Second,
		}

		url := fmt.Sprintf("https://open.er-api.com/v6/latest/%s", moedaOrigem)

		var resp *http.Response
		var err error

		for{
			resp, err = cliente.Get(url)

			if err != nil{
				fmt.Println("Erro de conexão ou o tempo esgotou:", err)
			
				var escolha string
				for{
					fmt.Print("\nDeseja tentar novamente? [sim/nao]: ")
					fmt.Scanln(&escolha)

					if simOuNao(escolha) == "sim"{
						break
					}else if simOuNao(escolha) == "nao"{
						return false
					}
				}
				continue
			}
			break
		}

		defer resp.Body.Close()

		var dados ExchangeRateResponse
		errDecodificacao := json.NewDecoder(resp.Body).Decode(&dados)

		if errDecodificacao != nil{
			fmt.Println("Erro ao traduzir o JSON da API:", errDecodificacao)
			return false
		}

		taxa := dados.Rates[moedaDestino]
		resultado := valor * taxa

		fmt.Printf("\nResultado da conversão é: %.2f\n", resultado)
		return true
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
	escolha = strings.ToLower(escolha)

		switch escolha{
			case "sim", "s":
				return "sim"
			case "nao", "não", "nn", "n":
				return "nao"
			default:
				fmt.Print("Erro: Escolha errada, digite novamente!\n")
				return "input inválido"
		}
}