# Naive Catalog

Ingenua por que não é ideal para produção. Ela não foi escrita olhando para requisitos 
de produção como performance.

## Como executar

```bash
    # para gerar o cli execute o script de build:
    ./scripts/build.sh naiveCatalog


    # para executar o terminal interativo:
    # ./bin/naivecatalog-cli da acesso a um terminal iterativo.

```

Esta implementação foi criada para entender melhor os arquivos de dados.  
Enquanto Era realizada, ela ajudava a amadurecer o conhecimento sobre os dados  
do catalogo, como é o caso da extração de categorias que foi feita.


O catalogo utilizado aqui foi gerado como descrito no [crawler Json](2.2-extraindodados_crawlerJSON.md)
e o resultado final é um diretorio com um json para cada produto do catálogo na 
data de hoje (8/12/2025).

As perguntas contidas nessa implementação são: 

1. Como o catalogo é pequeno eu posso lidar com ele completamente em memoria?
2. Como seria uma alternativa de pesquisa sem dependencia externa?

Explicação:

### Carregando o catalogo

O Carregamento do catalogo é feito a partir dos arquivos json em data/catalog e 
para trabalhar com um catalogo atualizado é necessario executar as instrucoes em 
[crawler Json](blog/2.2-extraindodados_crawlerJSON.md) e atualizar a variavel de 
ambiente NAIVE_CATALOG_PATH.

```GO
    // carregando o catalogo.
	c := NaiveCatalog{}
	err := c.Load(os.Getenv("NAIVE_CATALOG_PATH"))
```

Uma vez que o Catalog foi carregado somente parte dos dados são extraidas do Json e 
definidas como dados na struct Item. Cada item é um Produto/Serviço no Catalogo: 

```GO
type Item struct {
	Id               string `json:"productId"`
	Name             string `json:"productName"`
	Brand            string `json:"Brand"`
	ShortDescription string `json:"metaTagDescription"`
	Link             string `json:"link"`
	ImageUrl         string
	Json             string
}
```

### Realizando buscas

Toda busca é realizada via um string.Contains(ToUpper([search_string])), que atravessa
todo o catalogo buscando pela string no Item.Json e os itens que atendem a pesquisa 
são retornados como um novo Catalogo, um catalogo reduzido da busca.

O catalogo resultado da busca funciona como o Catalogo original e novas buscas nele 
acabam refinando a busca cada vez mais.

```GO
	c := NaiveCatalog{}
	c.Load(os.Getenv("NAIVE_CATALOG_PATH"))

    // retornar todos os produtos que tem a palavra café no Item.Json
	search := c.SearchBy("café")

    // retonar todos os produtos que contém a palavra café e a palavra adocicado.
	refined := c.SearchBy("adocicado")

    // para iterar pelo resultado 
    for _, item := range refined.Items{
        fmt.Printf(`Produto Id: %s \n 
           Nome: %s \n 
           Marca: %s \n 
           Descrição: %s \n 
           Link: %s \n 
        `)
    }
```

Outra maneira de buscar é utilizando a arvore de categorias. A pesquisa é realizada 
da mesma maneira utilizando **SearchBy** porém a string buscada permite refinar por 
 categoria devido a maneira como as categorias estão anotadas nos arquivos json.

```GO
	catalog := NaiveCatalog{}
	catalog.Load(os.Getenv("NAIVE_CATALOG_PATH"))

   
	cat := NaiveCategories{}
	cat.Load( os.Getenv("NAIVE_CATEGORY_PATH"))

    // todas as categorias informadas de uma vez...
	nav := c.Navigate([]string{"Cafés", "Intensidade", "Café Intenso"})

    // ou em uma pesquisa incremental em que o usuario navega pelas possibilidades.
	nav := c.Navigate([]string{"Cafés"})
	nav.Path = append(nav.Path, "Intensidade")
	nav = c.Navigate(nav.Path)
	nav.Path = append(nav.Path, "Café Intenso")
	nav = c.Navigate(nav.Path)

    // retornar todos os produtos que atendem a arvore de categoria: 
    // com path: /Cafés/Intensidade/Café Intenso 
	search := c.SearchBy(nav.ToString())

    // para iterar pelo resultado 
    for _, item := range search.Items{
        fmt.Printf(`Produto Id: %s \n 
           Nome: %s \n 
           Marca: %s \n 
           Descrição: %s \n 
           Link: %s \n 
        `)
    }
```
