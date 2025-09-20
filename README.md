# Sistema Antifraud - Detecção de Fraude em Transações

## 🎯 Objetivo

Desenvolver um sistema de detecção de fraude que analisa transações usando regras de negócio e padrões de comportamento, alertando sobre atividades suspeitas para uma plataforma admin.

## 📋 Contexto

No mercado financeiro moderno, sistemas de detecção de fraude precisam:

* Processar **milhares de transações por segundo** com alta performance
* Analisar **transações em tempo real** (menos de 50ms)
* Aplicar **regras de negócio** configuráveis para detectar fraudes
* Calcular **pontuações de risco** dinâmicas baseadas em padrões
* Enviar **alertas instantâneos** HTTP
* Manter **histórico de padrões** suspeitos e comportamentos
* Suportar **múltiplos cenários** de detecção simultaneamente
* Garantir **alta disponibilidade** e tolerância a falhas

## 🚀 Quick Start

### Pré-requisitos
- Go 1.21+
- PostgreSQL (opcional para histórico)
- Docker & Docker Compose (opcional)

### Instalação Rápida

```bash
# Clone o repositório
git clone <repository-url>
cd antifraud

# Execute com Make
make setup && make run-web

# Ou execute diretamente
go mod tidy
go run internal/services/web/cmd/main.go
```

### Acesso
- 📚 **API Base**: http://localhost:8080/api
- ❤️ **Health Check**: http://localhost:8080/api/health
- 🔍 **Análise de Transação**: POST http://localhost:8080/api/analyze
- 🚨 **Alertas**: http://localhost:8080/admin/alerts

### 🎯 O que implementar

Este é um **boilerplate** onde você deve implementar:

#### 1. **Fraud Engine** (`internal/services/engine/fraud/engine.go`)
- Análise em tempo real de transações
- Aplicação paralela de regras de detecção
- Cálculo de pontuação de risco combinada
- Decisão de aprovação/bloqueio automática

#### 2. **Rules Engine** (`internal/services/engine/rules/manager.go`)
- Aplicação dos 3 cenários de detecção
- Sistema de pesos e prioridades
- Análise de padrões de comportamento
- Detecção de anomalias temporais
- Reconhecimento de sequências suspeitas

#### 3. **Risk Calculator** (`internal/services/engine/risk/calculator.go`)
- Cálculo de pontuação de risco (0-100)
- Combinação de múltiplos fatores
- Classificação por níveis de risco
- Histórico de scores por usuário
- Armazenar o resultado do riscos

#### 6. **Handler Logic** (`internal/services/web/handlers/antifraud.go`)
- Implementar lógica real nos handlers
- Integrar com os services do engine
- Retornar JSON estruturado
- Tratamento de erros adequado

## 📊 Cenários de Detecção de Fraude

### 10 Cenários Implementados

| ID | Cenário | Descrição | Score Máximo | Criticidade |
|----|---------|-----------|--------------|-------------|
| 1 | **Viagem Impossível** | Transações em locais muito distantes | 90 | 🔴 CRÍTICA |
| 2 | **Valor Anômalo** | Transações muito acima do padrão | 80 | 🟠 ALTA |
| 3 | **Dispositivo Desconhecido** | Novo dispositivo suspeito | 70 | 🟡 MÉDIA |
| 4 | **Velocidade de Transações** | Muitas transações em pouco tempo | 85 | 🔴 CRÍTICA |
| 5 | **Horário Suspeito** | Transações em horários atípicos | 60 | 🟡 MÉDIA |
| 6 | **Sequência de Valores** | Padrões suspeitos de valores | 75 | 🟠 ALTA |
| 7 | **Localização Inconsistente** | IP vs GPS não batem | 80 | 🟠 ALTA |
| 8 | **Valor Redondo** | Transações com valores "perfeitos" | 65 | 🟡 MÉDIA |
| 9 | **Usuário Inativo** | Retorno após muito tempo | 70 | 🟡 MÉDIA |
| 10 | **Transações Consecutivas** | Mesmo valor repetido | 75 | 🟠 ALTA |

## Regras de Negócio Obrigatórias

### 1. Níveis de Risco

| Nível | Score | Cor | Ação Automática | Descrição |
|-------|-------|-----|-----------------|-----------|
| **LOW** | 0-30 | 🟢 | Aprovar | Transação normal, sem suspeitas |
| **MEDIUM** | 31-60 | 🟡 | Monitorar | Atenção, mas não bloquear |
| **HIGH** | 61-80 | 🟠 | Revisar | Requer análise humana |
| **CRITICAL** | 81-100 | 🔴 | Bloquear | Bloqueio automático |


