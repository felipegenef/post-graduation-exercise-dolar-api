package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Função para buscar a cotação do servidor HTTP local.
// Function to fetch the exchange rate from the local HTTP server.
func fetchCotacao(ctx context.Context) (string, error) {
	// Definindo o timeout para a requisição HTTP.
	// Setting the timeout for the HTTP request.
	reqCtx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel() // Garantir que o cancelamento do contexto ocorra no final.

	// Criando a requisição HTTP para o endpoint /cotacao.
	// Creating the HTTP request for the /cotacao endpoint.
	req, err := http.NewRequestWithContext(reqCtx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	// Enviando a requisição e obtendo a resposta.
	// Sending the request and receiving the response.
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close() // Garantir que o corpo da resposta seja fechado.

	// Verificando se o status da resposta é OK (200).
	// Checking if the response status is OK (200).
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("erro: status code %d", resp.StatusCode)
	}

	// Decodificando o corpo da resposta em formato JSON.
	// Decoding the response body from JSON format.
	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	// Retornando o valor da cotação (bid) obtido da resposta.
	// Returning the exchange rate (bid) value obtained from the response.
	return result["bid"], nil
}

func main() {
	// Criando um contexto para a requisição HTTP.
	// Creating a context for the HTTP request.
	ctx := context.Background()

	// Buscando a cotação usando a função fetchCotacao.
	// Fetching the exchange rate using the fetchCotacao function.
	cotacao, err := fetchCotacao(ctx)
	if err != nil {
		// Caso ocorra um erro ao obter a cotação, registrar no log e encerrar o programa.
		// If an error occurs while fetching the exchange rate, log it and terminate the program.
		log.Fatalf("Erro ao obter cotação: %v", err)
	}

	// Salvando a cotação em um arquivo de texto chamado cotacao.txt.
	// Saving the exchange rate to a text file named cotacao.txt.
	err = os.WriteFile("cotacao.txt", []byte(fmt.Sprintf("Dólar: %s", cotacao)), 0644)
	if err != nil {
		// Caso ocorra um erro ao salvar o arquivo, registrar no log e encerrar o programa.
		// If an error occurs while saving the file, log it and terminate the program.
		log.Fatalf("Erro ao salvar cotação no arquivo: %v", err)
	}

	// Registrando no log que a cotação foi salva com sucesso no arquivo.
	// Logging that the exchange rate has been successfully saved to the file.
	log.Println("Cotação salva em 'cotacao.txt'")
}
