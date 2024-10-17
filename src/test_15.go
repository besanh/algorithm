package test15

func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {
	if y1 < y2 {
		return 0 // No fine if the return is on or before the due date
	} else if y1 > y2 {
		return 10000 // Fine for returning in a later year
	} else { // y1 == y2
		if m1 < m2 {
			return 0 // No fine if the return is before the due month
		} else if m1 > m2 {
			return 500 * (m1 - m2) // Fine for returning in a later month
		} else { // m1 == m2
			if d1 <= d2 {
				return 0 // No fine if the return is on or before the due day
			} else {
				return 15 * (d1 - d2) // Fine for returning after the due day
			}
		}
	}
}
