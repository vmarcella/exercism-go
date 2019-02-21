package space

// Map the amount of time in seconds to earth years
func Age(age float64, planet string) float64 {
    // Map all the planets to their orbital period in seconds
    planetMap := map[string]float64{
        "Earth": 31557600,
        "Mercury":7600608,
        "Venus": 19400861,
        "Mars": 59356800,
        "Jupiter": 374099427,
        "Saturn": 928656297,
        "Uranus": 2649555255,
        "Neptune": 5196859068,
    }

    // Return the age divided by the orbital period in seconds
    return age / planetMap[planet]
}
