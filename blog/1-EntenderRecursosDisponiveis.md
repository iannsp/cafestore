# Entender os recursos Disponiveis.

Baseado no objetivo de ter um serviço de busca para os produtos e serviços da Café Store eu vou começar por:

## 1. Entender o catálogo

Nesse ponto quero entender que tipo de dado esta disponivel no site da loja, os meios de acesso a eles e quais deles podem ser utilizados por um software.

O primeiro ponto de acesso aos dados do produto é pela **página inicial do site** que dá acesso a outras páginas cada uma com sua url. Essas páginas todas juntas são a experiencia da loja na internet.

Ao acessar a página podemos ver os produtos mais vendidos, navegar pelas categorias, fazer buscas e também aprender sobre outros serviços como o programa de pontos.


Quando a gente olha o código da página mais de perto percebe que o html foi escrito para ajudar outros programas a consumir o conteudo do site com o uso de dados estruturados e microdata com tags [meta](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/meta), [schema.org/Store](http://schema.org/Store) e [schema.org/Product](http://schema.org/Product).

Desses detalhes da pra dizer que, ao menos em uma escala de uso pessoal, dá pra criar um crawler que navegue pelo conteudo sem bater num cloudfare da vida e que alguns use cases interessantes podem ser implementados.

- Saber das novidades.
  Se a pagina principal é mantida atualizada então visitar essa página e extrair dados do seu conteúdo pode trazer as coisas em destaque.

- Buscar por produtos usando nome, descrição ou tipo utilizando o link de busca

- Comparar preços.
  Fazer Pesquisa ativa e comparar com um preço registrado como resultado de uma busca mais antiga.
 
- monitorar mudanças de preço no site.
  Agendar pesquisas de produtos que eu tenha interesse para saber de mudanças de preço.

  - Consultar promoções.
  Tem um link na página principal que leva direto para uma página de promoções.
  
A partir daqui vamos coletar e processar dados.

[próximo post ::  Extraindo Dados](2-extraindodados.md)

    

