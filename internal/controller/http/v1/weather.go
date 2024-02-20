package v1

import (
	"github.com/evrone/go-clean-template/internal/usecase"
	"github.com/evrone/go-clean-template/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type weatherRoutes struct {
	w usecase.Weather
	l logger.Interface
}

func newWeatherRoutes(handler *gin.RouterGroup, w usecase.Weather, l logger.Interface) {
	r := &weatherRoutes{w: w, l: l}

	h := handler.Group("/weather")
	{
		h.GET("/:city", r.GetWeather)
		h.PUT("/:city", r.UpdateWeather)
	}
}

func (r *weatherRoutes) GetWeather(c *gin.Context) {
	city := c.Param("city")

	weather, err := r.w.Get(c, city)
	if err != nil {
		r.l.Error(err, "http - v1 - get")
		errorResponse(c, http.StatusInternalServerError, "failed to get")

		return
	}

	c.JSON(http.StatusOK, weather)
}

func (r *weatherRoutes) UpdateWeather(c *gin.Context) {
	city := c.Param("city")

	err := r.w.Update(c, city)
	if err != nil {
		r.l.Error(err, "http - v1 - update")
		errorResponse(c, http.StatusInternalServerError, "failed to update")

		return
	}

	c.JSON(http.StatusOK, "success")
}
