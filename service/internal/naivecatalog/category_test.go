package naivecatalog

import (
	"testing"
    "os"
)

func TestNewCategories(t *testing.T) {
	c := NaiveCategories{}
	err := c.Load( os.Getenv("NAIVE_CATEGORY_PATH"))
    if err != nil{
        t.Errorf("Fail to load Category json. Error: %s", err.Error())
    }
}

func TestInvalidTraverseAfterOne(t *testing.T) {
	c := NaiveCategories{}
	err := c.Load(os.Getenv("NAIVE_CATEGORY_PATH"))
    if err != nil{
        t.Errorf("Fail to load Category json. Error: %s", err.Error())
    }

	nav := c.Navigate([]string{})

    //slog.Info("Pesquisa por Categoria Cafés")
	nav = c.Navigate([]string{"Cafés"})
	nav.Path = append(nav.Path, "teste")

	//slog.Info("Pesquisa por Categoria Cafés/Teste", "path", nav.Path)
	nav = c.Navigate(nav.Path)
    if len(nav.Categories) > 0 {
        t.Errorf("Error Load Test Category. Expect 0 Categories, got %d", len(nav.Categories))
    }
}

func TestInvalidPath(t *testing.T) {
	c := NaiveCategories{}
	c.Load( os.Getenv("NAIVE_CATEGORY_PATH"))
	nav := c.Navigate([]string{"Categoria Invalida"})
	if len(nav.Categories) > 0 {
		t.Errorf("Categoria Pesquisada Invalida. Esperado 0 resultado recebido %d", len(nav.Categories))
	}
}

func TestValidTraverseToTheEndAndGetStringPath(t *testing.T) {
	c := NaiveCategories{}
	c.Load( os.Getenv("NAIVE_CATEGORY_PATH"))

	nav := c.Navigate([]string{"Cafés"})
	if len(nav.Categories) != 3 { // Intensidade, metodo, tipo
		t.Errorf("Erro de Pesquisa. Esperado 8 resultado recebido %d", len(nav.Categories))
	}

	nav.Path = append(nav.Path, "Intensidade")
	nav = c.Navigate(nav.Path)

	if len(nav.Categories) != 3 {
		t.Errorf("Erro de Pesquisa. Esperado 0 resultado recebido %d", len(nav.Categories))
	}

    fullpath := nav.ToString()

    if fullpath != "/Cafés/Intensidade/"{
		t.Errorf("Erro no path de navegacao. Esperado '/Cafés/Intensidade/', recebido %s", fullpath)
    }
}
