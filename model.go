package weather

import "fmt"

type Conditions struct {
	City  string
	Feel  string
	TempK float64
}

func (c Conditions) String() string {
	return fmt.Sprintf("Current wether for %s: %s %.1fK", c.City, c.Feel, c.TempK)
}

func (c Conditions) TempFahrenheit() float64 {
	return 1.8*c.TempCelsius() + 32
}

func (c Conditions) StringFahrenheit() string {
	return fmt.Sprintf("Current wether for %s: %s %.1fºF", c.City, c.Feel, c.TempFahrenheit())
}

func (c Conditions) TempCelsius() float64 {
	return c.TempK - 273.5
}

func (c Conditions) StringCelsius() string {
	return fmt.Sprintf("Current wether for %s: %s %.1fºC", c.City, c.Feel, c.TempCelsius())
}
