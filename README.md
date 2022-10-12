
# exttime
# [![Language Badge](https://img.shields.io/badge/Language-Go-blue.svg)](https://go.dev) [![License: GPL v3](https://img.shields.io/github/license/jdavid5815/exttime)](https://www.gnu.org/licenses/gpl-3.0) [![Issues](https://img.shields.io/github/issues/jdavid5815/exttime)](https://github.com/jdavid5815/exttime/issues) [![GoDoc](https://godoc.org/github.com/jdavid5815/exttime?status.svg)](https://godoc.org/github.com/jdavid5815/exttime)

Time and date related functions.

## *type Date*

This type is a struct that contains a Year, Month and Day field for use in various other time and date related functions.

## *type Moonphase*
Enumeration type that can assume the values NM (New Moon), FQ (First Quarter), FM (Full Moon) and LQ (Last Quarter).

## *type DateMoonphaseCombo*

This is a struct containing a *Date* and *Moonphase* field.

## *const Synodic_Month*

The constant *Synodic_Month* (also known as lunation) is the average period of the Moon's orbit with respect to the line joining the Sun and Earth: 29 d 12 h 44 min and 2.9 s. This is the period of the lunar phases, because the Moon's appearance depends on the position of the Moon with respect to the Sun as seen from Earth.

## *const NM,FQ,FM and LQ*

These are constants of Moonphase type.

## *func MonthToInteger (month string) int*

Returns the integer representation of a month in English or of its abbreviation. E.g. January, Feb, March, Arp, etc.. If the month cannot be converted, zero is returned.

## *func JulianDayNumber(gregorian Date) float32*

Calculate the Julian Day Number, when the Gregorian date is given. The function is only valid for the year 1582 and later. If an earlier year is given, 0.0 will be returned. The result is the Julian Day Number for the beginning of the date in question at 0 hours, UTC. Note that this always gives you a half day extra. That is because the Julian Day begins at noon, UTC. This is convenient for astronomers (who until recently only observed at night), but it is confusing. The *date* type is a struct containing the integers year, month and day.

## *func Leapyear(year int) bool*

Pass the year as an integer and the function will return true if it's a leapyear, or false otherwise.

## *func Moonphases(new_moon_ref Date, year int) []DateMoonphaseCombo*

Given a reference date of a well known New Moon - given by new_moon_ref - the function will calculate the dates of all the new moon, full moon, first quater and last quarter dates
for the given year - parameter year. The result is returned in a slice of DateMoonphaseCombo. Note that results might be off by a day. This is due the fact that we don't take the exact time into consideration, that the moon's orbit is not perfectly circular and also Gregorian days start at mightnight, Julian days at noon. So consider this routine to give you *ballpark* dates. The results can be tweaked by choosing a different reference date (where the New Moon is * earlier or later, depending on your situation.) Astronomers should
definitely not use this function!
