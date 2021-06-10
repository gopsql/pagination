package pagination_test

import (
	"testing"

	"github.com/gopsql/pagination/v2"
	"github.com/stretchr/testify/assert"
)

func TestSort01(t *testing.T) {
	s := pagination.Sort{}
	newCtx("/").Bind(&s)
	assert.Equal(t, "", s.OrderBy())
	assert.Equal(t, toJson(s.SortResult()), `{"Sort":null,"Order":null}`)
}

func TestSort02(t *testing.T) {
	s := pagination.Sort{
		DefaultSort: "created_at",
	}
	newCtx("/").Bind(&s)
	assert.Equal(t, "ORDER BY created_at DESC", s.OrderBy())
	assert.Equal(t, toJson(s.SortResult()), `{"Sort":"created_at","Order":"desc"}`)
}

func TestSort03(t *testing.T) {
	s := pagination.Sort{
		DefaultSort:  "created_at",
		DefaultOrder: "asc",
	}
	newCtx("/").Bind(&s)
	assert.Equal(t, "ORDER BY created_at ASC", s.OrderBy())
	assert.Equal(t, toJson(s.SortResult()), `{"Sort":"created_at","Order":"asc"}`)
}

func TestSort04(t *testing.T) {
	s := pagination.Sort{
		DefaultSort:  "created_at",
		DefaultOrder: "asc",
	}
	newCtx("/?sort=updated_at").Bind(&s)
	assert.Equal(t, "ORDER BY created_at ASC", s.OrderBy())
	assert.Equal(t, toJson(s.SortResult()), `{"Sort":"created_at","Order":"asc"}`)
}

func TestSort05(t *testing.T) {
	s := pagination.Sort{
		AllowedSorts: []string{"updated_at"},
		DefaultSort:  "created_at",
		DefaultOrder: "asc",
	}
	newCtx("/").Bind(&s)
	assert.Equal(t, "ORDER BY created_at ASC", s.OrderBy())
	assert.Equal(t, toJson(s.SortResult()), `{"Sort":"created_at","Order":"asc"}`)
	newCtx("/?sort=updated_at&order=desc").Bind(&s)
	assert.Equal(t, "ORDER BY updated_at DESC", s.OrderBy())
	assert.Equal(t, toJson(s.SortResult()), `{"Sort":"updated_at","Order":"desc"}`)
}

func TestSort06(t *testing.T) {
	s := pagination.Sort{
		AllowedSorts: map[string]string{
			"created_at": "updated_at",
			"updated_at": "created_at {} NULLS LAST",
		},
		DefaultSort: "created_at",
	}
	newCtx("/").Bind(&s)
	assert.Equal(t, "ORDER BY updated_at DESC", s.OrderBy())
	assert.Equal(t, toJson(s.SortResult()), `{"Sort":"created_at","Order":"desc"}`)
	newCtx("/?sort=updated_at&order=desc").Bind(&s)
	assert.Equal(t, "ORDER BY created_at DESC NULLS LAST", s.OrderBy())
	assert.Equal(t, toJson(s.SortResult()), `{"Sort":"updated_at","Order":"desc"}`)
}
