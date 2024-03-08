package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ade-iqbal/fga-assignment-3/helpers"
	"github.com/ade-iqbal/fga-assignment-3/models"
	"github.com/gin-gonic/gin"
)

func GetTemplate(ctx *gin.Context) {
	var (
		content = helpers.GetFile()
		Weather models.Weather
	)

	json.Unmarshal([]byte(content), &Weather)

	newWeather := helpers.GenerateRandom(Weather)
	helpers.UpdateFile(newWeather)

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"water": Weather.Status.Water,
		"wind": Weather.Status.Wind,
		"water_status": helpers.WaterStatus(Weather.Status.Water),
		"wind_status": helpers.WindStatus(Weather.Status.Wind),
	})
}