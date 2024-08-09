package annalyn

// CanFastAttack can be executed only when the knight is sleeping.
func CanFastAttack(knightIsAwake bool) bool {
	if knightIsAwake == true {
		return false
	} else {
		return true
	}
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	if knightIsAwake || archerIsAwake || prisonerIsAwake == true {
		return true
	} else {
		return false
	}
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	if archerIsAwake && prisonerIsAwake == true {
		return false
	} else if archerIsAwake == false && prisonerIsAwake == true {
		return true
	} else {
		return false
	}
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	switch {
	case knightIsAwake == false && archerIsAwake == true && prisonerIsAwake == false && petDogIsPresent == false:
		return false
	case knightIsAwake == true && archerIsAwake == false && prisonerIsAwake == true && petDogIsPresent == true:
		return true
	case knightIsAwake == false && archerIsAwake == false && prisonerIsAwake == false && petDogIsPresent == true:
		return true
	case knightIsAwake == false && archerIsAwake == false && prisonerIsAwake == true && petDogIsPresent == true:
		return true
	case knightIsAwake == false && archerIsAwake == false && prisonerIsAwake == true && petDogIsPresent == false:
		return true
	case knightIsAwake == false && archerIsAwake == false && prisonerIsAwake == false && petDogIsPresent == false:
		return false
	case knightIsAwake == true && archerIsAwake == false && prisonerIsAwake == false && petDogIsPresent == true:
		return true
	default:
		return false
	}
}
