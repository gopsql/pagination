package pagination_test

import (
	"testing"

	"github.com/gopsql/pagination/v2"
	"github.com/stretchr/testify/assert"
)

func TestPaginationQuery01(t *testing.T) {
	p := pagination.PaginationQuery{
		pagination.Pagination{},
		pagination.Query{},
	}
	newCtx("/?page=2&per=15&query=foo&sort=created_at&order=asc").Bind(&p)
	assert.Equal(t, "LIMIT 15 OFFSET 15", p.LimitOffset())
	assert.Equal(t, "%foo%", p.GetLikePattern())
	assert.Equal(t, toJson(p.PaginationQueryResult(18)),
		`{"CurrentPage":2,"NextPage":null,"PrevPage":1,"TotalPages":2,"TotalCount":18,"LimitValue":15,`+
			`"Query":"foo"}`)
}

func TestPaginationSort01(t *testing.T) {
	p := pagination.PaginationSort{
		pagination.Pagination{},
		pagination.Sort{
			AllowedSorts: []string{"created_at"},
		},
	}
	newCtx("/?page=2&per=15&query=foo&sort=created_at&order=asc").Bind(&p)
	assert.Equal(t, "ORDER BY created_at ASC LIMIT 15 OFFSET 15", p.OrderByLimitOffset())
	assert.Equal(t, toJson(p.PaginationSortResult(18)),
		`{"CurrentPage":2,"NextPage":null,"PrevPage":1,"TotalPages":2,"TotalCount":18,"LimitValue":15,`+
			`"Sort":"created_at","Order":"asc"}`)
}

func TestPaginationQuerySort01(t *testing.T) {
	p := pagination.PaginationQuerySort{
		pagination.Pagination{},
		pagination.Query{},
		pagination.Sort{
			AllowedSorts: []string{"created_at"},
		},
	}
	newCtx("/?page=2&per=15&query=foo&sort=created_at&order=asc").Bind(&p)
	assert.Equal(t, "ORDER BY created_at ASC LIMIT 15 OFFSET 15", p.OrderByLimitOffset())
	assert.Equal(t, "%foo%", p.GetLikePattern())
	assert.Equal(t, toJson(p.PaginationQuerySortResult(18)),
		`{"CurrentPage":2,"NextPage":null,"PrevPage":1,"TotalPages":2,"TotalCount":18,"LimitValue":15,`+
			`"Query":"foo","Sort":"created_at","Order":"asc"}`)
}

func TestQuerySort01(t *testing.T) {
	p := pagination.QuerySort{
		pagination.Query{},
		pagination.Sort{
			AllowedSorts: []string{"created_at"},
		},
	}
	newCtx("/?page=2&per=15&query=foo&sort=created_at&order=asc").Bind(&p)
	assert.Equal(t, "ORDER BY created_at ASC", p.OrderBy())
	assert.Equal(t, "%foo%", p.GetLikePattern())
	assert.Equal(t, toJson(p.QuerySortResult()),
		`{"Query":"foo","Sort":"created_at","Order":"asc"}`)
}
