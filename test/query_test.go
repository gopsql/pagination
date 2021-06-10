package pagination_test

import (
	"testing"

	"github.com/gopsql/pagination/v2"
	"github.com/stretchr/testify/assert"
)

func TestQuery01(t *testing.T) {
	q := pagination.Query{}
	newCtx("/").Bind(&q)
	assert.Equal(t, "", q.GetQuery())
	assert.Equal(t, "", q.GetLikePattern())
	assert.Equal(t, toJson(q.QueryResult()), `{"Query":null}`)
}

func TestQuery02(t *testing.T) {
	q := pagination.Query{}
	newCtx("/?query=+hello++world+++%251+").Bind(&q)
	assert.Equal(t, "hello world %1", q.GetQuery())
	assert.Equal(t, `%hello%world%\%1%`, q.GetLikePattern())
	assert.Equal(t, toJson(q.QueryResult()), `{"Query":"hello world %1"}`)
}
