package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPrepTime int) (time int) {
	if avgPrepTime <= 0 {
		avgPrepTime = 2
	}

	return len(layers) * avgPrepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodleGrams int, sauceLiters float64) {
	noodleLayers := 0
	sauceLayers := 0
	for _, layer := range layers {
		if layer == "noodles" {
			noodleLayers++
		}
		if layer == "sauce" {
			sauceLayers++
		}
	}
	return noodleLayers * 50, float64(sauceLayers) * float64(0.2)
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendIngredients []string, ownIngredients []string) {
	ownIngredients[len(ownIngredients)-1] = friendIngredients[len(friendIngredients)-1]

}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) []float64 {
	scaledAmounts := []float64{}
	scalingFactor := float64(portions) / 2.0

	for i := 0; i < len(amounts); i++ {
		scaledAmounts = append(scaledAmounts, amounts[i]*float64(scalingFactor))
	}
	return scaledAmounts
}
