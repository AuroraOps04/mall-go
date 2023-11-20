package field

import (
	"database/sql/driver"
	"time"
)

type DateTime time.Time

// ISO 8601
const layout = "2006-01-02T15:04:05.000Z"

func (d *DateTime) MarshalJSON() ([]byte, error) {
	str := time.Time(*d).Format(layout)
	return []byte("\"" + str + "\""), nil

}

func (d *DateTime) UnmarshalJSON(b []byte) error {
	// remove extra quote
	b = b[1 : len(b)-1]
	str := string(b)
	if str == "" {
		return nil
	}
	t, err := time.Parse(layout, str)
	*d = DateTime(t)
	if err != nil {
		return err
	}

	return nil
}
func (d *DateTime) Value() (driver.Value, error) {
	return time.Time(*d), nil

}
