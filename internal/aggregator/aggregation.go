package aggregator

import (
	"math"
	"time"

	"project/internal/baseprice"
	"project/internal/timecoefficients"
	"project/internal/traffic"
	"project/internal/weather"
)

var (
	minPrice = 99.0
	maxPrice = 20000.0
)

type PriceCalculator struct {
	TrafficClient traffic.TrafficClient
}

func ApplyPriceLimits(price float64) float64 {
	if price < minPrice {
		return minPrice
	} else if price > maxPrice {
		return maxPrice
	}
	return price
}

func (c *PriceCalculator) CalculatePrice(trip baseprice.TripParameters, now time.Time, w weather.WeatherData, lat, lng float64) float64 {
	base := baseprice.CalculateBasePrice(trip)
	timeMult := timecoefficients.GetTimeMultiplier(now)
	weatherMult := weather.GetWeatherMultiplier(w)
	trafficMult := traffic.GetTrafficMultiplier(c.TrafficClient.GetTrafficLevel(lat, lng))

	finalPrice := base * timeMult * weatherMult * trafficMult

	return ApplyPriceLimits(math.Round(finalPrice))
}
