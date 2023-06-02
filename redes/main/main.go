package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)

func main() {
	// Carregar os dados de treinamento
	trainData, err := loadCSV("dados_treinamento.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Converter os dados para instâncias
	trainInstances, err := convertToInstances(trainData)
	if err != nil {
		log.Fatal(err)
	}

	// Criar um classificador KNN
	cls := knn.NewKnnClassifier("euclidean", "linear", 2)

	// Treinar o classificador com os dados de treinamento
	err = cls.Fit(trainInstances)
	if err != nil {
		log.Fatal(err)
	}

	// Carregar os dados de teste
	testData, err := loadCSV("dados_teste.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Converter os dados para instâncias
	testInstances, err := convertToInstances(testData)
	if err != nil {
		log.Fatal(err)
	}

	// Classificar as instâncias de teste
	predictions, err := cls.Predict(testInstances)
	if err != nil {
		log.Fatal(err)
	}

	// Calcular a acurácia da classificação
	confusionMatrix, err := evaluation.GetConfusionMatrix(testInstances, predictions)
	if err != nil {
		log.Fatal(err)
	}

	accuracy := evaluation.GetAccuracy(confusionMatrix)
	fmt.Printf("Acurácia: %.2f%%\n", accuracy*100)
}

// Função para carregar os dados de um arquivo CSV
func loadCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Função para converter os dados CSV em instâncias
func convertToInstances(data [][]string) (*base.DenseInstances, error) {
	// Criar uma nova matriz de atributos
	attributes := make([]base.Attribute, len(data[0])-1)
	for i := range attributes {
		attributes[i] = base.NewFloatAttribute(fmt.Sprintf("Attr%d", i))
	}

	// Criar a matriz de instâncias
	instances := base.NewDenseInstances()
	instances.AddAttributes(attributes...)

	// Converter os dados para instâncias
	for _, row := range data {
		instance := base.NewDenseInstance()

		// Converter cada valor para float64 e adicionar à instância
		for i, value := range row[:len(row)-1] {
			floatValue, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
			if err != nil {
				return nil, err
			}
			instance.SetAttributeValue(i, floatValue)
		}

		// Converter o valor alvo para um atributo nominal
		targetValue, err := strconv.Atoi(strings.TrimSpace(row[len(row)-1]))
		if err != nil {
			return nil, err
		}
		targetAttr := base.NewNominalAttribute("Ações", []string{"0", "1"})
		targetAttr.SetStringValue(targetAttr.GetStringValueFromSysVal(targetValue))
		instance.SetClassAttribute(targetAttr)

		// Adicionar a instância ao conjunto de instâncias
		instances.AddInstance(instance)
	}

	return instances, nil
}
