module github.com/gopsql/pagination/v2/test

go 1.15

replace github.com/gopsql/pagination/v2 => ../

require (
	github.com/gofiber/fiber/v2 v2.17.0
	github.com/gopsql/pagination/v2 v2.0.0-00010101000000-000000000000
	github.com/labstack/echo/v4 v4.3.0
	github.com/stretchr/testify v1.4.0
	github.com/valyala/fasthttp v1.29.0
)
