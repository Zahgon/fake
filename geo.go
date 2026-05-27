package fake

// Latitude generates latitude (from -90.0 to 90.0)
func Latitude() float32 { _ = "STUB: not implemented"; return 0 }

// LatitudeDegrees generates latitude degrees (from -90 to 90)
func LatitudeDegrees() int { _ = "STUB: not implemented"; return 0 }

// LatitudeMinutes generates latitude minutes (from 0 to 60)
func LatitudeMinutes() int {
	_ = "STUB: not implemented"

	// LatitudeSeconds generates latitude seconds (from 0 to 60)
	return 0
}

func LatitudeSeconds() int {
	_ = "STUB: not implemented"

	// LatitudeDirection generates latitude direction (N(orth) o S(outh))
	return 0
}

func LatitudeDirection() string { _ = "STUB: not implemented"; return "" }

// Longitude generates longitude (from -180 to 180)
func Longitude() float32 { _ = "STUB: not implemented"; return 0 }

// LongitudeDegrees generates longitude degrees (from -180 to 180)
func LongitudeDegrees() int { _ = "STUB: not implemented"; return 0 }

// LongitudeMinutes generates (from 0 to 60)
func LongitudeMinutes() int {
	_ = "STUB: not implemented"

	// LongitudeSeconds generates (from 0 to 60)
	return 0
}

func LongitudeSeconds() int {
	_ = "STUB: not implemented"

	// LongitudeDirection generates (W(est) or E(ast))
	return 0
}

func LongitudeDirection() string { _ = "STUB: not implemented"; return "" }
