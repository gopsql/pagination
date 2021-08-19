# pagination

Convert pagination, query and/or sort query parameters to SQL expressions.

```go
func list(c echo.Context) error {
	q := pagination.PaginationQuerySort{
		pagination.Pagination{
			MaxPer:     50,
			DefaultPer: 20,
		},
		pagination.Query{},
		pagination.Sort{
			AllowedSorts: []string{
				"name",
				"created_at",
			},
			DefaultSort:  "created_at",
			DefaultOrder: "desc",
		},
	}
	c.Bind(&q)

	// ?query=admin&page=3&per=44&order=asc&sort=created_at

	fmt.Println(q.GetLikePattern())
	// %admin%

	fmt.Println(q.OrderByLimitOffset())
	// ORDER BY created_at ASC LIMIT 44 OFFSET 88

	count := 100
	b, _ := json.MarshalIndent(q.PaginationQuerySortResult(count), "", "  ")
	fmt.Println(string(b))
	// {
	//   "CurrentPage": 3,
	//   "NextPage": null,
	//   "PrevPage": 2,
	//   "TotalPages": 3,
	//   "TotalCount": 100,
	//   "LimitValue": 44,
	//   "Query": "admin",
	//   "Sort": "created_at",
	//   "Order": "asc"
	// }

	// ...
}
```
