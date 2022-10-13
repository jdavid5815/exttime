package exttime

func (m Moonphase) String() string {

	switch m {
	case NM:
		return ("NM")
	case FQ:
		return ("FQ")
	case FM:
		return ("FM")
	case LQ:
		return ("LQ")
	}
	return ""
}

/*
 * Given a reference date of a well known New Moon, the function will calculate
 * the dates of all the new moon, full moon, first quater and last quarter dates
 * for the given year. The result is returned in a slice. Note that some results
 * might be off by a day. This is due the fact that we don't take the exact time
 * into consideration. Also, Gregorian days start at mightnight, Julian days at
 * noon. So consider this routine to give you *ballpark* dates. The results can
 * be tweaked by choosing a different reference date (where the New Moon is
 * earlier or later, depending on your situation.)
 */
func Moonphases(new_moon_ref Date, year int) []DateMoonphaseCombo {

	var (
		julian_reference_date     float32
		julian_target_date        float32
		days_since_new_moon       float32
		new_moons_since_reference float32
		gregorian                 Date
		combo                     DateMoonphaseCombo
		phases                    []DateMoonphaseCombo
		days_in_month             int
		last_nm_day               int
		last_fq_day               int
		last_fm_day               int
		last_lq_day               int
	)
	// Get our reference date in Julian Day Number format.
	julian_reference_date = JulianDayNumber(new_moon_ref)
	gregorian.Year = year
	gregorian.Hour = 12
	gregorian.Minutes = 00
	last_nm_day = -1
	last_fq_day = -1
	last_fm_day = -1
	last_lq_day = -1
	for gregorian.Month = 1; gregorian.Month <= 12; gregorian.Month++ {
		switch gregorian.Month {
		case 1, 3, 5, 7, 8, 10, 12:
			days_in_month = 31
		case 4, 6, 9, 11:
			days_in_month = 30
		case 2:
			if Leapyear(year) {
				days_in_month = 29
			} else {
				days_in_month = 28
			}
		}
		for gregorian.Day = 1; gregorian.Day <= days_in_month; gregorian.Day++ {
			julian_target_date = JulianDayNumber(gregorian)
			days_since_new_moon = julian_target_date - julian_reference_date
			new_moons_since_reference = days_since_new_moon / Synodic_Month
			new_moons_fraction := new_moons_since_reference - float32(int(new_moons_since_reference))
			switch moon_days := new_moons_fraction * Synodic_Month; {
			case moon_days < 1 || moon_days > 29:
				if gregorian.Day-last_nm_day > 1 {
					combo.Date = gregorian
					combo.Phase = NM
					phases = append(phases, combo)
					if gregorian.Day == days_in_month {
						last_nm_day = 0
					} else {
						last_nm_day = gregorian.Day
					}
				}
			case moon_days > 7.38 && moon_days < 8.38:
				if gregorian.Day-last_fq_day > 1 {
					combo.Date = gregorian
					combo.Phase = FQ
					phases = append(phases, combo)
					if gregorian.Day == days_in_month {
						last_fq_day = 0
					} else {
						last_fq_day = gregorian.Day
					}
				}
			case moon_days > 14.77 && moon_days < 15.77:
				combo.Date = gregorian
				combo.Phase = FM
				phases = append(phases, combo)
				if gregorian.Day == days_in_month {
					last_fm_day = 0
				} else {
					last_fm_day = gregorian.Day
				}
			case moon_days > 22.14 && moon_days < 23.14:
				combo.Date = gregorian
				combo.Phase = LQ
				phases = append(phases, combo)
				if gregorian.Day == days_in_month {
					last_lq_day = 0
				} else {
					last_lq_day = gregorian.Day
				}
			}
			if gregorian.Day == days_in_month {
				if last_nm_day != 0 {
					last_nm_day = -1
				}
				if last_fq_day != 0 {
					last_fq_day = -1
				}
				if last_fm_day != 0 {
					last_fm_day = -1
				}
				if last_lq_day != 0 {
					last_lq_day = -1
				}
			}
		}
	}
	return phases
}
