# API Golang - Persons

## Descrição
Este projeto é uma API desenvolvida em Golang para gerenciar informações de pessoas. Ele fornece endpoints para criar, ler, atualizar e excluir registros de pessoas.

## Tecnologias Utilizadas
- **Golang**: Linguagem de programação principal.
- **HTTP Framework**: [Gorilla Mux](https://github.com/gorilla/mux) 
- **Banco de Dados**: PostgreSQL

## Endpoints
### Exemplos de Endpoints
- `GET /persons`: Retorna a lista de todas as pessoas.
- `GET /persons/{id}`: Retorna os detalhes de uma pessoa específica.
- `POST /persons`: Cria um novo registro de pessoa.
- `PUT /persons/{id}`: Atualiza os dados de uma pessoa existente.
- `DELETE /persons/{id}`: Remove uma pessoa do sistema.

## Como Executar o Projeto
- WIP

### Pré-requisitos
- [Golang](https://golang.org/dl/) instalado na máquina.
- (Opcional) [Docker](https://www.docker.com/) para execução em container.

### Passos
1. Clone o repositório:
   ```bash
   git clone https://github.com/ln0rd/api-golang-persons.git
   ```
2. Instale dependencias
   ```
   go mod tidy
   ```



### Reminder

Init project, to reference the repository with the code
```
go mod init github.com/ln0rd/api-golang-persons
```




