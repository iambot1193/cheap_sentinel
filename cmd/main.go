package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"cheap_sentinel/internal" // Certifica-te que este caminho bate com o teu go.mod

	"github.com/playwright-community/playwright-go"
)

func main() {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("❌ Erro Playwright: %v", err)
	}

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	if err != nil {
		log.Fatalf("❌ Erro Browser: %v", err)
	}
	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("❌ Erro Page: %v", err)
	}

	fmt.Println("🛰️ Sentinel acessando AliExpress...")
	url := "https://pt.aliexpress.com/item/1005006357879509.html"
	if _, err := page.Goto(url); err != nil {
		log.Fatalf("❌ Erro URL: %v", err)
	}

	time.Sleep(5 * time.Second)
	title, _ := page.Title()
	words := strings.Fields(title)
	if len(words) > 5 {
		title = strings.Join(words[:5], " ") + "..."
	}

	priceLocator := page.Locator("span[class*='price-default--current']")
	if err := priceLocator.WaitFor(); err != nil {
		page.Screenshot(playwright.PageScreenshotOptions{Path: playwright.String("erro_preco.png")})
		return
	}

	rawPrice, _ := priceLocator.InnerText()
	cleanPrice := strings.NewReplacer("R$", "", " ", "", ".", "", ",", ".").Replace(rawPrice)
	priceNumeric, err := strconv.ParseFloat(strings.TrimSpace(cleanPrice), 64)

	if err != nil {
		log.Printf("⚠️ Erro na conversão: %v", err)
	} else {
		fmt.Printf("🔢 Valor capturado: %.2f\n", priceNumeric)
		csvPath := "data/prices.csv"

		// --- ADAPTAÇÃO: LÓGICA INVERTIDA ---

		// 1. Calcula estatísticas do passado ANTES de guardar o novo preço
		avgPrice, minPrice, err := internal.CalculateStats(csvPath)

		score := 0.0
		if err == nil && avgPrice > 0 {
			// Só calcula se houver diferença entre média e mínima para evitar div/0
			if (avgPrice - minPrice) > 0.01 {
				score = ((avgPrice - priceNumeric) / (avgPrice - minPrice)) * 100
			} else if priceNumeric < avgPrice {
				// Caso simplificado: se o preço baixou mas não temos range histórico
				score = 50.0
			}
		}

		fmt.Printf("📊 Opportunity Score: %.2f\n", score)

		// 2. Agora sim, guarda a nova entrada com o Score calculado
		err = internal.SaveEntry(csvPath, title, priceNumeric, score)
		if err != nil {
			log.Printf("❌ Erro ao guardar: %v", err)
		} else {
			fmt.Println("💾 Dados persistidos com sucesso!")
		}
	}

	page.Screenshot(playwright.PageScreenshotOptions{Path: playwright.String("prova_de_vida.png")})
	pw.Stop()
}
