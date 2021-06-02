/*
Package pagination is useful for simple search and pagination.

	func list(c echo.Context) error {
		q := pagination.Query{
			MaxPer:     50,
			DefaultPer: 20,
		}
		c.Bind(&q)
		cond := "WHERE app_id = $1"
		qt, qa := q.GetQuery()
		if qt != "" {
			cond += " AND (name ILIKE $2 OR message ILIKE $2)"
		}
		count := modelPost.MustCount(append([]interface{}{cond, appId}, qa...)...)
		posts := []models.Post{}
		sql := cond + " ORDER BY created_at DESC " + q.LimitOffset()
		modelPost.Find(append([]interface{}{sql, appId}, qa...)...).MustQuery(&posts)
		return c.JSON(200, struct {
			Posts      []models.Post
			Pagination pagination.Result
		}{posts, q.Result(count)})
	}
*/
package pagination

import (
	"fmt"
	"math"
	"strings"
)

const (
	DefaultMaxPer     = 100
	DefaultDefaultPer = 10
)

type (
	Query struct {
		Query string `query:"query"`
		Page  int    `query:"page"`
		Per   int    `query:"per"`

		MaxPer     int
		DefaultPer int
	}

	Result struct {
		CurrentPage *int
		NextPage    *int
		PrevPage    *int
		TotalPages  *int
		TotalCount  *int
		LimitValue  *int
	}
)

func (q Query) GetQuery() (query string, args []interface{}) {
	query = strings.TrimSpace(q.Query)
	if query == "" {
		return
	}
	args = append(args, "%"+strings.ReplaceAll(strings.ReplaceAll(query, "%", `\%`), "_", `\_`)+"%")
	return
}

func (q Query) LimitOffset() string {
	return fmt.Sprintf("LIMIT %d OFFSET %d", q.GetPer(), (q.GetPage()-1)*q.GetPer())
}

func (q Query) GetPage() int {
	page := q.Page
	if page < 1 {
		page = 1
	}
	return page
}

func (q Query) GetPer() int {
	per, maxPer, defaultPer := q.Per, q.MaxPer, q.DefaultPer
	if maxPer < 1 {
		maxPer = DefaultMaxPer
	}
	if defaultPer < 1 {
		defaultPer = DefaultDefaultPer
	}
	if per < 1 {
		per = defaultPer
	} else if per > maxPer {
		per = maxPer
	}
	return per
}

func (q Query) Result(count int) (p Result) {
	per, page := q.GetPer(), q.GetPage()
	p = Result{
		TotalCount:  &count,
		LimitValue:  &per,
		CurrentPage: &page,
	}
	tp, np, pp := int(math.Ceil(float64(count)/float64(per))), page+1, page-1
	p.TotalPages = &tp
	if np <= tp {
		p.NextPage = &np
	}
	if pp > 0 {
		p.PrevPage = &pp
	}
	return
}
