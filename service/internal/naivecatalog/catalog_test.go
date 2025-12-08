package naivecatalog

import (
	"testing"
    "os"
)


func TestNewNaiveCatalog(t *testing.T) {
	c := NaiveCatalog{}
	err := c.Load(os.Getenv("NAIVE_CATALOG_PATH"))

    if err != nil{
        t.Errorf("Error loading Catalog. Err: %s", err.Error())
    }

    if c.Len() != 571 {
        t.Errorf("Fail Load Catalog. Expected 571 Items with 'Café', got %d", c.Len())
    }
}

func TestSearchBy(t *testing.T) {
	c := NaiveCatalog{}
	c.Load(os.Getenv("NAIVE_CATALOG_PATH"))

	search := c.SearchBy("café")
    if search.Len() != 457 {
        t.Errorf("Fail search for 'Café'. Expected 25 Items with 'Café', got %d", search.Len())
    }

	refined := c.SearchBy("adocicado")
    if refined.Len() != 25 {
        t.Errorf("Fail search for Café Adocicado. Expected 25 cafés, got %d", refined.Len())
    }
}

func TestSearchCatalogValidIfReturnCafe(t *testing.T){
	c := NaiveCatalog{}
	c.Load("./data/catalog/")
	c.Load(os.Getenv("NAIVE_CATALOG_PATH"))

	cat := NaiveCategories{}
	c.Load(os.Getenv("NAIVE_CATEGORY_PATH"))
	nav := cat.Navigate([]string{"Cafés","Intensidade","Café Suave"})
	search := c.SearchBy(nav.ToString())
    if search.Len() != 7 {
        t.Errorf("Fail search for %s. Expected 7 cafés, got %d", nav.ToString(), search.Len())
    }
}
