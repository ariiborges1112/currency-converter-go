# 💱 Currency Converter CLI

[![Go Version](https://img.shields.io/badge/Go-1.20%2B-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

Um conversor de moedas dinâmico em linha de comando (CLI) desenvolvido em **Go (Golang)**. O projeto consome taxas de câmbio em tempo real diretamente de uma API REST pública e realiza conversões entre **BRL**, **USD** e **EUR** com validação estrita de dados.

---

## 🚀 Funcionalidades

- **Cotação em Tempo Real**: Consulta a API pública de câmbio para obter taxas atualizadas na execução.
- **Validação de Entradas**: Garante que apenas moedas suportadas (`BRL`, `USD`, `EUR`) sejam processadas.
- **Tratamento de Input**: Normalização automática de caracteres para maiúsculas (`strings.ToUpper`), permitindo entradas em minúsculo ou maiúsculo.
- **Tratamento Explicito de Erros**: Validação de falhas de conexão de rede, parse de JSON e entradas inválidas.

---

## 🛠️ Tecnologias Utilizadas

- **[Go (Golang)](https://go.dev/)**: Linguagem principal do projeto.
- **`net/http`**: Pacote nativo do Go para requisições HTTP e consumo de APIs.
- **`encoding/json`**: Mapeamento e deserialização de JSON utilizando *structs* e *struct tags*.
- **ExchangeRate-API**: API REST pública utilizada para obter as cotações financeiras.

---

## 🔧 Como Executar

### Pré-requisitos
- **Go** instalado em sua máquina (versão 1.18 ou superior).

### Passo a Passo

1. Clone este repositório:
   ```bash
   git clone [https://github.com/ariiborges1112/currency-converter-go.git](https://github.com/ariiborges1112/currency-converter-go.git)