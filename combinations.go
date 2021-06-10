package pagination

import "strings"

type (
	PaginationQuery struct {
		Pagination
		Query
	}

	PaginationQueryResult struct {
		PaginationResult
		QueryResult
	}
)

func (pq PaginationQuery) PaginationQueryResult(count int) PaginationQueryResult {
	return PaginationQueryResult{
		pq.PaginationResult(count),
		pq.QueryResult(),
	}
}

type (
	PaginationSort struct {
		Pagination
		Sort
	}

	PaginationSortResult struct {
		PaginationResult
		SortResult
	}
)

func (ps PaginationSort) OrderByLimitOffset() string {
	return strings.TrimSpace(ps.OrderBy() + " " + ps.LimitOffset())
}

func (ps PaginationSort) PaginationSortResult(count int) PaginationSortResult {
	return PaginationSortResult{
		ps.PaginationResult(count),
		ps.SortResult(),
	}
}

type (
	PaginationQuerySort struct {
		Pagination
		Query
		Sort
	}

	PaginationQuerySortResult struct {
		PaginationResult
		QueryResult
		SortResult
	}
)

func (pqs PaginationQuerySort) OrderByLimitOffset() string {
	return strings.TrimSpace(pqs.OrderBy() + " " + pqs.LimitOffset())
}

func (pqs PaginationQuerySort) PaginationQuerySortResult(count int) PaginationQuerySortResult {
	return PaginationQuerySortResult{
		pqs.PaginationResult(count),
		pqs.QueryResult(),
		pqs.SortResult(),
	}
}

type (
	QuerySort struct {
		Query
		Sort
	}

	QuerySortResult struct {
		QueryResult
		SortResult
	}
)

func (qs QuerySort) QuerySortResult() QuerySortResult {
	return QuerySortResult{
		qs.QueryResult(),
		qs.SortResult(),
	}
}
