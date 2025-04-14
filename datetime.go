package cryptocloud

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	DateFormat     = "02.01.2006"
	DateTimeFormat = "2006-01-02 15:04:05.999999"
)

type Date time.Time

func NewDate(year int, month time.Month, day int) Date {
	return Date(time.Date(year, month, day, 0, 0, 0, 0, time.Local))
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d).Format(DateFormat))
}

func (d Date) Time() time.Time {
	return time.Time(d)
}

type DateTime time.Time

func (d *DateTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("failed to unmarshal string: %w", err)
	}

	t, err := time.Parse(DateTimeFormat, s)
	if err != nil {
		return fmt.Errorf("failed to parse time: %w", err)
	}

	*d = DateTime(t)

	return nil
}

func (d *DateTime) Time() time.Time {
	if d == nil {
		return time.Time{}
	}

	return time.Time(*d)
}
