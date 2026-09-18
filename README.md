# 💱 Currency Converter CLI

[![Go Version](https://img.shields.io/badge/Go-1.20%2B-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

Um conversor de moedas dinâmico em linha de comando (CLI) desenvolvido em **Go (Golang)**. O projeto consome taxas de câmbio em tempo real diretamente de uma API REST pública e realiza conversões entre **BRL**, **USD** e **EUR** com validação estrita de dados.

---

## 🚀 Funcionalidades

- **Cotação em Tempo Real**: Consulta a API pública de câmbio para obter taxas atualizadas na execução.
- **Resiliência e Timeout**: Utiliza um cliente HTTP customizado com tempo limite (*timeout*) para evitar travamentos em caso de instabilidade na internet.
- **Validação Estrita de Entrada**: 
  - Tratamento contra inputs inválidos (letras onde deveriam ser números, valores menores ou iguais a zero).
  - Bloqueio automático caso o usuário tente converter uma moeda para ela mesma.
- **Listagem Flexível**: Exibe as principais moedas logo no início e permite ao usuário consultar uma lista expandida sob demanda.
- **Fluxo Contínuo e Interativo**: Opção de tentar novamente em caso de falha de conexão ou realizar novas conversões em loop sem precisar reiniciar o programa manualmente.
- **Limpeza de Terminal Nativa**: Compatível com múltiplos sistemas operacionais (`Windows` e sistemas Unix/Linux/macOS).

---

## 🛠️ Tecnologias Utilizadas

- **[Go (Golang)](https://go.dev/)**: Linguagem principal do projeto.
- **Pacotes Nativos**:
  - `net/http` & `encoding/json`: Para requisições HTTP seguras e desserialização de JSON.
  - `os/exec` & `runtime`: Para manipulação e limpeza do terminal baseada no S.O.
  - `time`: Para o gerenciamento de *timeouts* de rede.
- **Boas Práticas de Engenharia**:
  - Modularização de código em funções especializadas.
  - Tratamento explícito de erros(da linguagem Go).
  - Gerenciamento de escopo e controle de fluxo avançado (`for`, `switch`, `goto`).
- **ExchangeRate-API**: API REST pública utilizada para obter as cotações financeiras.

---

## 🔧 Como Executar

### Pré-requisitos
- **Go** instalado em sua máquina (versão 1.18 ou superior).

### Passo a Passo

1. Clone este repositório:
   ```bash
   git clone [https://github.com/ariiborges1112/currency-converter-go.git](https://github.com/ariiborges1112/currency-converter-go.git)