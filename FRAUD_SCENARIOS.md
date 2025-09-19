# 🚨 Cenários de Detecção de Fraude

Este documento detalha todos os cenários de detecção de fraude implementados no sistema antifraude, organizados por complexidade e tempo de implementação.

## 📊 Diagramas de Detecção por Cenário

### 🔴 Cenário 1: Viagem Impossível
```mermaid
graph TD
    A[Transação com Localização] --> B[Buscar Última Transação do Usuário]
    B --> C{Última Transação Existe?}
    C -->|Não| D[Score: 0 - Primeira Transação]
    C -->|Sim| E[Calcular Distância]
    E --> F[Calcular Tempo Decorrido]
    F --> G[Calcular Velocidade Máxima Possível]
    G --> H{Distância > Velocidade Máxima?}
    H -->|Não| I[Score: 0 - Viagem Possível]
    H -->|Sim| J[Score: 80 - Viagem Impossível]
    D --> K[Retornar Resultado]
    I --> K
    J --> K
    
    style A fill:#e1f5fe
    style J fill:#ffebee
    style K fill:#e8f5e8
```

### 📊 Cenário 2: Valor Anômalo
```mermaid
graph TD
    A[Transação com Valor] --> B[Buscar Histórico do Usuário]
    B --> C{Histórico >= 5 transações?}
    C -->|Não| D[Score: 0 - Histórico Insuficiente]
    C -->|Sim| E[Calcular Média dos Valores]
    E --> F[Calcular Desvio Padrão]
    F --> G[Definir Threshold: Média + 3σ]
    G --> H{Valor > Threshold?}
    H -->|Não| I[Score: 0 - Valor Normal]
    H -->|Sim| J[Score: 70 - Valor Anômalo]
    D --> K[Retornar Resultado]
    I --> K
    J --> K
    
    style A fill:#e1f5fe
    style J fill:#ffebee
    style K fill:#e8f5e8
```

### 📱 Cenário 3: Dispositivo Desconhecido
```mermaid
graph TD
    A[Transação com IP e GPS] --> B[Geolocalizar IP Address]
    B --> C[Obter Coordenadas GPS da Transação]
    C --> D[Calcular Distância entre IP e GPS]
    D --> E{Distância > 50km?}
    E -->|Não| F[Score: 0 - Localização Consistente]
    E -->|Sim| G{Distância > 200km?}
    G -->|Não| H[Score: 30 - Inconsistência Moderada]
    G -->|Sim| I[Score: 60 - Inconsistência Crítica]
    F --> J[Retornar Resultado]
    H --> J
    I --> J
    
    style A fill:#e1f5fe
    style I fill:#ffebee
    style J fill:#e8f5e8
```


### ⚡ Cenário 4: Velocidade de Transações
```mermaid
graph TD
    A[Transação Recebida] --> B[Definir Janela de Tempo: 5 minutos]
    B --> C[Contar Transações do Usuário na Janela]
    C --> D{Contagem >= 10?}
    D -->|Não| E[Score: 0 - Velocidade Normal]
    D -->|Sim| F{Contagem >= 20?}
    F -->|Não| G[Score: 25 - Velocidade Alta]
    F -->|Sim| H[Score: 50 - Velocidade Crítica]
    E --> I[Retornar Resultado]
    G --> I
    H --> I
    
    style A fill:#e1f5fe
    style H fill:#ffebee
    style I fill:#e8f5e8
```

### 🕐 Cenário 5: Horário Suspeito
```mermaid
graph TD
    A[Transação com Timestamp] --> B[Extrair Hora da Transação]
    B --> C{Hora entre 00h-06h?}
    C -->|Não| D[Score: 0 - Horário Normal]
    C -->|Sim| E{Hora entre 02h-04h?}
    E -->|Não| F[Score: 20 - Horário Suspeito]
    E -->|Sim| G[Score: 30 - Horário Muito Suspeito]
    D --> H[Retornar Resultado]
    F --> H
    G --> H
    
    style A fill:#e1f5fe
    style G fill:#ffebee
    style H fill:#e8f5e8
```

