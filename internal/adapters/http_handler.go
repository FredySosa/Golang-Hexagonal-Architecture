package adapters

import (
	"net/http"

	"github.com/FredySosa/hexArch/internal/core/domain"
	"github.com/FredySosa/hexArch/internal/ports"
	"github.com/labstack/echo/v4"
)

type (
	HTTPHandler struct {
		countriesService ports.CountriesService
	}
)

func NewHTTPHandler(cService ports.CountriesService) *echo.Echo {
	h := HTTPHandler{
		countriesService: cService,
	}

	e := echo.New()
	h.initRoutes(e)

	return e
}

func (h HTTPHandler) initRoutes(e *echo.Echo) {
	e.GET("/ping", h.Ping)

	countriesGroup := e.Group("/countries")
	countriesGroup.GET("", h.GetCountries)
}

func (h HTTPHandler) Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, "pong")
}

func (h HTTPHandler) GetCountries(c echo.Context) error {
	ctx := c.Request().Context()
	limit, offset := "100", "0"

	queryData := domain.QueryData{
		Offset: offset,
		Limit:  limit,
	}
	countries, newOffset, err := h.countriesService.GetCountries(ctx, queryData)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, "some error happened")
	}
	c.Response().Header().Add("Offset", newOffset)

	return c.JSON(http.StatusOK, countries)
}
