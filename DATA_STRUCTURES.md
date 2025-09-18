# 📊 Estruturas de Dados - Sistema Antifraude

Este documento algumas sugestões de estruturas de dados principais para implementação do sistema de detecção de fraude, organizadas por funcionalidade e complexidade.

## 🎯 Estruturas Principais

### Transaction Models

```go
type Transaction struct {
    ID            string            `json:"id" db:"id"`
    UserID        string            `json:"user_id" db:"user_id"`
    Amount        float64           `json:"amount" db:"amount"`
    Currency      string            `json:"currency" db:"currency"`
    Location      TransactionLocation `json:"location" db:"location"`
    DeviceInfo    DeviceInfo        `json:"device_info" db:"device_info"`
    Timestamp     time.Time         `json:"timestamp" db:"timestamp"`
    Type          TransactionType   `json:"type" db:"type"`
    MerchantInfo  MerchantInfo      `json:"merchant_info" db:"merchant_info"`
}

type TransactionLocation struct {
    Country     string  `json:"country" db:"country"`
    City        string  `json:"city" db:"city"`
    Latitude    float64 `json:"latitude" db:"latitude"`
    Longitude   float64 `json:"longitude" db:"longitude"`
    IPAddress   string  `json:"ip_address" db:"ip_address"`
}

type DeviceInfo struct {
    DeviceID    string `json:"device_id" db:"device_id"`
    Platform    string `json:"platform" db:"platform"`
    AppVersion  string `json:"app_version" db:"app_version"`
    UserAgent   string `json:"user_agent" db:"user_agent"`
    IsKnown     bool   `json:"is_known" db:"is_known"`
}

type TransactionType string
const (
    PURCHASE    TransactionType = "PURCHASE"
    TRANSFER    TransactionType = "TRANSFER"
    WITHDRAWAL  TransactionType = "WITHDRAWAL"
    DEPOSIT     TransactionType = "DEPOSIT"
)

type MerchantInfo struct {
    MerchantID   string `json:"merchant_id" db:"merchant_id"`
    MerchantName string `json:"merchant_name" db:"merchant_name"`
    Category     string `json:"category" db:"category"`
}
```

## 🚨 Modelos de Análise de Fraude

```go
type FraudAnalysis struct {
    TransactionID string             `json:"transaction_id" db:"transaction_id"`
    RiskScore     float64            `json:"risk_score" db:"risk_score"`
    RiskLevel     RiskLevel          `json:"risk_level" db:"risk_level"`
    Triggers      []DetectionTrigger `json:"triggers" db:"triggers"`
    Action        FraudAction        `json:"action" db:"action"`
    AnalyzedAt    time.Time          `json:"analyzed_at" db:"analyzed_at"`
}

type DetectionTrigger struct {
    RuleID      string  `json:"rule_id"`
    RuleName    string  `json:"rule_name"`
    Score       float64 `json:"score"`
    Confidence  float64 `json:"confidence"`
    Description string  `json:"description"`
}

type RiskLevel string
const (
    LOW      RiskLevel = "LOW"
    MEDIUM   RiskLevel = "MEDIUM"
    HIGH     RiskLevel = "HIGH"
    CRITICAL RiskLevel = "CRITICAL"
)

type FraudAction string
const (
    APPROVE   FraudAction = "APPROVE"
    REVIEW    FraudAction = "REVIEW"
    BLOCK     FraudAction = "BLOCK"
    ESCALATE  FraudAction = "ESCALATE"
)
```

## 🔔 Sistema de Alertas (Priority Queue)

```go
type Alert struct {
    ID          string         `json:"id"`
    Priority    int            `json:"priority"`    // Lower = Higher Priority
    RiskScore   float64        `json:"risk_score"`
    Transaction Transaction    `json:"transaction"`
    Analysis    FraudAnalysis  `json:"analysis"`
    CreatedAt   time.Time      `json:"created_at"`
    Status      AlertStatus    `json:"status"`
}

type AlertStatus string
const (
    PENDING   AlertStatus = "PENDING"
    REVIEWED  AlertStatus = "REVIEWED"
    RESOLVED  AlertStatus = "RESOLVED"
    IGNORED   AlertStatus = "IGNORED"
)

// Min-Heap Implementation
type AlertHeap []*Alert

func (h AlertHeap) Len() int           { return len(h) }
func (h AlertHeap) Less(i, j int) bool { return h[i].Priority < h[j].Priority }
func (h AlertHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *AlertHeap) Push(x interface{}) {
    *h = append(*h, x.(*Alert))
}

func (h *AlertHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}
```

## 📈 Estruturas de Cache e Histórico

