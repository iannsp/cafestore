package naivecatalog

import (
	"encoding/json"
	"os"
)

type NaiveCategory struct {
	Name       string          `json:"name"`
	Categories []NaiveCategory `json:"children, omitempty"`
}

type NaiveCategoryTree struct {
	Categories []NaiveCategory `json:"categories_tree"`
}

type NaiveCategories struct {
	tree NaiveCategoryTree
}

func (c *NaiveCategories) Load(filePath string) error {
	var tree NaiveCategoryTree
	catFContent, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(catFContent, &tree)
	if err == nil {
		c.tree = tree
	}
	return nil
}

func (c *NaiveCategories) Navigate(path []string) NaiveCategoryNavigation {

	nav := c.tree.Categories

	if len(path) > 0 {
		for _, step := range path {
			found := false
			for _, cat := range nav {
				if step == cat.Name {
					found = true
					nav = cat.Categories
				}
			}
			if !found {
				nav = []NaiveCategory{}
			}
		}
	}
	return NaiveCategoryNavigation{Path: path, Categories: nav}
}

