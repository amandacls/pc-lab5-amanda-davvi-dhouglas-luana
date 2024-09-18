# Laboratório 05 - Programação Concorrente - 2024.1
### Grupo: Amanda Santana | Davvi Duarte | Dhouglas Bandeira | Luana Liberato

## Visão geral:
Este repositório apresenta nossa solução para o laboratório 5 da disciplina de Programação Concorrente para o semestre 2024.1

## Objetivo: 
Explorar uma construção essencial de golang para concorrência: select statements. Como motivação, iremos construir uma aplicação distribuída. Essa aplicação funciona como um buscador de arquivos para filesharing (pense em Bittorrent). 
Um cliente deve procurar que máquinas na rede armazenam o arquivo buscado. Como chave de busca, ele deve passar o hash do arquivo (calcular tal como o sum dos labs passados). Considere que haverá, pelo menos, um grupo de quatro máquinas.

## Solução:
Utilizamos um esquema de organização baseado na arquitetura cliente-servidor. Onde os papeis a eles designados são os seguintes:
- Cliente:
  1. Envia os hashs dos arquivos que ele possui
  2. Consulta quais máquinas possuem arquivos com um dado hash 
- Servidor:
  1. Salva os hashs em um mapa contendo: IP da máquina e os hashs dos arquivos que ela possui
  2. Consulta quais máquinas tem um determinado hash e envia a informação para o cliente

## Como executar:
Antes de tudo para gerar os arquivos execute o comando:

```bash make_dataset.sh <numero>```

Servidor:

```go run server.go```

Cliente:
1. Para cadastrar os arquivos, execute:

```go run clientRegister.go```

2. Para consultar um arquivo, execute:

```go run clientSearch.go <hash>```
