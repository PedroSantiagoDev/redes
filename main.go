package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type CompletionRequest struct {
	Prompt     string `json:"prompt"`
	MaxTokens  int    `json:"max_tokens"`
	Model      string `json:"model"`
	Token      string `json:"token"`
	Completion string `json:"completion"`
}

type CompletionResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text        string  `json:"text"`
		Rank        int     `json:"rank"`
		LogProbs    float64 `json:"logprobs"`
		FinishReason string  `json:"finish_reason"`
		Index       int     `json:"index"`
	} `json:"choices"`
}

func main() {
	// Configurar os parâmetros para a solicitação de conclusão da OpenAI API
	requestData := CompletionRequest{
		Prompt:     "Once upon a time",
		MaxTokens:  50,
		Model:      "text-davinci-003",
		Token:      "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",
		Completion: "",
	}

	// Codificar a estrutura de solicitação em JSON
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		log.Fatal("Falha ao codificar a solicitação JSON:", err)
	}

	// Fazer a solicitação HTTP para a API da OpenAI
	response, err := http.Post(
		"https://api.openai.com/v1/engines/davinci/completions",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		log.Fatal("Falha ao fazer a solicitação HTTP:", err)
	}

	// Verificar o código de status da resposta HTTP
	if response.StatusCode != http.StatusOK {
		log.Fatalf("Erro na resposta da API: %s", response.Status)
	}

	// Decodificar a resposta JSON em uma estrutura de resposta
	var responseData CompletionResponse
	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		log.Fatal("Falha ao decodificar a resposta JSON:", err)
	}

	// Exibir a resposta da API
	fmt.Println(responseData.Choices[0].Text)
}
