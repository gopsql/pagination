package pagination_test

import (
	"testing"

	"github.com/gopsql/pagination/v2"
	"github.com/stretchr/testify/assert"
)

func TestBind(t *testing.T) {
	p := pagination.PaginationQuerySort{
		pagination.Pagination{},
		pagination.Query{},
		pagination.Sort{
			AllowedSorts: []string{"created_at"},
		},
	}
	c := newFiberCtx("/", "page=2&per=15&query=foo&sort=created_at&order=asc")
	pagination.Bind(&p, c.QueryParser)
	assert.Equal(t, "ORDER BY created_at ASC LIMIT 15 OFFSET 15", p.OrderByLimitOffset())
	assert.Equal(t, "%foo%", p.GetLikePattern())
	assert.Equal(t, toJson(p.PaginationQuerySortResult(18)),
		`{"CurrentPage":2,"NextPage":null,"PrevPage":1,"TotalPages":2,"TotalCount":18,"LimitValue":15,`+
			`"Query":"foo","Sort":"created_at","Order":"asc"}`)
}
