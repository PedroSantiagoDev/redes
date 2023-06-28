package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/levigross/grequests"
	"github.com/shomali11/go-interview/openai"
)

type Monitor struct {
	ID   int    `json:"id"`
	Name string `json:"friendly_name"`
}

type UptimeRobotResponse struct {
	Monitors []Monitor `json:"monitors"`
}

func main() {
	// Configure a biblioteca da OpenAI
	err := api.Configure("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")
	if err != nil {
		log.Fatal("Falha ao configurar a API da OpenAI:", err)
	}

	// Obtenha os dados da API UptimeRobot
	monitors, err := fetchMonitorsFromUptimeRobot()
	if err != nil {
		log.Fatal("Falha ao obter dados da API UptimeRobot:", err)
	}

	// Gerar relatório da rede com base nos dados obtidos
	networkReport := generateNetworkReport(monitors)

	// Exibir o relatório da rede
	fmt.Println(networkReport)
}

func fetchMonitorsFromUptimeRobot() ([]Monitor, error) {
	// Faça a chamada à API UptimeRobot para obter os monitores
	resp, err := grequests.Get("https://api.uptimerobot.com/v2/getMonitors", &grequests.RequestOptions{
		Params: map[string]string{
			"api_key": "u2192484-8d91af3483cf40a9108a311d",
			"format":  "json",
		},
	})
	if err != nil {
		return nil, err
	}

	// Verifique o código de status da resposta
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Erro na resposta da API UptimeRobot: %d", resp.StatusCode)
	}

	// Parse a resposta JSON
	var uptimeRobotResp UptimeRobotResponse
	err = json.Unmarshal(resp.Bytes(), &uptimeRobotResp)
	if err != nil {
		return nil, err
	}

	// Retorne a lista de monitores obtida
	return uptimeRobotResp.Monitors, nil
}

func generateNetworkReport(monitors []Monitor) string {
	// Realize a análise dos monitores para gerar o relatório da rede

	// Exemplo de chamada à API da OpenAI para processamento de linguagem natural
	response, err := api.Complete(api.CompletionRequest{
		Model:  "text-davinci-003",
		Params: api.CompletionParams{Prompt: "Relatório da Rede\n\nFalhas e Melhorias:\n\n", MaxTokens: 100},
	})
	if err != nil {
		log.Fatal("Falha na chamada à API da OpenAI:", err)
	}

	// Construa o relatório da rede com base nos monitores e nas respostas da OpenAI
	report := response.Choices[0].Text

	return report
}
