package pagination

import (
	"regexp"
	"strings"
)

type (
	Query struct {
		Query string `query:"query"`
	}

	QueryResult struct {
		// Effective sort value, null (nil) means no sort
		Query *string
	}
)

var (
	regexpSpace = regexp.MustCompile(`\s+`)
)

func (q Query) GetQuery() (query string) {
	query = strings.TrimSpace(q.Query)
	if query != "" {
		query = regexpSpace.ReplaceAllString(query, " ")
	}
	return
}

func (q Query) GetLikePattern() (pattern string) {
	query := q.GetQuery()
	if query == "" {
		return ""
	}
	pattern = strings.ReplaceAll(query, "%", `\%`)
	pattern = strings.ReplaceAll(pattern, "_", `\_`)
	pattern = strings.ReplaceAll(pattern, " ", `%`)
	return "%" + pattern + "%"
}

func (q Query) QueryResult() (r QueryResult) {
	if query := q.GetQuery(); query != "" {
		r.Query = &query
	}
	return
}
