package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "modernc.org/sqlite"
)

// Estrutura para mapear a resposta da API de cotação.
// Structure to map the API response for the currency exchange rate.
type CotacaoResponse struct {
	USDBRL USDBRL `json:"USDBRL"`
}

// Estrutura para mapear a cotação do dólar.
// Structure to map the USD to BRL exchange rate.
type USDBRL struct {
	Bid string `json:"bid"`
}

func main() {
	// Abrindo a conexão com o banco de dados uma vez. A conexão será reutilizada.
	// Opening the database connection once. The connection will be reused.
	db, err := sql.Open("sqlite", "./cotacoes.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // Garantir que a conexão será fechada quando o programa terminar.
	// Ensure the connection will be closed when the program finishes.

	// Criando a tabela "cotacoes" no banco, caso não exista.
	// Creating the "cotacoes" table in the database, if it does not exist.
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS cotacoes (id INTEGER PRIMARY KEY, cotacao TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	// Definindo o handler para a rota /cotacao.
	// Setting the handler for the /cotacao route.
	http.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context() // Obtendo o contexto da requisição. // Getting the request context.

		// Obtendo a cotação atual do USD para BRL chamando a função fetchCotacao.
		// Fetching the current USD to BRL exchange rate by calling the fetchCotacao function.
		cotacao, err := fetchCotacao(ctx)
		if err != nil {
			// Se ocorrer um erro ao buscar a cotação, retorna um erro HTTP 500.
			// If an error occurs when fetching the exchange rate, return HTTP 500 error.
			http.Error(w, fmt.Sprintf("Erro ao obter cotação: %v", err), http.StatusInternalServerError)
			return
		}

		// Salvando a cotação no banco de dados.
		// Saving the exchange rate to the database.
		if err := saveCotacao(ctx, db, cotacao); err != nil {
			// Se ocorrer um erro ao salvar a cotação no banco de dados, retorna um erro HTTP 500.
			// If an error occurs while saving the exchange rate to the database, return HTTP 500 error.
			http.Error(w, fmt.Sprintf("Erro ao salvar cotação: %v", err), http.StatusInternalServerError)
			return
		}

		// Definindo o tipo de conteúdo como JSON.
		// Setting the content type as JSON.
		w.Header().Set("Content-Type", "application/json")

		// Respondendo com a cotação obtida.
		// Responding with the obtained exchange rate.
		json.NewEncoder(w).Encode(map[string]string{"bid": cotacao})
	})

	// Iniciando o servidor HTTP na porta 8080.
	// Starting the HTTP server on port 8080.
	log.Println("Servidor iniciado na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Função para buscar a cotação do USD para BRL da API externa.
// Function to fetch the USD to BRL exchange rate from the external API.
func fetchCotacao(ctx context.Context) (string, error) {
	reqCtx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel() // Garantir que o timeout seja cancelado ao final.

	// Criando a requisição HTTP para a API externa.
	// Creating the HTTP request for the external API.
	req, err := http.NewRequestWithContext(reqCtx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req) // Enviando a requisição para a API.
	if err != nil {
		fmt.Printf("Erro ao chamar a API %v", err)
		return "", err
	}
	defer resp.Body.Close() // Garantir que o corpo da resposta será fechado após o uso.

	var cotacao CotacaoResponse
	// Fazendo o "decode" da resposta JSON da API.
	// Decoding the JSON response from the API.
	if err := json.NewDecoder(resp.Body).Decode(&cotacao); err != nil {
		fmt.Printf("Erro ao realizar o decode do payload %v", err)
		return "", err
	}

	// Retornando o valor da cotação (bid) obtido da API.
	// Returning the exchange rate value (bid) obtained from the API.
	return cotacao.USDBRL.Bid, nil
}

// Função para salvar a cotação no banco de dados.
// Function to save the exchange rate to the database.
func saveCotacao(ctx context.Context, db *sql.DB, cotacao string) error {
	dbCtx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel() // Garantir que o timeout seja cancelado ao final.

	// Preparando a consulta SQL para inserir a cotação.
	// Preparing the SQL statement to insert the exchange rate.
	stmt, err := db.PrepareContext(dbCtx, "INSERT INTO cotacoes (cotacao) VALUES (?)")
	if err != nil {
		return err
	}
	defer stmt.Close() // Garantir que o statement seja fechado após o uso.

	// Executando o comando SQL para inserir a cotação.
	// Executing the SQL command to insert the exchange rate.
	_, err = stmt.ExecContext(dbCtx, cotacao)
	return err
}
