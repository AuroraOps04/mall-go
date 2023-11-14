package field

import "strconv"

type LikeNumber int8

func (l LikeNumber) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(int(l))), nil
}

func (l *LikeNumber) UnmarshalJSON(bytes []byte) error {
	// if is quote number  bytes[0] = 34
	if bytes[0] == 34 {
		bytes = bytes[1 : len(bytes)-1]
	}
	num, err := strconv.ParseInt(string(bytes), 10, 8)
	if err != nil {
		return err
	}
	*l = LikeNumber(num)
	return nil
}
