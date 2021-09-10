package pagination

import (
	"fmt"
	"math"
)

var (
	// Default MaxPer value if Pagination's MaxPer is not set.
	DefaultMaxPer = 100

	// Default DefaultPer value if Pagination's DefaultPer is not set.
	DefaultDefaultPer = 10
)

type (
	Pagination struct {
		// The page value.
		Page int `query:"page"`

		// The per value.
		Per int `query:"per"`

		// Maximum per value.
		MaxPer int

		// DefaultPer is used if the per value is empty.
		DefaultPer int
	}

	PaginationResult struct {
		CurrentPage *int
		NextPage    *int
		PrevPage    *int
		TotalPages  *int
		TotalCount  *int
		LimitValue  *int
	}
)

func (p Pagination) Limit() int {
	return p.GetPer()
}

func (p Pagination) Offset() int {
	return (p.GetPage() - 1) * p.GetPer()
}

func (p Pagination) LimitOffset() string {
	return fmt.Sprintf("LIMIT %d OFFSET %d", p.Limit(), p.Offset())
}

func (p Pagination) GetPage() int {
	page := p.Page
	if page < 1 {
		page = 1
	}
	return page
}

func (p Pagination) GetPer() int {
	per, maxPer, defaultPer := p.Per, p.MaxPer, p.DefaultPer
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

func (p Pagination) PaginationResult(count int) (r PaginationResult) {
	per, page := p.GetPer(), p.GetPage()
	r = PaginationResult{
		TotalCount:  &count,
		LimitValue:  &per,
		CurrentPage: &page,
	}
	tp, np, pp := int(math.Ceil(float64(count)/float64(per))), page+1, page-1
	r.TotalPages = &tp
	if np <= tp {
		r.NextPage = &np
	}
	if pp > 0 {
		r.PrevPage = &pp
	}
	return
}
