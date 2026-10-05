**Aluno**: Matheus Figueiredo Silva (20260035032)

**Linguagem**: Go

# Jack Compiler

Projeto em Go com uma implementação parcial de um compilador para a linguagem
Jack, usada no curso Nand2Tetris.

## Estado atual

No momento, o projeto implementa a **análise léxica**: o scanner percorre o
código-fonte Jack e produz tokens para palavras-chave, identificadores,
constantes, símbolos e fim de arquivo. Também reconhece comentários e ignora
espaços em branco.

As etapas seguintes de um compilador — como análise sintática e geração de
código — ainda não estão implementadas. O executável atual é apenas um ponto de
entrada inicial.

## Requisitos

- Go 1.27 ou superior

## Testes

A validação é feita por testes Go. O teste do scanner processa os arquivos Jack
de exemplo em `tests/Square/` e compara os tokens gerados com os XMLs de
referência correspondentes.

Na raiz do repositório, execute todos os testes com:

```sh
go test ./...
```

Para executar somente os testes do scanner:

```sh
go test ./scanner
```

Para ver o resultado de cada caso de teste:

```sh
go test -v ./...
```

## Estrutura do projeto

- `scanner/`: scanner e testes da análise léxica.
- `token/`: tipos de token e formatação da representação XML.
- `tests/Square/`: fontes Jack e XMLs de referência usados nos testes.
- `cmd/jack_compiler/`: ponto de entrada do executável.

