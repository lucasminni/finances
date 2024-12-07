package income

import "time"

type Income struct {
	Name        string
	Description string
	Value       float64
	Date        time.Time
}
