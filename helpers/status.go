package helpers

func WaterStatus(water uint8) string {
	switch {
	case water <= 5:
		return "aman"
	case water >= 6 && water <= 8:
		return "siaga"
	default:
		return "bahaya"
	}
}

func WindStatus(wind uint8) string {
	switch {
	case wind <= 6:
		return "aman"
	case wind >= 7 && wind <= 15:
		return "siaga"
	default:
		return "bahaya"
	}
}