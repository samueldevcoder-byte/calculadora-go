# Calculadora em Go

Uma calculadora de terminal desenvolvida em Go como projeto prático de aprendizado e desenvolvimento de habilidades em programação.

## Funcionalidades

- Adição, subtração, multiplicação e divisão.
- Suporte a números inteiros, negativos e decimais.
- Validação de entradas.
- Tratamento de divisão por zero.
- Testes automatizados.

## Tecnologias

- Go
- Pacote `testing` para testes automatizados
- Git (para controle de versão, quando configurado)

## Como executar

1. Instale o Go: https://go.dev/dl/
2. Clone ou baixe este repositório.
3. Abra o terminal na pasta do projeto.
4. Execute:

```bash
go run .
```

Digite a operação no formato `2*2`, `10/5` ou `2.5+3`.

## Como executar os testes

Na pasta do projeto, execute:

```bash
go test -v
```

Para verificar possíveis problemas no código:

```bash
go vet ./...
```

## Objetivo

Projeto desenvolvido para praticar fundamentos de Go, funções, tratamento de erros, validação de entradas e testes automatizados.