package weather

import "fmt"

type response struct {
	city  string
	feel  string
	temp  float64
	scale string
}

func (r response) String() string {
	return fmt.Sprintf("Current wether for %s: %s %.1f%s", r.city, r.feel, r.temp, r.scale)
}
