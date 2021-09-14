// Package weather works with weather data.
package weather

var (
	// CurrentCondition keeps the current weather.
	CurrentCondition string
	// CurrentLocation stores the current location.
	CurrentLocation string
)

// Log prepares a message for logging the weather.
func Log(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
