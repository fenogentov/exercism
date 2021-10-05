package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, l := range layers {
		if l == "noodles" {
			noodles += 50
		}
		if l == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) []string {
	myList = append(myList, friendsList[len(friendsList)-1])
	return myList
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(inputList []float64, portions int) (out []float64) {
	for _, e := range inputList {
		out = append(out, e/2*float64(portions))
	}
	return out
}
