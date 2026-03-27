package main

import (
	"fmt"
	"log"
	"time"

	"github.com/playwright-community/playwright-go"
)

func main() {
	// 1. Inicia o driver do Playwright
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("Não foi possível iniciar o Playwright: %v", err)
	}

	// 2. Abre o navegador (Chromium)
	// Headless: false permite que você VEJA o navegador abrindo. 
	// Depois que o bot estiver pronto, mudamos para true.
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	if err != nil {
		log.Fatalf("Não foi possível abrir o navegador: %v", err)
	}
	defer browser.Close() // Garante que o navegador feche ao terminar

	// 3. Abre uma nova aba (página)
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("Erro ao criar nova página: %v", err)
	}

	// 4. Navega para um produto do AliExpress (exemplo de uma B550)
	fmt.Println("🛰️ Sentinel acessando AliExpress...")
	url := "https://www.aliexpress.com/item/1005006063640245.html" 
	
	if _, err = page.Goto(url); err != nil {
		log.Fatalf("Erro ao acessar a URL: %v", err)
	}

	// 5. Espera um pouco para o JavaScript carregar (o Ali é pesado)
	time.Sleep(5 * time.Second)

	// 6. Captura o título do produto para testar
	title, _ := page.Title()
	fmt.Printf("✅ Sucesso! Produto encontrado: %s\n", title)

	// 7. Tira um print de prova (vai salvar na raiz do projeto)
	if _, err = page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String("prova_de_vida.png"),
	}); err != nil {
		log.Fatalf("Erro ao tirar screenshot: %v", err)
	}

	fmt.Println("📸 Screenshot 'prova_de_vida.png' gerado com sucesso.")
}