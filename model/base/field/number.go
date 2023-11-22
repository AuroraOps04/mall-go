package field

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

type LikeNumberInt64 int64
type LikeNumberFloat64 float64

func (l LikeNumberInt64) MarshalJSON() ([]byte, error) {
	s := strconv.FormatInt(int64(l), 10)
	return []byte(s), nil
}

func (l *LikeNumberInt64) UnmarshalJSON(bytes []byte) error {
	// if is quote number  bytes[0] = 34
	if bytes[0] == 34 {
		bytes = bytes[1 : len(bytes)-1]
	}
	num, err := strconv.ParseInt(string(bytes), 10, 64)
	if err != nil {
		return err
	}
	*l = LikeNumberInt64(num)
	return nil
}

func (l LikeNumberInt64) Value() (driver.Value, error) {
	return int64(l), nil
}

func (l *LikeNumberInt64) Scan(src any) error {
	switch v := src.(type) {
	case int64:
		*l = LikeNumberInt64(v)
		return nil
	case []uint8:
		s := string(v)
		digit, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}
		*l = LikeNumberInt64(digit)
		return nil

	}
	return fmt.Errorf("not exceped type: %T", src)
}

func (l LikeNumberFloat64) MarshalJSON() ([]byte, error) {
	str := strconv.FormatFloat(float64(l), 'e', 10, 64)
	return []byte(str), nil
}

func (l *LikeNumberFloat64) UnmarshalJSON(b []byte) error {
	if b[0] == '"' {
		b = b[1 : len(b)-1]
	}
	f, err := strconv.ParseFloat(string(b), 64)
	if err != nil {
		return err
	}
	*l = LikeNumberFloat64(f)
	return nil
}

func (l LikeNumberFloat64) Value() (driver.Value, error) {
	return float64(l), nil
}

func (l *LikeNumberFloat64) Scan(src any) error {
	switch v := src.(type) {
	case float64:
		*l = LikeNumberFloat64(v)
		return nil
	case []uint8:
		digit, err := strconv.ParseFloat(string(v), 64)
		if err != nil {
			return err
		}
		*l = LikeNumberFloat64(digit)
		return nil
	}

	return fmt.Errorf("not support type: %T\n", src)
}
