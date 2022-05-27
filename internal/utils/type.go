package utils

import (
	"encoding/json"
)

func TypeConverter[R any](data any) (res *R, err error) {
	var b []byte
	b, err = json.Marshal(&data)
	if err != nil {
		res = nil
		return
	}
	err = json.Unmarshal(b, &res)
	if err != nil {
		return
	}
	return
}
