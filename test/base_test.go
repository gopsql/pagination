package pagination_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"
)

func newCtx(path string) echo.Context {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, path, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c
}

func newFiberCtx(path, qs string) *fiber.Ctx {
	app := fiber.New()
	c := &fasthttp.RequestCtx{}
	c.Request.Header.SetMethod("GET")
	c.URI().SetPath(path)
	c.URI().SetQueryString(qs)
	return app.AcquireCtx(c)
}

func toJson(obj interface{}) string {
	b, _ := json.Marshal(obj)
	return string(b)
}
