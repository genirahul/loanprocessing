package custom

import (
	"encoding/json"
	"time"
)

const (
	FORMAT = "2006-01-02"
)

// Date Custom time.Time type which help JSON binding of date in 'yyyy-mm-dd' format
type Date time.Time

var _ json.Unmarshaler = &Date{}

// UnmarshalJSON used to unmarshel time from json in 'yyyy-mm-dd' format
func (d *Date) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation(FORMAT, s, time.UTC)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

// Sub return difference between two dates.
func (d Date) Sub(d1 Date) time.Duration {
	t1 := time.Time(d)
	t2 := time.Time(d1)
	return t1.Sub(t2)
}

func (d Date) String() string {
	return time.Time(d).Format(FORMAT)
}

// Parse Converts string to Date
func Parse(s string) (Date, error) {
	t, err := time.ParseInLocation(FORMAT, s, time.UTC)
	return Date(t), err
}
