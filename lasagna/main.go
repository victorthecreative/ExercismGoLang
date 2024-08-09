package lasagna

// TODO: define the 'OvenTime' constant
const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actual int) int {

	return OvenTime - actual
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	const timelayer = 2
	preparationLayerTime := timelayer * numberOfLayers
	return preparationLayerTime

}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {

	time_preparation := PreparationTime(numberOfLayers)
	elapsedTime := actualMinutesInOven + time_preparation
	return elapsedTime
}
