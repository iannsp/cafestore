package naivecatalog

import (
	"os"
	"strings"
)

type NaiveCatalog struct {
	Items []Item
}

func (c *NaiveCatalog) Len() int {
	return len(c.Items)
}

func (c *NaiveCatalog) Load(dirPath string) error {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if !entry.IsDir() { // Process only files, skip subdirectories
			filePath := dirPath + "/" + entry.Name()
			content, err := os.ReadFile(filePath) // Use os.ReadFile for reading file content
			if err != nil {
				return err
			}
			// has only '[]' what means the product with this Id don't exist.
			trimmedJson := strings.TrimSpace(string(content))
			if len(trimmedJson) <= 2 {
				continue
			}
			removedfromArray := trimmedJson[1 : len(trimmedJson)-1]
			item, err := NewItem(removedfromArray)
			if err == nil {
				c.Items = append(c.Items, item)
			}
		}
	}
	return nil
}

func (c *NaiveCatalog) SearchBy(by string) NaiveCatalog {
	found := []Item{}
	for _, item := range c.Items {
		if strings.Contains(strings.ToLower(item.Json), strings.ToLower(by)) {
			item, err := NewItem(item.Json)
			if err == nil {
				found = append(found, item)
			}
		}
	}
	return NaiveCatalog{Items: found}
}
