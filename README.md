# Cotação API - Golang SQLite (Eng version below)

## Descrição

Este projeto faz parte de um exercício/exame de uma pós-graduação em Golang. Ele consiste em um servidor HTTP que consulta uma API externa para obter a cotação atual do dólar (USD) em relação ao real brasileiro (BRL) e armazena essa cotação em um banco de dados SQLite. O servidor disponibiliza um endpoint HTTP que pode ser acessado para obter a cotação mais recente.

## Funcionalidades

- Consulta a API externa para obter a cotação do dólar (USD) para o real (BRL).
- Armazena a cotação obtida no banco de dados SQLite.
- Endpoint HTTP (`/cotacao`) para consultar a cotação atual.
- Utiliza prepared statements para inserir dados no banco de dados com segurança.

## Requisitos

- Go 1.23.3 ou superior
- Pacote SQLite `modernc.org/sqlite`
- Acesso à internet para consumir a API externa (https://economia.awesomeapi.com.br/json/last/USD-BRL)

## Como Rodar o Servidor

### 1. Clone o Repositório

```bash
git clone https://github.com/felipegenef/post-graduation-exercise-dolar-api.git
cd cotacao-api
```

### 2. Instale as Dependências

Este projeto utiliza o pacote modernc.org/sqlite para interagir com o banco de dados SQLite. Você pode instalar as dependências executando:

```bash
go mod tidy
```

### 3. Execute o Servidor

Para executar o servidor, use o seguinte comando:

```bash
go run server/server.go
```

Isso iniciará o servidor HTTP na porta 8080.

### 4. Rodando o Cliente

Após o servidor estar rodando, você pode usar o cliente para buscar a cotação e salvá-la em um arquivo de texto.

Depois de executar o servidor, rode o cliente com o seguinte comando:

```bash
go run client/client.go
```
O cliente irá fazer uma requisição para o servidor, obter a cotação atual e salvar o valor em um arquivo chamado cotacao.txt.

### 5. Verifique o Arquivo

O cliente cria um arquivo chamado cotacao.txt contendo a cotação do dólar. O arquivo estará localizado no mesmo diretório onde o cliente foi executado.

Exemplo de conteúdo do arquivo cotacao.txt:

```txt
Dólar: 5.25
```
# Exchange Rate API - Golang SQLite

## Description

This project is part of an exercise/exam for a Golang postgraduate course. It consists of an HTTP server that queries an external API to get the current exchange rate of USD to BRL and stores this rate in a SQLite database. The server exposes an HTTP endpoint that can be accessed to retrieve the latest exchange rate.

## Features

- Queries an external API to get the USD to BRL exchange rate.
- Stores the retrieved exchange rate in a SQLite database.
- HTTP endpoint (/cotacao) to query the current exchange rate.
- Uses prepared statements to insert data into the database securely.

## Requirements

- Go 1.23.3 or higher
- SQLite package modernc.org/sqlite
- Internet access to query the external API (https://economia.awesomeapi.com.br/json/last/USD-BRL)

## How to Run the Server

### 1. Clone the Repository

```bash
git clone https://github.com/felipegenef/post-graduation-exercise-dolar-api.git
cd cotacao-api
```

### 2. Install Dependencies

This project uses the modernc.org/sqlite package to interact with the SQLite database. You can install the dependencies by running:

```bash
go mod tidy
```

### 3. Run the Server

To run the server, use the following command:

```bash
go run server/server.go
```
This will start the HTTP server on port 8080.

### 4. Running the Client

After running the server, you can run the client to fetch the exchange rate and save it in a text file.

Once the server is running, execute the client with the following command:

```bash
go run client/client.go
```

The client will make a request to the server, retrieve the current exchange rate, and save the value in a file called cotacao.txt.

### 5. Check the File

The client creates a file called cotacao.txt containing the exchange rate. The file will be located in the same directory where the client was executed.

Example content of the cotacao.txt file:

```txt
Dólar: 5.25
```
