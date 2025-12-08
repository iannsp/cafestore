package naivecatalog

import (
	"strings"
)


type NaiveCategoryNavigation struct {
	Path       []string
	Categories []NaiveCategory
}

func (cn NaiveCategoryNavigation) ToString() string {
    if len(cn.Path) > 0 {
	    return "/" + strings.Join(cn.Path, "/") + "/"
    }
    return ""
}
