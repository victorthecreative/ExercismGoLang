package main

import "fmt"

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	productionRate_float := float64(productionRate)
	successRate_norm := successRate / 100
	var result float64 = productionRate_float * successRate_norm
	return result
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	result_per_hour := CalculateWorkingCarsPerHour(productionRate, successRate)
	result_per_hour_int := int(result_per_hour)
	result_per_minute := result_per_hour_int / 60

	return result_per_minute
}

func CalculateCost(carsCount int) uint {
	const costs_per_cars = 10000
	const costs_then_cars = 95000
	groupThen := carsCount / 10
	unit_cars := carsCount % 10
	calc_cost := groupThen*costs_then_cars + unit_cars*costs_per_cars

	total_cost := uint(calc_cost)

	return total_cost
}

func main() {
	var result int = 37 / 10
	fmt.Println(result)
}

//Portanto, o custo para 37 carros é: 3*95.000+7*10.000=355.000