### 🔢 Cenário 6: Sequência de Valores
```mermaid
graph TD
    A[Transação com Valor] --> B[Buscar Últimas 5 Transações do Usuário]
    B --> C{Histórico >= 3 transações?}
    C -->|Não| D[Score: 0 - Histórico Insuficiente]
    C -->|Sim| E[Verificar Sequência Aritmética]
    E --> F{É Sequência Aritmética?}
    F -->|Não| G[Score: 0 - Sem Padrão]
    F -->|Sim| H{Diferença >= 100?}
    H -->|Não| I[Score: 20 - Sequência Pequena]
    H -->|Sim| J[Score: 40 - Sequência Suspeita]
    D --> K[Retornar Resultado]
    G --> K
    I --> K
    J --> K
    
    style A fill:#e1f5fe
    style J fill:#ffebee
    style K fill:#e8f5e8
```

### 🌍 Cenário 7: Localização Inconsistente
```mermaid
graph TD
    A[Transação com IP e GPS] --> B[Geolocalizar IP Address]
    B --> C[Obter Coordenadas GPS da Transação]
    C --> D[Calcular Distância entre IP e GPS]
    D --> E{Distância > 50km?}
    E -->|Não| F[Score: 0 - Localização Consistente]
    E -->|Sim| G{Distância > 200km?}
    G -->|Não| H[Score: 30 - Inconsistência Moderada]
    G -->|Sim| I[Score: 60 - Inconsistência Crítica]
    F --> J[Retornar Resultado]
    H --> J
    I --> J
    
    style A fill:#e1f5fe
    style I fill:#ffebee
    style J fill:#e8f5e8
```

### 💰 Cenário 8: Valor Redondo
```mermaid
graph TD
    A[Transação com Valor] --> B[Verificar se Valor é Inteiro]
    B --> C{Valor é Inteiro?}
    C -->|Não| D[Score: 0 - Valor com Centavos]
    C -->|Sim| E{Valor >= 1000?}
    E -->|Não| F[Score: 0 - Valor Pequeno]
    E -->|Sim| G{Valor é Múltiplo de 1000?}
    G -->|Não| H[Score: 15 - Valor Redondo]
    G -->|Sim| I[Score: 25 - Valor Muito Redondo]
    D --> J[Retornar Resultado]
    F --> J
    H --> J
    I --> J
    
    style A fill:#e1f5fe
    style I fill:#fff3e0
    style J fill:#e8f5e8
```

### 👤 Cenário 9: Usuário Inativo
```mermaid
graph TD
    A[Transação de Usuário] --> B[Buscar Última Transação do Usuário]
    B --> C{Última Transação Existe?}
    C -->|Não| D[Score: 0 - Primeira Transação]
    C -->|Sim| E[Calcular Tempo desde Última Transação]
    E --> F{Tempo > 90 dias?}
    F -->|Não| G[Score: 0 - Usuário Ativo]
    F -->|Sim| H{Tempo > 180 dias?}
    H -->|Não| I[Score: 20 - Usuário Inativo]
    H -->|Sim| J[Score: 40 - Usuário Muito Inativo]
    D --> K[Retornar Resultado]
    G --> K
    I --> K
    J --> K
    
    style A fill:#e1f5fe
    style J fill:#ffebee
    style K fill:#e8f5e8
```

### 🔄 Cenário 10: Transações Consecutivas
```mermaid
graph TD
    A[Transação com Valor] --> B[Buscar Últimas 3 Transações do Usuário]
    B --> C{Histórico >= 3 transações?}
    C -->|Não| D[Score: 0 - Histórico Insuficiente]
    C -->|Sim| E[Verificar se Valores são Iguais]
    E --> F{Últimas 3 transações têm mesmo valor?}
    F -->|Não| G[Score: 0 - Valores Diferentes]
    F -->|Sim| H{Valor >= 1000?}
    H -->|Não| I[Score: 15 - Consecutivas Pequenas]
    H -->|Sim| J[Score: 35 - Consecutivas Grandes]
    D --> K[Retornar Resultado]
    G --> K
    I --> K
    J --> K
    
    style A fill:#e1f5fe
    style J fill:#ffebee
    style K fill:#e8f5e8
```

## 📊 Sistema de Pontuação

```go
type RiskScore struct {
    BaseScore    float64 `json:"base_score"`
    Triggers     []Trigger `json:"triggers"`
    FinalScore   float64 `json:"final_score"`
    RiskLevel    RiskLevel `json:"risk_level"`
}

type Trigger struct {
    RuleID       string  `json:"rule_id"`
    RuleName     string  `json:"rule_name"`
    Score        float64 `json:"score"`
    Description  string  `json:"description"`
}
```

---

**💡 Dica:** Comece implementando os cenários mais simples e vá adicionando complexidade gradualmente. Isso garante que você tenha um sistema funcional desde o início do hackathon!
