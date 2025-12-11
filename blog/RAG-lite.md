# Rag Lite

* Esta versão de prompt irá trabalhar somente com cafés e ignorar outros produtos do catálogo.


## Como executar

```bash
    # para gerar o cli execute o script de build:
    ./scripts/build.sh raglite

    # export sua chave de API de GEMINI.
    export GEMINI_API_KEY=""

    # export o path para seu diretorio service/ui (service/ui)
    export UI_PATH='/projetos/cafestore/service/ui/'

    # export o path para o diretório de dados sobre cafés.(data/cafes)
    export DATA_PATH="/projetos/cafestore/data/cafes/"

    # execute o app e acesse através de http://127.0.0.1:8080:
    ./bin/raglite
    
```
Essa implementação se beneficia da implementação de NaiveCatalog. Com mais experiencia no modelo de dados disponivel foi possivel:

* Estabelecer um json minimo para reduzir o custo de tokens numa abordagem inicial de injeção do json.
* selecionar um grupo de produtos(café) para realizar um experimento menor.
* Desenvolver uma jornada com o mínimo de use cases afim de reduzir a complexidade.

---
## Use Cases

Objetivo:  Guiar o cliente interessado em café na sua jornada de atendimento.

1. Descobrir Produtos/Recomendação: (exemplo: "To procurando um café suave, voce recomenda algum torra clara?")  
Ação: Utilizar as caracteristicas e descrição do café para elaborar sugestões.

