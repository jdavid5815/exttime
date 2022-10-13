package exttime

type Date struct {
	Year    int
	Month   int
	Day     int
	Hour    int
	Minutes int
}
type Moonphase uint8
type DateMoonphaseCombo struct {
	Date  Date
	Phase Moonphase
}

const Synodic_Month = 29.5305888531
const (
	NM Moonphase = iota // New Moon
	FQ                  // First Quarter
	FM                  // FUll Moon
	LQ                  // Last Quarter
)
