// Package space implements a function to calculate someone's age on a given planet
package space

import (
	"fmt"
	"strconv"
)

type Planet string

const (
	EarthYearSeconds = 31557600
)

var PlanetToEarth = map[Planet]float64{
	"Earth": 1,
	"Mercury": 0.2408467,
	"Venus": 0.61519726,
	"Mars": 1.8808158,
	"Jupiter": 11.862615,
	"Saturn": 29.447498,
	"Uranus": 84.016846,
	"Neptune": 164.79132,
}

// Age takes a number of seconds and a planet and returns the number of years on that planet corresponding to the number of seconds
func Age(seconds float64, planet Planet) float64 {
	age, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", seconds/EarthYearSeconds/PlanetToEarth[planet]), 2)
	return age
}