2. Informação Sobre produto: (exemplo: "O que você pode me dizer sobre o café Café Baggio Bourbon. Qual o preço dele?)  
Ação: Descrever o produto, falar seu preço, e mostrar a sugerir a compra com a URL de compra.

3. Navegar pelas Categorias: (exemplo: "Que marcas de café voce vende?", "Você tem algum café premiado?")  
Ação: Utilizar as informações em categorias.json para buscar opções.

4. Compra: (e.g., "Como fa,o pra comprar?")  
Ação: Mostrar o link do produto e dar informaçõesde preço e disponibilidade.

5. Perguntas sem contexto: (e.g., "Onde esta minha compra?", "Posso falar com um humano?", "Qual o horário da sua loja?")  
Ação: Ter plano de tratamento definidos no prompt.

---
## Catalogo

Utilizando o playground do gemini criei duas versões reduzidas do catalogo de café. A primeira focada em maximizar a economia de tokens(1) e a segunda em maximizar quantidade de dados X custo de tokens(2).


1. Menor Custo de tokens [json](../data/cafes/catalog_cafe_small.json) | [toon](../data/cafes/catalog_cafe_small.toon)
2. Balancear custo de token com disponibilidade de Informação [json](../data/cafes/catalogo_cafe_details.json) | [toon](../data/cafes/catalogo_cafe_details.toon)

#### Análise do catalogo de menor custo de tokens(1)

*   **Caracteres totais:** ~38.000 caracteres.
*   **Quantidade de Tokens (Estimativa padrão OpenAI):** **~11.200 tokens**.

> **Nota:** Links (URLs) e formatação JSON (`{`, `}`, `"`) consomem mais tokens do que texto corrido simples.

#### Custo de Input(Json)
*Cotação usada: US$ 1,00 = R$ 6,00 (aproximado).*

| Modelo | Preço por 1M tokens (Input) | Custo em Dólar (USD) | Custo em Reais (BRL) |
| :--- | :--- | :--- | :--- |
| **GPT-4o-mini** (Mais econômico) | US$ 0.15 | **$ 0.0017** | **R$ 0,01** |
| **GPT-4o** (Alta inteligência) | US$ 2.50 | **$ 0.0280** | **R$ 0,17** |
| **Claude 3.5 Sonnet** (Anthropic) | US$ 3.00 | **$ 0.0336** | **R$ 0,20** |
| **Claude 3 Haiku** (Rápido) | US$ 0.25 | **$ 0.0028** | **R$ 0,02** |
| **Gemini 1.5 Flash** (Google) | US$ 0.075 | **$ 0.0008** | **< R$ 0,01** |

O mesmo catalogo formatado usando TOON ([Token-Oriented Object Notation](https://github.com/toon-format/spec))

Redução de quantidade de tokens de **~31%**.

*   **Original Tokens:** ~11,200
*   **New TOON Tokens:** **~7,750**

#### Custo de Input (TOON)

| Model | Cost (USD) | Cost (BRL) |
| :--- | :--- | :--- |
| **GPT-4o** | $0.0194 | R$ 0,12 |
| **GPT-4o-mini** | $0.0011 | R$ 0,006 |
| **Claude 3.5 Sonnet** | $0.0232 | R$ 0,14 | 

### Análise do catalogo que balanceia nível de informacao com custo(2)

*   **Caracteres totais:** ~48,000 caracteres.
*   **Quantidade de Tokens (Estimativa padrão OpenAI):** **~12,500 tokens**.

### 2. Custo de Input(Json)

| Model | Price per 1M Input Tokens | Cost for this JSON (USD) | Cost for this JSON (BRL)* |
| :--- | :--- | :--- | :--- |
| **GPT-4o** | $2.50 | **$0.03** (3 cents) | **R$ 0,18** |
| **GPT-4o-mini** | $0.15 | **$0.0019** (0.2 cents) | **R$ 0,01** |
| **GPT-4 Turbo** | $10.00 | **$0.12** (12 cents) | **R$ 0,72** |
| **GPT-3.5 Turbo**| $0.50 | **$0.006** (0.6 cents) | **R$ 0,04** |

### 2. Custo de Input(TOON)

Redução de quantidade de tokens de **~21.7%**.

*   **Original Tokens:** ~12,500
*   **New TOON Tokens:** **~9,795**---

| Model | Price per 1M Input Tokens | Cost (USD) | Cost (BRL)* |
| :--- | :--- | :--- | :--- |
| **GPT-4o-mini** | $0.15 | **$0.0015** (0.15 cents) | **R$ 0,01** |
| **GPT-4o** | $2.50 | **$0.024** (2.4 cents) | **R$ 0,15** |
Explicação:


# Prompt

```markdown
# PERSONA
Você é a **"Barista Virtual da Café Store"**.
Seu tom de voz é de **Vizinho Amigo**: entusiasta, caloroso, mas **direto e eficiente**.
Você ama café, mas valoriza o tempo do cliente. Fale sempre em português (Brasil).

# OBJETIVO
Ajudar clientes a navegar pelo catálogo, tirar dúvidas e fechar compras.

# INFORMAÇÕES QUE VOCÊ POSSUI (CONTEXTO)
1.  **Catálogo de Produtos:** 'productName', 'description', 'Intensidade', 'Sabor', 'Aroma', 'Torra', 'Acidez', 'Corpo', 'Variedade', 'Origem', 'link', 'price' e 'image_url' (se disponível).
2.  **Categorias:** Estrutura de categoria de produtos.

# REGRAS DE COMUNICAÇÃO

1.  **CONCISÃO É CHAVE:**
    *   Evite introduções longas ou "conversa fiada" excessiva.
    *   Vá direto à resposta ou recomendação.
    *   Pense no que vai responder e então reduza o texto.

2.  **RECOMENDAÇÕES (LIMITE DE 3):**
    *   Sugira no **máximo 3 produtos** por vez. Se houver mais opções relevantes, pergunte se o cliente quer ver mais.
    *   Ao recomendar, destaque **apenas** o atributo sensorial mais relevante para o pedido do cliente (ex: "Notas de chocolate" para quem pediu intensidade).

3.  **FORMATO DE EXIBIÇÃO DO PRODUTO:**
    Para cada produto sugerido, use estritamente este layout visual:

    **{productName}**
    {Breve motivo da escolha baseado no gosto do cliente}
    💰 **{price}** | [Comprar Agora]({link})

4.  **FILTRAGEM:**
    *   Se o pedido for vago (ex: "Quero café"), faça **uma** pergunta rápida de qualificação (ex: "Grão ou cápsula? Suave ou intenso?") antes de listar produtos.

# LIMITAÇÕES
*   Você **NÃO** acessa dados de pedidos, entregas ou contas.
*   **Fallback:** Para assuntos fora do catálogo, direcione para o atendimento humano/fale conosco de forma breve.

# CONTEXTO ATUAL (PRODUTOS/CATEGORIAS)
**INSTRUÇÃO:** Baseie-se apenas nos dados abaixo.
```

# Codígo

O Código da aplicação RAG Lite foi escrito em GO.

A intenção dessa implementação não é suportar vários usuários; Essa feature será implementada em RAG1.

O que você vai ver no Código.

1. Em [/service/internal/raglite/httserver.go](../service/internal/raglite/httserver.go)

A separação do [chat](../service/internal/raglite/chat.go) e dos handlers de [Https](../service/internal/raglite/httserver.go).

Ao invés de utilizar http server direto eu gosto de organizar o server de maneira que eu seja obrigado a explicitamente "pendurar" as rotas e os handlers nele através de AttachRoutes(route string, handler HttpServerHandlerFunc):

```go
type HttpServerHandlerFunc func(http.ResponseWriter, *http.Request)

func (hs *HttpServer) AttachRoutes(route string, handler HttpServerHandlerFunc){
    hs.handler.HandleFunc(route, handler)
}

// em cmd/raglite/main.go

hs := raglite.NewHttpServer("8080")
hs.AttachRoutes("/", index)
hs.AttachRoutes("/api/chat", handleMessage)

```
2. Em [/service/internal/raglite/chat.go](../service/internal/raglite/chat.go)

Utilizei genai.ChatSession para manter o contexto da conversa.
  

```go
type Chat struct{
    geminiApiKey string
    prompt string
    session *genai.ChatSession
}

```

E em [/service/cmd/raglite/main.go](../service/cmd/raglite/main.go), voce pode ver o setup da aplicação.

```go
hs := raglite.NewHttpServer("8080")

chat = raglite.NewChat(geminiApiKey)
chat.Prompt(loadPrompt(datapath))
chat.AttachRoutes(&hs)

hs.AttachRoutes("/", index)
hs.AttachRoutes("/api/chat", handleMessage)

err := hs.ListenAndServe()

``` 

