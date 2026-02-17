package main

import (
	"fmt"
	"time"

	"project/internal/aggregator"
	"project/internal/baseprice"
	"project/internal/traffic"
	w "project/internal/weather"
)

func main() {
	calculator := aggregator.PriceCalculator{
		TrafficClient: &traffic.RealTrafficClient{}, // В продакшене используется настоящий клиент, а мы подключим структуру-заглушку для имитации его работы
	}

	price := calculator.CalculatePrice(
		baseprice.TripParameters{Distance: 8.5, Duration: 20},
		time.Now(),
		w.WeatherData{Condition: w.HeavyRain, WindSpeed: 10},
		55.751244, 37.618423,
	)

	fmt.Printf("Ваша цена: %.0f руб.\n", price)
}