### 2. Sistema de Alertas

**Prioridades de Fila**:
- **P0 (CRITICAL)**: Score 81-100 → Alerta imediato
- **P1 (HIGH)**: Score 61-80 → Alerta em 5s
- **P2 (MEDIUM)**: Score 31-60 → Alerta em 30s
- **P3 (LOW)**: Score 0-30 → Log apenas

## 📈 API Endpoints Obrigatórios

| Método | Endpoint | Descrição | Status Esperado |
|--------|----------|-----------|-----------------|
| POST | `/analyze` | Analisar transação | 200 (aprovada) / 400 (bloqueada) |
| GET | `/admin/alerts` | Listar alertas ativos | 200 |
| GET | `/admin/risk/{transactionId}` | Score específico | 200 / 404 |
| POST | `/rules` | Configurar regras | 201 / 400 |
| GET | `/stats` | Estatísticas de fraude | 200 |
| GET | `/health` | Health check | 200 |

## 🧪 Cenários de Teste Essenciais

### Executar Todos os Testes
```bash
make test
```

### 3 Cenários Obrigatórios
- ✅ **Transação aprovada** (score < 30) → Status 200
- ✅ **Transação bloqueada** (score > 80) → Status 400  
- ✅ **Alerta gerado** (score 61-80) → WebSocket notification

## 🏗️ Arquitetura do Sistema

### Estrutura do Projeto
```
antifraud/
├── internal/
│   ├── domain/                         # Entidades de negócio
│   │   ├── transaction.go             # Struct Transaction
│   │   ├── analysis.go                # Struct FraudAnalysis
│   │   ├── alert.go                   # Struct Alert
│   │   └── errors.go                  # Erros de domínio
│   ├── services/
│   │   ├── web/                       # Serviço Web (API REST)
│   │   │   ├── cmd/
│   │   │   │   └── main.go           # Entry point web
│   │   │   ├── handlers/
│   │   │   │   ├── antifraud.go      # Handlers de análise
│   │   │   │   ├── alerts.go         # Handlers de alertas
│   │   │   │   └── api.go            # Handlers gerais
│   │   ├── engine/                   # Serviço Engine (Fraude)
│   │   │   ├── cmd/
│   │   │   │   └── main.go          # Entry point engine
│   │   │   ├── fraud/
│   │   │   │   └── engine.go        # Fraud Engine
│   │   │   ├── risk/
│   │   │   │   └── calculator.go    # Risk Calculator
│   │   │   ├── alerts/
│   │   │   │   └── manager.go       # Alert Manager
│   │   └── shared/                  # Componentes compartilhados
│   │       ├── validators/
│   │       │   └── business.go      # Validações de negócio
│   │       └── config/
│   │           └── config.go        # Configurações
├── Makefile                         # Automação completa
├── docker-compose.yml               # Ambiente containerizado
└── go.mod                           # Dependências Go
```

## 🎮 Exemplo Prático

### Cenário: Detecção de Viagem Impossível

```bash
# 1. Usuário faz transação em São Paulo
curl -X POST http://localhost:8080/api/analyze \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user-123",
    "amount": 100.0,
    "type": "pix",
    "direction": "credito"
    "location": {
      "country": "BR",
      "city": "São Paulo", 
      "latitude": -23.5505,
      "longitude": -46.6333
    },
    "timestamp": "2024-01-01T10:00:00Z"
  }'

# Resposta: Transação aprovada
{
  "transaction_id": "tx-001",
  "risk_score": 15,
  "risk_level": "LOW",
  "approved": true,
  "alerts": []
}

# 2. Transação impossível em Nova York 30min depois
curl -X POST http://localhost:8080/api/analyze \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user-123",
    "amount": 200.0,
    "location": {
      "country": "US",
      "city": "New York",
      "latitude": 40.7128,
      "longitude": -74.0060
    },
    "timestamp": "2024-01-01T10:30:00Z"
  }'

# Resposta: Transação bloqueada
{
  "transaction_id": "tx-002", 
  "risk_score": 95,
  "risk_level": "CRITICAL",
  "approved": false,
  "reason": "Impossible travel detected",
  "alerts": [
    {
      "type": "IMPOSSIBLE_TRAVEL",
      "priority": "P0",
      "message": "Travel from São Paulo to New York in 30 minutes"
    }
  ]
}
```