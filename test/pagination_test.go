package pagination_test

import (
	"testing"

	"github.com/gopsql/pagination/v2"
	"github.com/stretchr/testify/assert"
)

func TestPagination01(t *testing.T) {
	p := pagination.Pagination{}
	newCtx("/").Bind(&p)
	assert.Equal(t, "LIMIT 10 OFFSET 0", p.LimitOffset())
	assert.Equal(t, toJson(p.PaginationResult(33)),
		`{"CurrentPage":1,"NextPage":2,"PrevPage":null,"TotalPages":4,"TotalCount":33,"LimitValue":10,"From":1,"To":10}`)
}

func TestPagination02(t *testing.T) {
	p := pagination.Pagination{}
	newCtx("/?per=20").Bind(&p)
	assert.Equal(t, "LIMIT 20 OFFSET 0", p.LimitOffset())
	assert.Equal(t, toJson(p.PaginationResult(34)),
		`{"CurrentPage":1,"NextPage":2,"PrevPage":null,"TotalPages":2,"TotalCount":34,"LimitValue":20,"From":1,"To":20}`)
}

func TestPagination03(t *testing.T) {
	p := pagination.Pagination{}
	newCtx("/?per=101").Bind(&p)
	assert.Equal(t, "LIMIT 100 OFFSET 0", p.LimitOffset())
	assert.Equal(t, toJson(p.PaginationResult(0)),
		`{"CurrentPage":1,"NextPage":null,"PrevPage":null,"TotalPages":0,"TotalCount":0,"LimitValue":100,"From":null,"To":null}`)
}

func TestPagination04(t *testing.T) {
	p := pagination.Pagination{}
	newCtx("/?per=50&page=4").Bind(&p)
	assert.Equal(t, "LIMIT 50 OFFSET 150", p.LimitOffset())
	assert.Equal(t, toJson(p.PaginationResult(444)),
		`{"CurrentPage":4,"NextPage":5,"PrevPage":3,"TotalPages":9,"TotalCount":444,"LimitValue":50,"From":151,"To":200}`)
}

func TestPagination05(t *testing.T) {
	p := pagination.Pagination{
		MaxPer:     5,
		DefaultPer: 1,
	}
	newCtx("/").Bind(&p)
	assert.Equal(t, "LIMIT 1 OFFSET 0", p.LimitOffset())
	assert.Equal(t, toJson(p.PaginationResult(9)),
		`{"CurrentPage":1,"NextPage":2,"PrevPage":null,"TotalPages":9,"TotalCount":9,"LimitValue":1,"From":1,"To":1}`)
}

func TestPagination06(t *testing.T) {
	p := pagination.Pagination{
		MaxPer:     5,
		DefaultPer: 1,
	}
	newCtx("/?per=3&page=6").Bind(&p)
	assert.Equal(t, "LIMIT 3 OFFSET 15", p.LimitOffset())
	assert.Equal(t, toJson(p.PaginationResult(100)),
		`{"CurrentPage":6,"NextPage":7,"PrevPage":5,"TotalPages":34,"TotalCount":100,"LimitValue":3,"From":16,"To":18}`)
}

func TestPagination07(t *testing.T) {
	p := pagination.Pagination{
		MaxPer:     5,
		DefaultPer: 1,
	}
	newCtx("/?per=10").Bind(&p)
	assert.Equal(t, "LIMIT 5 OFFSET 0", p.LimitOffset())
	assert.Equal(t, toJson(p.PaginationResult(20)),
		`{"CurrentPage":1,"NextPage":2,"PrevPage":null,"TotalPages":4,"TotalCount":20,"LimitValue":5,"From":1,"To":5}`)
}
