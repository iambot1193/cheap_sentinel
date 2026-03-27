# cheap_sentinel

> **Motor de monitoramento inteligente para e-commerce focado em análise de oportunidade e custo real.**

## 1. Visão Geral e Motivação
O **cheap_sentinel** é um motor de monitoramento de preços para e-commerce (foco inicial: AliExpress) desenvolvido em **Go**. Diferente de scrapers convencionais, o objetivo aqui é a extração de alta performance e a geração de um **Opportunity Score** baseado em histórico real, considerando variáveis críticas do mercado brasileiro como frete e impostos (**Remessa Conforme**).

## 2. Stack Tecnológica (The Architecture)
* **Linguagem:** `Go 1.2x` (Performance, Tipagem Estática, Concorrência).
* **Extração:** `PuerkitoBio/goquery` (Parsing HTML) ou `playwright-go` (para SPAs dinâmicos).
* **Storage:** Arquivo CSV versionado no GitHub (**Database-as-a-Git**).
* **Infra/Automação:** `GitHub Actions` (Cron-job para execução 24/7 sem custo).
* **Notificação:** `Email Bot API`.

## 3. Roadmap de Desenvolvimento (Milestones)

### Fase 1: Core Engine (Local)
- [ ] Configurar ambiente Go e estrutura de pastas.
- [ ] Implementar Scraper básico para AliExpress (parsing de Preço + Frete + Taxa).
- [ ] Criar módulo de persistência local em CSV.
- [ ] **Desafio Técnico:** Simular cabeçalhos de navegação (Cookies/User-Agent) para obter preços da malha brasileira.

### Fase 2: Data Intelligence & Logic
- [ ] Implementar o cálculo de Média Móvel (últimas 10-20 capturas).
- [ ] Desenvolver o algoritmo do **Opportunity Score (0-100)**.
- [ ] Criar lógica de exportação formatada para o GitHub (limpeza de dados e ordenação).

### Fase 3: Cloud & Automation
- [ ] Criar `workflow.yml` para o GitHub Actions.
- [ ] Configurar GitHub Secrets (Tokens e IDs).
- [ ] Implementar o **Committer**: O próprio binário em Go realiza o upload do CSV atualizado para o repositório via API.

### Fase 4: Observabilidade & Alerta
- [ ] Integração com Email Bot: Disparo de alertas se $Score > 80$.
- [ ] *(Opcional)* Conectar o CSV ao **Grafana Cloud** para visualização de Candlestick de preços.

## 4. Lógica de Oportunidade
O sistema utiliza uma fórmula de desvio para identificar anomalias positivas de preço:

$$Score = \left( \frac{Preço_{Médio} - Preço_{Atual}}{Preço_{Médio} - Preço_{Mínimo}} \right) \times 100$$

## 5. Modelo de Dados (CSV Structure)
Os dados são persistidos no arquivo `data/history.csv` seguindo a estrutura:

| timestamp | product_id | final_price_brl | tax_included | score | status |
| :--- | :--- | :--- | :--- | :--- | :--- |
| 2026-03-27T15:00Z | am4_b550_msi | 750.40 | true | 92 | 🔥 COMPRA |

## 6. Diferenciais de Engenharia (Portfolio Highlights)
* **Zero Infrastructure Cost:** Operação em nuvem gratuita via GitHub Actions.
* **Cross-Compilation:** Binários compilados para múltiplos ambientes (Windows/Linux/ARM).
* **Mathematical Insight:** Uso de estatística simples para tomada de decisão automatizada.
* **Clean Architecture:** Separação clara entre coleta de dados, processamento lógico e persistência.

---
Desenvolvido por [Felipe Gonçalves || github.com/iambot1193]
#Projeto iniciado na data de 27/03/2026