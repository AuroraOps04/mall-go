package field

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestName(t *testing.T) {

	type A struct {
		Num LikeNumber `json:"num"`
	}
	a := A{
		Num: 3,
	}
	var b A
	s, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("str: ", string(s))
	err = json.Unmarshal(s, &b)
	fmt.Println(b)
	err = json.Unmarshal([]byte(`{"num": 2}`), &b)
	fmt.Println(b)

	err = json.Unmarshal([]byte(`{"num": "123.3"}`), &b)
	fmt.Println(b, err)

}
