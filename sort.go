package pagination

import (
	"fmt"
	"strings"
)

type (
	Sort struct {
		// The sort value.
		Sort string `query:"sort"`

		// The order value.
		Order string `query:"order"`

		// List of allowed sort values.
		// Can be []string{} or map[string]string.
		// When it is a map, the key is the sort value, value is the
		// sort expression in the ORDER BY clause. If value is empty,
		// key is used. You can use {} in value to represent the order
		// value (sort direction). For example:
		// map[string]string{
		// 	"created_at": "",
		// 	"posts_count":  "posts_count {} NULLS LAST",
		// }
		AllowedSorts interface{}

		// DefaultSort is used if the sort value is empty or not in the
		// AllowedSorts.
		DefaultSort string

		// DefaultOrder is used if the order value is empty.
		// Must be "asc" or "desc". Default value is "desc".
		DefaultOrder string

		// Make "DefaultSort" as the "fallback" sort method.
		AlwaysUseDefaultSort bool
	}

	SortResult struct {
		// Effective sort value, null (nil) means no sort
		Sort *string

		// Effective order value, null (nil) means no sort
		Order *string
	}
)

func (s Sort) IsAllowed(sort string) bool {
	switch sorts := s.AllowedSorts.(type) {
	case []string:
		for i := 0; i < len(sorts); i++ {
			if sorts[i] == sort {
				return true
			}
		}
		return false
	case map[string]string:
		if _, ok := sorts[sort]; ok {
			return true
		}
		return false
	default:
		return false
	}
}

func (s Sort) GetSort() string {
	if s.Sort != "" && s.IsAllowed(s.Sort) {
		return s.Sort
	}
	if s.DefaultSort != "" {
		return s.DefaultSort
	}
	return ""
}

func (s Sort) GetOrder() string {
	order := s.Order
	if order == "" {
		order = s.DefaultOrder
	}
	if order == "asc" {
		return "ASC"
	} else {
		return "DESC"
	}
}

func (s Sort) OrderByValue() string {
	if sort := s.GetSort(); sort != "" {
		out := s.getSortExpression(sort, s.GetOrder())
		if s.AlwaysUseDefaultSort && sort != s.DefaultSort {
			defaultOrder := s.DefaultOrder
			if defaultOrder == "asc" {
				defaultOrder = "ASC"
			} else {
				defaultOrder = "DESC"
			}
			defaultSort := s.getSortExpression(s.DefaultSort, defaultOrder)
			return out + ", " + defaultSort
		}
		return out
	}
	return ""
}

func (s Sort) OrderBy() string {
	if value := s.OrderByValue(); value != "" {
		return "ORDER BY " + value
	}
	return ""
}

func (s Sort) SortResult() (r SortResult) {
	if sort := s.GetSort(); sort != "" {
		order := strings.ToLower(s.GetOrder())
		r.Sort = &sort
		r.Order = &order
	}
	return
}

func (s Sort) getSortExpression(in, order string) string {
	var sort string
	switch sorts := s.AllowedSorts.(type) {
	case []string:
		sort = in
	case map[string]string:
		if out, ok := sorts[in]; ok && out != "" {
			sort = out
		} else {
			sort = in
		}
	default:
		sort = in
	}
	if strings.Contains(sort, "{}") {
		return strings.Replace(sort, "{}", order, -1)
	} else {
		return fmt.Sprintf("%s %s", sort, order)
	}
}