```go
type UserProfile struct {
    UserID              string                 `json:"user_id" db:"user_id"`
    KnownDevices        []string               `json:"known_devices" db:"known_devices"`
    TransactionHistory  []TransactionSummary   `json:"transaction_history" db:"transaction_history"`
    AverageAmount       float64                `json:"average_amount" db:"average_amount"`
    LastTransactionAt   time.Time              `json:"last_transaction_at" db:"last_transaction_at"`
    RiskProfile         UserRiskProfile        `json:"risk_profile" db:"risk_profile"`
    CreatedAt           time.Time              `json:"created_at" db:"created_at"`
    UpdatedAt           time.Time              `json:"updated_at" db:"updated_at"`
}

type TransactionSummary struct {
    Amount      float64   `json:"amount"`
    Timestamp   time.Time `json:"timestamp"`
    Location    string    `json:"location"`
    DeviceID    string    `json:"device_id"`
}

type UserRiskProfile struct {
    RiskLevel       RiskLevel `json:"risk_level"`
    FraudCount      int       `json:"fraud_count"`
    LastFraudAt     time.Time `json:"last_fraud_at"`
    TrustScore      float64   `json:"trust_score"`
}

type DeviceProfile struct {
    DeviceID        string    `json:"device_id" db:"device_id"`
    UserID          string    `json:"user_id" db:"user_id"`
    Platform        string    `json:"platform" db:"platform"`
    FirstSeenAt     time.Time `json:"first_seen_at" db:"first_seen_at"`
    LastSeenAt      time.Time `json:"last_seen_at" db:"last_seen_at"`
    TransactionCount int      `json:"transaction_count" db:"transaction_count"`
    IsTrusted       bool      `json:"is_trusted" db:"is_trusted"`
}
```

## 🎯 Estruturas para Cenários Específicos

### Viagem Impossível
```go
type LocationAnalysis struct {
    PreviousLocation TransactionLocation `json:"previous_location"`
    CurrentLocation  TransactionLocation `json:"current_location"`
    Distance         float64             `json:"distance"`
    TimeDifference   time.Duration       `json:"time_difference"`
    MaxPossibleSpeed float64             `json:"max_possible_speed"`
    IsImpossible     bool                `json:"is_impossible"`
}
```

### Valor Anômalo
```go
type AmountAnalysis struct {
    UserID           string    `json:"user_id"`
    CurrentAmount    float64   `json:"current_amount"`
    AverageAmount    float64   `json:"average_amount"`
    StandardDev      float64   `json:"standard_deviation"`
    Threshold        float64   `json:"threshold"`
    IsAnomalous      bool      `json:"is_anomalous"`
    AnomalyScore     float64   `json:"anomaly_score"`
}
```

### Velocidade de Transações
```go
type VelocityAnalysis struct {
    UserID           string        `json:"user_id"`
    WindowDuration   time.Duration `json:"window_duration"`
    TransactionCount int           `json:"transaction_count"`
    MaxAllowed       int           `json:"max_allowed"`
    IsExceeded       bool          `json:"is_exceeded"`
    VelocityScore    float64       `json:"velocity_score"`
}
```

## 🔧 Estruturas de Configuração

```go
type RuleConfig struct {
    RuleID          string  `json:"rule_id"`
    RuleName        string  `json:"rule_name"`
    IsEnabled       bool    `json:"is_enabled"`
    Weight          float64 `json:"weight"`
    Threshold       float64 `json:"threshold"`
    Description     string  `json:"description"`
    CreatedAt       time.Time `json:"created_at"`
    UpdatedAt       time.Time `json:"updated_at"`
}

type SystemConfig struct {
    MaxRiskScore        float64       `json:"max_risk_score"`
    AnalysisTimeout     time.Duration `json:"analysis_timeout"`
    CacheTTL            time.Duration `json:"cache_ttl"`
    AlertRetentionDays  int           `json:"alert_retention_days"`
    MaxConcurrentWorkers int          `json:"max_concurrent_workers"`
}
```

## 📊 Estruturas de Métricas e Estatísticas

```go
type SystemStats struct {
    TotalTransactions   int64     `json:"total_transactions"`
    FraudDetected       int64     `json:"fraud_detected"`
    FalsePositives      int64     `json:"false_positives"`
    AverageResponseTime float64   `json:"average_response_time"`
    ThroughputPerSecond float64   `json:"throughput_per_second"`
    LastUpdated         time.Time `json:"last_updated"`
}

type RuleStats struct {
    RuleID              string  `json:"rule_id"`
    TriggerCount        int64   `json:"trigger_count"`
    TruePositives       int64   `json:"true_positives"`
    FalsePositives      int64   `json:"false_positives"`
    Accuracy            float64 `json:"accuracy"`
    LastTriggeredAt     time.Time `json:"last_triggered_at"`
}
```

## 🚀 Estruturas para APIs

### Request/Response Models
```go
type AnalyzeRequest struct {
    Transaction Transaction `json:"transaction"`
}

type AnalyzeResponse struct {
    TransactionID string         `json:"transaction_id"`
    RiskScore     float64        `json:"risk_score"`
    RiskLevel     RiskLevel      `json:"risk_level"`
    Action        FraudAction    `json:"action"`
    Triggers      []DetectionTrigger `json:"triggers"`
    ProcessedAt   time.Time      `json:"processed_at"`
}

type AlertsResponse struct {
    Alerts []Alert `json:"alerts"`
    Total  int     `json:"total"`
    Page   int     `json:"page"`
    Limit  int     `json:"limit"`
}

type StatsResponse struct {
    System SystemStats `json:"system"`
    Rules  []RuleStats `json:"rules"`
}
```
---

**💡 Dica:** Comece implementando as estruturas essenciais e vá adicionando as avançadas conforme o tempo permitir durante o hackathon!
