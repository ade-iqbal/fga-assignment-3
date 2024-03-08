package helpers

import (
	"math/rand"

	"github.com/ade-iqbal/fga-assignment-3/models"
)

func GenerateRandom(weather models.Weather) models.Weather {
	for {
		randomWater := rand.Intn(101)
		randomWind := rand.Intn(101)

		if (randomWater >= 1 && randomWater <= 100) && (randomWind >= 1 && randomWind <= 100) {
			weather.Status.Water = uint8(randomWater)
			weather.Status.Wind = uint8(randomWind)
			break
		}
	}

	return weather
}