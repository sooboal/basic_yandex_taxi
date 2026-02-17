package wether

type WeatherCondition int

const (
	Clear WeatherCondition = iota // Ключевое слово iota присваивает каждой константе числовое значение по порядку (0, 1, 2, 3 и т.д.)
	Rain
	HeavyRain
	Snow
)

type WeatherData struct {
	Condition WeatherCondition
	WindSpeed int
}

func GetWeatherMultiplier(weather WeatherData) float64 {
	coef := 1.0

	switch weather.Condition {
	case HeavyRain:
		coef += 0.2
	case Rain:
		coef += 0.125
	case Snow:
		coef += 0.15
	}

	if weather.WindSpeed > 15 {
		coef += 0.1
	}

	return coef
}
