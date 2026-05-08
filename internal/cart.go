package internal

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

// CalculateStats lê apenas o que já está no ficheiro
func CalculateStats(path string) (float64, float64, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, 0, err // Se o ficheiro não existe, é a primeira execução
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return 0, 0, err
	}

	var sum, min float64
	count := 0

	for i, record := range records {
		if i == 0 || len(record) < 2 {
			continue
		}
		// Preço está na segunda coluna (index 1)
		val, _ := strconv.ParseFloat(record[1], 64)

		sum += val
		if count == 0 || val < min {
			min = val
		}
		count++
	}

	if count == 0 {
		return 0, 0, nil
	}
	return sum / float64(count), min, nil
}

// SaveEntry adiciona a nova linha no fim do CSV
func SaveEntry(path string, name string, price float64, score float64) error {
	_, err := os.Stat(path)
	isNewFile := os.IsNotExist(err)

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if isNewFile {
		header := []string{"Produto", "Preço (R$)", "Data e Hora", "Score"}
		writer.Write(header)
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	row := []string{
		name,
		fmt.Sprintf("%.2f", price),
		timestamp,
		fmt.Sprintf("%.2f", score),
	}
	return writer.Write(row)
}
